"""
Подготовка окружения для локальной LLM (Ollama).

- проверяет/устанавливает зависимости из requirements.txt
- проверяет наличие Ollama CLI
- (опционально) запускает ollama serve в detached режиме
- проверяет и загружает модель
- устанавливает Ollama если отсутствует
- устанавливает Python-библиотеки из requirements.txt
"""

import shutil
import platform
import subprocess
import sys
import time
import requests
import ollama
from .config import OLLAMA_MODEL
from .logger import get_logger
from pathlib import Path

logger = get_logger("setup_env")  # логгер для процесса настройки окружения

# ---------------------------------------
# Установка зависимостей из requirements.txt
# ---------------------------------------


def install_requirements():
    """
    Устанавливает все зависимости из requirements.txt через pip.
    Если возникнут ошибки, процесс завершится с исключением.
    """
    req_file = Path(__file__).parent / "requirements.txt"
    if not req_file.exists():
        logger.warning(
            "requirements.txt не найден, пропускаем установку зависимостей")
        return

    logger.info(f"Устанавливаем зависимости из {req_file} ...")
    try:
        subprocess.run([sys.executable, "-m", "pip", "install",
                       "-r", str(req_file)], check=True)
        logger.info("Зависимости успешно установлены")
    except subprocess.CalledProcessError as e:
        logger.error(f"Ошибка при установке зависимостей: {e}")
        raise RuntimeError(
            "Не удалось установить зависимости из requirements.txt") from e


# ---------------------------------------
#  Проверка, работает ли Ollama сервер
# ---------------------------------------
def ollama_is_running() -> bool:
    """Проверяет доступность локального Ollama сервера через HTTP запрос."""
    try:
        r = requests.get("http://127.0.0.1:11434/api/tags", timeout=2)
        return r.status_code == 200
    except Exception:
        return False


# ---------------------------------------
#  Установка Ollama если отсутствует
# ---------------------------------------
def install_ollama_if_missing():
    """Автоматическая установка Ollama для macOS или Linux, если CLI не найден."""
    if shutil.which("ollama"):
        logger.info("Ollama CLI найден, установка не требуется")
        return

    system = platform.system()
    logger.info(f"Ollama CLI не найден, пробуем установить для {system}...")

    try:
        if system == "Darwin":
            subprocess.run(["brew", "install", "ollama"], check=True)
        elif system == "Linux":
            url = "https://ollama.com/download/ollama-latest-linux.tar.gz"
            subprocess.run(
                ["curl", "-L", url, "-o", "/tmp/ollama.tar.gz"], check=True)
            subprocess.run(
                ["tar", "-xzf", "/tmp/ollama.tar.gz", "-C", "/tmp/"], check=True)
            subprocess.run(["sudo", "mv", "/tmp/ollama",
                           "/usr/local/bin/ollama"], check=True)
        else:
            raise RuntimeError(
                f"Автоматическая установка Ollama не поддерживается для {system}")

        logger.info("Ollama успешно установлен")
    except subprocess.CalledProcessError as e:
        logger.error(f"Ошибка при установке Ollama: {e}")
        raise RuntimeError("Не удалось установить Ollama автоматически") from e


# ---------------------------------------
#  Запуск ollama serve
# ---------------------------------------
def start_ollama_server(timeout: int = 25):
    """Фоновый запуск Ollama сервера."""
    if ollama_is_running():
        logger.info("Ollama сервер уже запущен")
        return

    logger.info("Ollama не запущена. Стартуем ollama serve...")

    proc = subprocess.Popen(
        ["ollama", "serve"],
        stdout=subprocess.PIPE,
        stderr=subprocess.PIPE
    )

    start = time.time()
    while time.time() - start < timeout:
        if ollama_is_running():
            logger.info("Ollama успешно запущена")
            return
        time.sleep(0.5)

    proc.terminate()
    raise RuntimeError("Ollama не была запущена в отведённое время")


# ---------------------------------------
#  Проверка наличия модели
# ---------------------------------------
def ensure_model():
    """Проверяет, установлена ли нужная модель в Ollama. Скачивает при отсутствии."""
    models = ollama.list()
    if hasattr(models, "models"):
        models = models.models

    names = [m.get("name") if isinstance(m, dict) else getattr(m, "model", None)
             for m in models]

    if OLLAMA_MODEL not in names:
        logger.info(f"Модель {OLLAMA_MODEL} отсутствует. Скачиваем...")
        ollama.pull(OLLAMA_MODEL)
    else:
        logger.info(f"Модель {OLLAMA_MODEL} уже установлена")


# ---------------------------------------
#  Основная функция подготовки среды
# ---------------------------------------
def ensure_environment():
    """Полная подготовка локальной среды."""
    if platform.system() not in ("Darwin", "Linux"):
        logger.error("Ollama автоматом запускается только на macOS/Linux")
        return

    install_requirements()       # установка Python-зависимостей
    install_ollama_if_missing()  # установка Ollama если нет
    start_ollama_server()        # старт сервера в фоне
    ensure_model()               # проверка и установка модели

    logger.info("Среда готова к работе")


if __name__ == "__main__":
    ensure_environment()
