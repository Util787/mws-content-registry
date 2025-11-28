# ===========================
# Заглушка для модели Ollama
# ===========================
import ollama
from config import OLLAMA_MODEL


class ModelRunner:
    def __init__(self):
        self.model_name = OLLAMA_MODEL

    def process(self, user_query, stats_json, context):
        # Формирование запроса к модели
        prompt = f"User query: {user_query}\nStats: {stats_json}\nContext: {context}"
        response = ollama.generate(
            model=self.model_name, prompt=prompt, max_tokens=500)
        return {
            "answer_text": response.text,
            "reasoning": "",
            "structured": ""
        }
