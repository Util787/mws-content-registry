from .setup_llm import ensure_environment
from .server import serve

if __name__ == "__main__":
    ensure_environment()  # подготовка среды
    serve()               # запуск gRPC сервера
