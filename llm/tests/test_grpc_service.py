import pytest
import requests
import subprocess
import time
import grpc
from concurrent import futures
import os
import signal

from service.model_runner import ModelRunner
from service.config import OLLAMA_MODEL, GRPC_PORT, OLLAMA_PORT
from llm_grpc import llm_service_pb2_grpc, llm_service_pb2
from service.logger import get_logger
from service.server import LLMService

logger = get_logger("test_auto")  # логгер для автоматических тестов


# -----------------------------
# Fixture для Ollama сервера
# -----------------------------
@pytest.fixture(scope="session")
def ollama_server():
    """
    Поднимает локальный Ollama сервер перед тестами.
    После окончания тестов сервер корректно останавливается.
    """
    # Запускаем Ollama в фоне (stdout/stderr перенаправлены)
    proc = subprocess.Popen(
        ["ollama", "serve"],
        stdout=subprocess.PIPE,
        stderr=subprocess.PIPE,
        preexec_fn=os.setsid  # чтобы можно было убить процесс группой
    )
    logger.info("Запуск Ollama сервера...")

    # Ждём доступность сервера
    timeout = 30
    start_time = time.time()
    while time.time() - start_time < timeout:
        try:
            r = requests.get(
                f"http://localhost:{OLLAMA_PORT}/api/tags", timeout=3)
            if r.status_code == 200:
                logger.info("Ollama сервер готов")
                break
        except requests.RequestException:
            time.sleep(1)
    else:
        # Если не поднялся, убиваем процесс и падаем
        os.killpg(os.getpgid(proc.pid), signal.SIGTERM)
        pytest.fail("Не удалось поднять Ollama сервер")

    yield  # сюда пойдут тесты

    # Teardown: останавливаем сервер после тестов
    os.killpg(os.getpgid(proc.pid), signal.SIGTERM)
    logger.info("Ollama сервер остановлен")


# --------------------------------
# Fixture для gRPC сервера
# --------------------------------
@pytest.fixture(scope="session")
def grpc_server(ollama_server):
    """
    Поднимает gRPC сервер с ModelRunner.
    Требует, чтобы Ollama сервер уже был запущен.
    """
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=2))
    llm_service_pb2_grpc.add_LLMServiceServicer_to_server(
        LLMService(ModelRunner(ensure_ollama=False)), server)
    server.add_insecure_port(f"[::]:{GRPC_PORT}")
    server.start()
    logger.info(f"gRPC сервер запущен на порту {GRPC_PORT}")

    yield  # сюда пойдут тесты

    # Teardown: останавливаем gRPC сервер
    server.stop(0)
    logger.info("gRPC сервер остановлен")


# --------------------------------
# Тест прямого вызова модели
# --------------------------------
@pytest.mark.model
def test_model_generate_direct(ollama_server):
    """
    Проверка работы прямого вызова ollama.generate.
    Модель должна вернуть непустой ответ.
    """
    import ollama
    prompt = "Привет, скажи что-нибудь короткое"

    response = ollama.generate(model=OLLAMA_MODEL, prompt=prompt, stream=False)
    assert response is not None
    logger.info(f"Модель ответила: {response}")


# --------------------------------
# Тест интеграции ModelRunner
# --------------------------------
@pytest.mark.integration
def test_integration_model_runner(ollama_server):
    """
    Проверка полной интеграции через ModelRunner:
    формирование prompt, вызов Ollama и получение ответа.
    """
    runner = ModelRunner(model_name=OLLAMA_MODEL)
    query = "Сколько просмотров у поста?"
    stats = '{"views": 123, "likes": 10}'

    result = runner.process(query, stats, context="Тест")
    assert "answer" in result
    assert result["answer"]  # ответ не пустой
    assert result["model"] == OLLAMA_MODEL
    logger.info(f"Интеграционный тест успешен: {result['answer']}")


# --------------------------------
# Тест gRPC сервиса
# --------------------------------
@pytest.mark.grpc
def test_grpc_service(grpc_server):
    """
    Проверка работы gRPC сервиса с настоящей моделью.
    Отправляем AnalyzeRequest, ожидаем AnalyzeResponse с непустым answer.
    """
    channel = grpc.insecure_channel(f"localhost:{GRPC_PORT}")
    stub = llm_service_pb2_grpc.LLMServiceStub(channel)

    request = llm_service_pb2.AnalyzeRequest(
        user_query="Привет, как дела?",
        stats_json='{"views": 10, "likes": 2}',
        context="Тест"
    )

    response = stub.Analyze(request, timeout=5)
    assert hasattr(response, "answer")
    assert response.answer  # проверяем, что answer не пустой
    logger.info(f"gRPC сервис ответил: {response.answer}")
