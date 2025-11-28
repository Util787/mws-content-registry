"""
Этот скрипт отвечает за подготовку окружения для работы локальной LLM.

- проверяет и устанавливает Python-зависимости из requirements.txt;
- убеждается, что Ollama установлена (для macOS и Linux может поставить сам);
- поднимает локальный ollama-server, если он ещё не запущен;
- проверяет, скачана ли нужная модель, и при необходимости загружает её.

По сути, после запуска этого файла среда полностью готова к использованию модели. В случае с Windows пользователю будет выведена инструкция по установке Ollama вручную.
"""


import sys
import os
import platform
import subprocess
import time
import ollama
from .config import OLLAMA_MODEL


def check_python_requirements():
    # Папка, где находится setup_env.py
    base_dir = os.path.dirname(os.path.abspath(__file__))
    req_file = os.path.join(base_dir, "requirements.txt")

    if os.path.exists(req_file):
        print("=== Проверка и установка зависимостей Python ===")
        subprocess.run([sys.executable, "-m", "pip",
                       "install", "--upgrade", "pip"])
        subprocess.run([sys.executable, "-m", "pip",
                       "install", "-r", req_file])
    else:
        print(f"requirements.txt не найден! Ожидалось здесь: {req_file}")


def check_ollama_installed():
    try:
        subprocess.run(["ollama", "--version"], check=True,
                       stdout=subprocess.DEVNULL)
        return True
    except (FileNotFoundError, subprocess.CalledProcessError):
        return False


def install_ollama_unix():
    system = platform.system()
    if system == "Darwin":
        print("=== Установка Ollama через Homebrew ===")
        if not subprocess.run(["brew", "--version"], stdout=subprocess.DEVNULL).returncode == 0:
            print("Homebrew не найден, устанавливаем...")
            subprocess.run(
                ["/bin/bash", "-c",
                    "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"],
                check=True
            )
        subprocess.run(["brew", "install", "--cask", "ollama"], check=True)
    elif system == "Linux":
        print("=== Установка Ollama через официальный скрипт ===")
        subprocess.run(
            "curl -fsSL https://ollama.com/install.sh | sh",
            shell=True,
            check=True
        )


def setup_ollama_server(timeout: int = 30, interval: float = 1.0):
    """
    Запускает Ollama сервер, если он не работает.
    Ждёт, пока сервер станет доступен, или пока не истечёт timeout.

    :param timeout: максимальное время ожидания (в секундах)
    :param interval: интервал между проверками (в секундах)
    """
    import requests
    import subprocess
    import time

    def server_is_up() -> bool:
        try:
            r = requests.get("http://localhost:11434/api/tags", timeout=3)
            return r.status_code == 200
        except requests.RequestException:
            return False

    if server_is_up():
        print("Ollama сервер уже работает")
        return

    print("Сервер Ollama не отвечает, пробуем запустить...")
    proc = subprocess.Popen(["ollama", "serve"],
                            stdout=subprocess.PIPE, stderr=subprocess.PIPE)

    start_time = time.time()
    while time.time() - start_time < timeout:
        if server_is_up():
            print("Ollama сервер успешно запущен")
            return
        time.sleep(interval)

    # Если по таймауту сервер не поднялся
    proc.terminate()
    raise RuntimeError(f"Ollama сервер не запустился за {timeout} секунд")


def pull_model():
    result = ollama.list()  # возвращает ListResponse

    # result.models — список объектов модели
    try:
        models_data = result.models
    except AttributeError:
        models_data = result  # fallback, если нет .models

    installed_models = []
    for m in models_data:
        if hasattr(m, "model"):
            installed_models.append(m.model)
        else:
            print("Не удалось определить имя модели для элемента:", m)

    if OLLAMA_MODEL not in installed_models:
        print(f"Модель {OLLAMA_MODEL} не найдена! Скачиваем...")
        ollama.pull(OLLAMA_MODEL)
    else:
        print(f"Модель {OLLAMA_MODEL} уже установлена")


def windows_install_instructions():
    print("""
          Ollama не установлена на Windows.
          1. Установите через winget (PowerShell от имени администратора):
          winget install Ollama.Ollama
          2. Или скачайте и установите вручную: https://ollama.com/download
          После установки перезапустите этот скрипт.
          """)


def ensure_environment():
    check_python_requirements()

    system = platform.system()
    if system in ["Darwin", "Linux"]:
        if not check_ollama_installed():
            install_ollama_unix()
        setup_ollama_server()
        pull_model()
    elif system.startswith("MINGW") or system == "Windows" or system.startswith("CYGWIN"):
        if not check_ollama_installed():
            windows_install_instructions()
            sys.exit(1)
        setup_ollama_server()
        pull_model()
    else:
        print(f"Неизвестная ОС: {system}")
        sys.exit(1)

    print("=== Среда готова к работе ===")


if __name__ == "__main__":
    ensure_environment()
