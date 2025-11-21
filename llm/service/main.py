import ollama

response = ollama.generate(
    model="qwen2.5:7b",
    prompt="Перечисли ключевые метрики анализа соцсетей."
)

print(response["response"])
