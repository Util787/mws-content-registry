# MWS-CONTENT-REGISTRY
Это программа собирающая аналитические данные из различных api (на данный момент поддерживает только youtube) и добавляет их в таблицу MWSTables, также поддерживает функции AI: чат-бота(учитывает контекст MWSTables), анализа данных, используя LLM API, но также готова локальная модель для работы по grpc

# Структура проекта
- Backend: Golang, gin
- Frontend: React + Vite
- LLM: Python, Ollama

# Quick start

1. **Clone the repository**:
   ```bash
   git clone https://github.com/Util787/mws-content-registry
   cd mws-content-registry/
   ```

2. **Create and configure the `.env` file**:
   ```bash
   cp .env.example .env
   ```

2. **Copy `.env` file to backend/**:
   ```bash
   cp .env ./backend/
   ```

3. **Build and run compose**:
   ```bash
   docker-compose up --build
   ```

# Команда Bezrabotine
- Рябов Влад - backend
- Чуприна Георгий - backend
- Дегтярев Александр - da/ml/backend
- Шокота Даниил - frontend

# Дополнительно
- ссылка на MWSTables: https://tables.mws.ru/share/shrKQgGkVPUXjeRBnuyGF