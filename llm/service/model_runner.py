import time
import requests
import ollama
from .config import OLLAMA_MODEL
from .logger import get_logger

# логгер для отслеживания работы ModelRunner
logger = get_logger("ModelRunner")


def is_ollama_up(timeout=1) -> bool:
    """
    Проверяет, запущен ли локальный Ollama сервер.
    Делает HTTP GET-запрос к /api/tags, ожидает статус 200.
    """
    try:
        r = requests.get("http://localhost:11434/api/tags", timeout=timeout)
        return r.status_code == 200
    except Exception:
        return False


def wait_for_ollama(timeout=10, interval=0.2):
    """
    Ждём, пока Ollama сервер станет доступен.

    :param timeout: максимальное время ожидания в секундах
    :param interval: интервал между проверками
    """
    start = time.time()
    while time.time() - start < timeout:
        if is_ollama_up(timeout=1):
            return True  # сервер доступен, можно продолжать
        time.sleep(interval)  # ждём перед следующей проверкой
    # превышено время ожидания
    raise RuntimeError("Ollama не стартовала в отведённое время")


class ModelRunner:
    """
    Обёртка для работы с моделью Ollama.
    Позволяет формировать prompt и получать ответ модели.
    """

    def __init__(self, model_name: str = OLLAMA_MODEL, ensure_ollama: bool = False, wait_timeout: int = 10):
        self.model_name = model_name
        # если ensure_ollama=True, ждём пока сервер Ollama запустится перед созданием объекта
        if ensure_ollama:
            wait_for_ollama(timeout=wait_timeout)

    def process(self, user_query: str, stats_json: str, context: str = "") -> dict:
        """
        Отправляет запрос в модель и возвращает ответ.

        :param user_query: текстовый запрос пользователя
        :param stats_json: JSON-строка с дополнительной статистикой
        :param context: контекст для уточнения ответа модели
        :return: словарь с ключами 'answer' и 'model'
        """
        prompt = f"User query: {user_query}\nStats: {stats_json}\nContext: {context}"
        logger.debug("Отправляем порт в модель: %s", prompt)

        last_exc = None
        # делаем несколько попыток вызова модели на случай временной ошибки
        for attempt in range(3):
            try:
                raw = ollama.generate(
                    model=self.model_name, prompt=prompt, stream=False)  # вызываем модель без стриминга
                logger.debug("Получен сырой ответ: %s", raw)
                # возвращаем результат
                return {"answer": str(raw), "model": self.model_name}
            except Exception as e:
                last_exc = e
                logger.warning(
                    "Ollama генерация провелилась (попытка %d): %s", attempt + 1, e)
                time.sleep(0.5)  # небольшой перерыв перед следующей попыткой

        # если все попытки не удались — логируем и поднимаем исключение
        logger.error("Все попытки вызова Ollama провалились: %s", last_exc)
        raise RuntimeError("Ошибка получения ответа от Ollama") from last_exc
