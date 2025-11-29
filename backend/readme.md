# Quick start

1. **Clone the repository**:
   ```bash
   git clone https://github.com/Util787/mws-content-registry
   cd mws-content-registry/backend
   ```

2. **Create and configure the `.env` file**:
   ```bash
   cp .env.example .env
   ```

3. **Build and run dockerfile**:
   ```bash
   docker build -t mcr-backend .
   ```
    ```bash
   docker run -p 8000:8000 mcr-backend
   ```

## API Документация

### 1. **Добавить YouTube видео по URL**

**POST** `/api/v1/add-yt-video`

**Тело запроса:**

```json
{
  "url": "https://www.youtube.com/watch?v=some_video_id"
}
```

**Ответ:**

- Успех: `200 OK`
  ```json
  {
    "message": "YouTube video added successfully"
  }
  ```
- Ошибка: `400 Bad Request` или `500 Internal Server Error`

---

### 2. **Добавить последние популярные YouTube видео**

**POST** `/api/v1/add-yt-videos/recent`

**Ответ:**

- Успех: `200 OK`
  ```json
  {
    "message": "Recent YouTube videos added successfully"
  }
  ```
- Ошибка: `500 Internal Server Error`

---

### 3. **Анализировать контент с помощью LLM**

**POST** `/api/v1/add-llm-analyze/:recordId`

**Path parameter:**

- `recordId` — ID записи для анализа.

**Ответ:**

- Успех: `200 OK`
  ```json
  {
    "message": "content analyzed by llm successfully"
  }
  ```
- Ошибка: `400 Bad Request` или `500 Internal Server Error`

---

### 4. **Получить записи таблицы**

**GET** `/api/v1/records`

**Query parameters:**

| Параметр     | Тип                | Обязательный | Описание                          |
|--------------|--------------------|--------------|-----------------------------------|
| `pageNum`    | `int`              | Да           | Номер страницы.                  |
| `pageSize`   | `int`              | Да           | Размер страницы.                 |
| `sort`       | `[{ field: 'field1', order: 'asc' }]` | Нет         | Сортировка записей.              |
| `recordId`   | `string`           | Нет          | Фильтр по ID записи.             |
| `fields`     | `[]string`         | Нет          | Поля для выборки.                |

**Пример запроса:**

```http
GET /api/v1/records?pageNum=1&pageSize=10
```

**Ответ:**

- Успех: `200 OK`
  ```json
  {
    "records": [
      {
        "recordId": "123",
        "fields": {
          "id": 1,
          "url": "https://example.com",
          "published_at": 1633024800,
          "views": 1000,
          "topic": "Example Topic",
          "likes": 100,
          "comments": "Example comment",
          "comments_count": 10,
          "description": "Example description",
          "author": "Author Name",
          "recomendations": "Example recommendations",
          "comments_summary": "Summary",
          "comments_tone": "Positive",
          "ai_analyze_date": 1633024800
        }
      }
    ]
  }
  ```
- Ошибка: `400 Bad Request` или `500 Internal Server Error`

---
