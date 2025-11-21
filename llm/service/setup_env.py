import sys
import os
import platform
import subprocess
import time

MODEL_NAME = "qwen2.5:7b"


def check_python_requirements():
    if os.path.exists("requirements.txt"):
        print("=== Проверка и установка зависимостей Python ===")
        subprocess.run([sys.executable, "-m", "pip",
                       "install", "--upgrade", "pip"])
        subprocess.run([sys.executable, "-m", "pip",
                       "install", "-r", "requirements.txt"])
    else:
        print("requirements.txt не найден!")


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


def setup_ollama_server():
    import requests
    try:
        r = requests.get("http://localhost:11434/api/tags", timeout=3)
        if r.status_code == 200:
            print("Ollama сервер работает")
            return
    except requests.RequestException:
        pass

    print("Сервер Ollama не отвечает. Запускаем...")
    subprocess.Popen(["ollama", "serve"])
    time.sleep(5)


def pull_model():
    import ollama
    installed_models = [m["name"] for m in ollama.list()]
    if MODEL_NAME not in installed_models:
        print(f"Модель {MODEL_NAME} не найдена! Скачиваем...")
        ollama.pull(MODEL_NAME)
    else:
        print(f"Модель {MODEL_NAME} уже установлена")


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
