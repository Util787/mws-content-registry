# ===========================
# Заглушка для модели Ollama
# ===========================
import ollama
from .config import OLLAMA_MODEL


class ModelRunner:
    def __init__(self, model_name=OLLAMA_MODEL):
        self.model_name = model_name

    def process(self, user_query, stats_json, context=""):
        """
        Формирует prompt из пользовательского запроса, статистики и контекста,
        и передаёт его в Ollama для генерации ответа.
        """
        prompt = f"User query: {user_query}\nStats: {stats_json}\nContext: {context}"

        raw_response = ollama.generate(model=self.model_name, prompt=prompt)

        result = {
            # или распарсить JSON, если модель возвращает JSON-строку
            "answer": str(raw_response),
            "model": self.model_name,
        }

        return result
