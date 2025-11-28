import threading
import grpc
from concurrent import futures
from .logger import get_logger
from .config import GRPC_PORT
from .model_runner import ModelRunner, wait_for_ollama
from llm_grpc import llm_service_pb2_grpc, llm_service_pb2

logger = get_logger("LLMService")  # логгер для gRPC сервиса
# блокировка для последовательного доступа к модели при многопоточном gRPC
_lock = threading.Lock()


class LLMService(llm_service_pb2_grpc.LLMServiceServicer):
    """
    gRPC сервис для работы с моделью.
    Наследуется от сгенерированного класса Servicer из .proto файла.
    """

    def __init__(self, model_runner: ModelRunner):
        self.runner = model_runner  # экземпляр ModelRunner, который выполняет вызовы модели

    def Analyze(self, request, context):
        """
        Метод Analyze вызывается клиентом через gRPC.
        Формирует prompt из запроса пользователя и вызывает ModelRunner.
        """
        logger.info("Получен запрос: user_query=%s", request.user_query)
        with _lock:  # синхронизация вызова модели в многопоточном окружении
            result = self.runner.process(
                request.user_query, request.stats_json, request.context)
        logger.info("Responding")  # логируем отправку ответа
        return llm_service_pb2.AnalyzeResponse(
            # преобразуем результат в protobuf-объект
            answer=result.get("answer", ""),
            # дополнительное поле reasoning, если есть
            reasoning=result.get("reasoning", ""),
            # дополнительное поле structured, если есть
            structured=result.get("structured", "")
        )


def serve(max_workers: int = 4):
    """
    Поднимает gRPC сервер и регистрирует сервис LLMService.
    max_workers: количество потоков в ThreadPoolExecutor.
    """
    try:
        # проверка доступности Ollama перед стартом
        wait_for_ollama(timeout=10)
    except Exception:
        logger.warning(
            "Ollama не доступен перед стартом gRPC; сервер будет пытаться обращаться по мере необходимости")

    # создаём экземпляр ModelRunner без ожидания Ollama
    runner = ModelRunner(ensure_ollama=False)
    server = grpc.server(
        futures.ThreadPoolExecutor(max_workers=max_workers),
        options=[
            # максимальный размер сообщения gRPC
            ('grpc.max_send_message_length', 1024 * 1024 * 128),
            ('grpc.max_receive_message_length', 1024 * 1024 * 128),
        ]
    )
    # регистрируем наш сервис в gRPC сервере
    llm_service_pb2_grpc.add_LLMServiceServicer_to_server(
        LLMService(runner), server)
    # слушаем все интерфейсы на заданном порте
    server.add_insecure_port(f"[::]:{GRPC_PORT}")
    logger.info("gRPC запущен на порте %s", GRPC_PORT)
    server.start()  # запускаем сервер
    server.wait_for_termination()  # держим процесс живым, пока сервер работает
    logger.info("gRPC сервер остановлен")
