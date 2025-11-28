#!/usr/bin/env python3
"""
Создание виртуального окружения Python 3.12 и установка зависимостей.
- Создаёт папку .venv
- Устанавливает pip/setuptools/wheel
- Устанавливает зависимости из requirements.txt
"""

import os
import sys
import subprocess
from pathlib import Path

ROOT = Path(__file__).parent.resolve()
VENV_DIR = ROOT / "../.venv"
REQ = ROOT / "requirements.txt"


def run(cmd, **kwargs):
    """Утилита для короткого вызова команд."""
    subprocess.run(cmd, check=True, **kwargs)


def ensure_python312():
    """Проверяет, что доступен python3.12."""
    for name in ("python3.12", "python3", "python"):
        path = shutil.which(name)
        if not path:
            continue
        # проверяем версию
        try:
            out = subprocess.check_output(
                [path, "-c", "import sys; print(sys.version_info[:2])"])
            if b"(3, 12)" in out:
                return path
        except Exception:
            pass
    raise RuntimeError("Python 3.12 не найден. Установите Python 3.12.")


def create_venv(python):
    """Создаёт виртуальное окружение .venv."""
    if VENV_DIR.exists():
        print(".venv уже существует, пропускаю создание.")
        return
    print("Создаю виртуальное окружение .venv ...")
    run([python, "-m", "venv", str(VENV_DIR)])


def install_requirements(python_bin):
    """Установка зависимостей."""
    if not REQ.exists():
        print("requirements.txt не найден — пропускаю установку.")
        return

    print("Обновляю pip...")
    run([python_bin, "-m", "pip", "install",
        "--upgrade", "pip", "setuptools", "wheel"])

    print("Устанавливаю зависимости из requirements.txt ...")
    run([python_bin, "-m", "pip", "install", "-r", str(REQ)])

    print("Готово.")


def main():
    python = ensure_python312()
    create_venv(python)

    # путь к python внутри venv
    if os.name == "nt":
        python_bin = VENV_DIR / "Scripts" / "python.exe"
    else:
        python_bin = VENV_DIR / "bin" / "python"

    install_requirements(str(python_bin))


if __name__ == "__main__":
    import shutil
    main()
