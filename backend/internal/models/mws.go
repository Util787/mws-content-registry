package models

// Response — общий ответ MWS Tables API
type MWSTableResponse struct {
	Code    int          `json:"code"`
	Success bool         `json:"success"`
	Message string       `json:"message"`
	Data    MWSTableData `json:"data"`
}

// Data — контейнер для записей и информации о пагинации
type MWSTableData struct {
	Records  []MWSTableRecord `json:"records"`
	PageNum  int              `json:"pageNum"`
	PageSize int              `json:"pageSize"`
	Total    int              `json:"total"`
}

// Record — отдельная запись в таблице
type MWSTableRecord struct {
	RecordID string                  `json:"recordId"`
	Fields   MWSTableGetRecordFields `json:"fields"`
}

// Fields — поля записи таблицы
type MWSTableGetRecordFields struct {
	ID              int      `json:"id"`
	URL             string   `json:"url.text"`
	PublishedAt     string   `json:"published_at"`
	Views           int      `json:"views"`
	Topic           string   `json:"topic"`
	Likes           int      `json:"likes"`
	Comments        []string `json:"comments"`
	Description     string   `json:"description"`
	Author          string   `json:"author"`
	Recomendations  string   `json:"recomendations"`
	CommentsSummary string   `json:"comments_summary"`
	CommentsTone    string   `json:"comments_tone"`
	AiAnalyzeDate   int      `json:"ai_analyze_date"`
}

type MWSTableAddRecordFields struct {
	URL             *string   `json:"url"`
	PublishedAt     *string   `json:"published_at"`
	Views           *int      `json:"views"`
	Topic           *string   `json:"topic"`
	Likes           *int      `json:"likes"`
	Comments        *[]string `json:"comments"`
	Description     *string   `json:"description"`
	Author          *string   `json:"author"`
	Recomendations  *string   `json:"recomendations"`
	CommentsSummary *string   `json:"comments_summary"`
	CommentsTone    *string   `json:"comments_tone"`
}

type MWSTableNewRecord struct {
	Fields MWSTableAddRecordFields `json:"fields"`
}
