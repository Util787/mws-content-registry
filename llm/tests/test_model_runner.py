import pytest
import requests
import time
import grpc
from concurrent import futures

from service.model_runner import ModelRunner
from service.config import OLLAMA_MODEL, GRPC_PORT, OLLAMA_PORT
from service.logger import get_logger
from llm_grpc import llm_service_pb2_grpc, llm_service_pb2

# Логгер для тестов, которые используют реальный сервер Ollama и gRPC
logger = get_logger("test_with_server")


# --------------------------------
# Тест доступности Ollama сервера
# --------------------------------
@pytest.mark.server
def test_server_running():
    """
    Проверка доступности локального Ollama сервера через HTTP.
    Пять попыток с таймаутом 3 секунды.
    """
    for attempt in range(5):
        try:
            r = requests.get(
                f"http://localhost:{OLLAMA_PORT}/api/tags", timeout=3)
            assert r.status_code == 200  # сервер отвечает OK
            logger.info("Ollama сервер доступен")
            return
        except requests.RequestException as e:
            logger.warning(f"Попытка {attempt+1}: {e}")
            time.sleep(2)
    pytest.fail("Сервер Ollama не отвечает после нескольких попыток")


# --------------------------------
# Тест прямого вызова модели
# --------------------------------
@pytest.mark.model
def test_model_generate_direct():
    """
    Проверка прямого вызова ollama.generate без использования ModelRunner.
    Модель должна вернуть непустой ответ.
    """
    import ollama
    prompt = "Привет, скажи что-нибудь короткое"

    for attempt in range(5):
        try:
            response = ollama.generate(
                model=OLLAMA_MODEL, prompt=prompt, stream=False)
            break  # если вызов успешен — выходим из цикла
        except Exception as e:
            logger.warning(f"Попытка {attempt+1}: {e}")
            time.sleep(2)
    else:
        pytest.fail("Модель не отвечает на прямой вызов generate")

    assert response is not None
    logger.info(f"Модель ответила: {response}")


# --------------------------------
# Тест интеграции ModelRunner
# --------------------------------
@pytest.mark.integration
def test_integration_model_runner():
    """
    Полная интеграция через ModelRunner:
    формируем prompt, вызываем Ollama и проверяем результат.
    """
    runner = ModelRunner(model_name=OLLAMA_MODEL)
    query = "Сколько просмотров у поста?"
    stats = '{"views": 123, "likes": 10}'

    for attempt in range(5):
        try:
            result = runner.process(query, stats, context="Тест")
            break
        except Exception as e:
            logger.warning(f"Попытка {attempt+1}: {e}")
            time.sleep(2)
    else:
        pytest.fail("Модель не отвечает через ModelRunner")

    # Проверяем, что результат содержит ответ и модель правильная
    assert "answer" in result
    assert result["answer"], "Ответ модели пустой"
    assert result["model"] == OLLAMA_MODEL
    logger.info(f"Интеграционный тест успешен: {result['answer']}")


# --------------------------------
# Тест gRPC сервиса
# --------------------------------
@pytest.mark.grpc
def test_grpc_service():
    """
    Проверка работы gRPC сервиса:
    отправляем AnalyzeRequest и ожидаем AnalyzeResponse с непустым answer.
    Пять попыток на случай временной недоступности.
    """
    channel = grpc.insecure_channel(f"localhost:{GRPC_PORT}")
    stub = llm_service_pb2_grpc.LLMServiceStub(channel)

    request = llm_service_pb2.AnalyzeRequest(
        user_query="Привет, как дела?",
        stats_json='{"views": 10, "likes": 2}',
        context="Тест"
    )

    for attempt in range(5):
        try:
            response = stub.Analyze(request, timeout=5)
            assert response.answer != ""  # проверяем, что ответ не пустой
            logger.info(f"gRPC response: {response}")
            break
        except grpc.RpcError as e:
            logger.warning(f"Попытка {attempt+1}: {e}")
            time.sleep(2)
    else:
        pytest.fail("gRPC сервис не отвечает")
