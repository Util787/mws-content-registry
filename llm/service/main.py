from service.setup_env import ensure_environment
from service.server import serve

if __name__ == "__main__":
    ensure_environment()  # подготовка среды
    serve()               # запуск gRPC сервера
