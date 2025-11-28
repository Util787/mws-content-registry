# Проект LLM с Ollama

## Содержание

- `service/model_runner.py` — класс `ModelRunner` для работы с Ollama.
- `service/server.py` — gRPC сервер для вызова модели через `LLMService`.
- `service/setup_env.py` — скрипт для подготовки среды (проверка Ollama, запуск сервера, установка модели и зависимостей).
- `tests/` — тесты: проверка сервера, прямой вызов модели, интеграция и gRPC.
- `llm_grpc/` — сгенерированные gRPC Python-модули (`llm_service_pb2.py`, `llm_service_pb2_grpc.py`).

---

## Установка и подготовка среды

- На macOS и Linux скрипт `setup_env.py` автоматически устанавливает зависимости из `requirements.txt`, проверяет наличие Ollama CLI, устанавливает Ollama при необходимости, запускает `ollama serve` и подтягивает модель `OLLAMA_MODEL`.
- На Windows необходимо установить Ollama вручную.

---

## Запуск gRPC сервера

```bash
python -m service.server
```

По умолчанию gRPC сервер слушает порт, указанный в `service/config.py` (`GRPC_PORT`).

## Использование ModelRunner

```python
from service.model_runner import ModelRunner
from service.config import OLLAMA_MODEL

runner = ModelRunner(model_name=OLLAMA_MODEL)
result = runner.process("Сколько просмотров у поста?", '{"views":123,"likes":10}')
print(result["answer"])
```

Параметры:

- `user_query` — вопрос пользователя.
- `stats_json` — дополнительные данные в формате JSON.
- `context` — контекст запроса (опционально).

## Тесты

Все тесты находятся в папке `tests/` и делятся на четыре категории:

### Server (@pytest.mark.server)

- Проверяет доступность локального Ollama сервера по HTTP.
- Используется, чтобы убедиться, что сервер поднят перед интеграцией.

### Model (@pytest.mark.model)

- Проверяет прямой вызов `ollama.generate`.
- Модель должна возвращать непустой ответ без использования `ModelRunner`.

### Integration (@pytest.mark.integration)

- Полная интеграция через `ModelRunner`.
- Проверяет формирование prompt, вызов модели и корректность ответа.

### gRPC (@pytest.mark.grpc)

- Проверяет работу gRPC сервиса (`LLMService`).
- Отправляет `AnalyzeRequest` и проверяет, что `AnalyzeResponse.answer` не пустой.

## Запуск тестов

Для запуска всех тестов используйте команду:

```bash
pytest -v tests/
```

Эта команда выполнит все тесты с маркерами: `server`, `model`, `integration`, `grpc`.

## Отличия тестовых файлов

| Параметр                  | `test_grpc_service.py`          | `test_model_runner.py`              |
| ------------------------- | ------------------------------- | ----------------------------------- |
| Подъём Ollama сервера     | Да, автоматически через fixture | Нет, сервер должен быть уже запущен |
| Подъём gRPC сервера       | Да, автоматически через fixture | Нет, сервер должен быть уже запущен |
| Зависимость от окружения  | Изолированные тесты             | Зависит от внешнего окружения       |
| Использование ModelRunner | Да                              | Да                                  |
| Тестирование gRPC         | Да                              | Да                                  |
| Уровень тестирования      | e2e / интеграционный            | unit / интеграционный               |
