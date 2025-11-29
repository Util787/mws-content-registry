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
	ID              int    `json:"id"`
	URL             string `json:"url"`
	PublishedAt     int64  `json:"published_at"`
	Views           uint64 `json:"views"`
	Topic           string `json:"topic"`
	Likes           uint64 `json:"likes"`
	Comments        string `json:"comments"`
	CommentsCount   uint64 `json:"comments_count"`
	Description     string `json:"description"`
	Author          string `json:"author"`
	Recomendations  string `json:"recomendations"`
	CommentsSummary string `json:"comments_summary"`
	CommentsTone    string `json:"comments_tone"`
	AiAnalyzeDate   int    `json:"ai_analyze_date"`
}

type MWSTableAddRecordFields struct {
	URL             *string `json:"url,omitempty"`
	PublishedAt     *int64  `json:"published_at,omitempty"`
	Views           *uint64 `json:"views,omitempty"`
	Topic           *string `json:"topic,omitempty"`
	Likes           *uint64 `json:"likes,omitempty"`
	Comments        *string `json:"comments,omitempty"`
	CommentsCount   *uint64 `json:"comments_count,omitempty"`
	Description     *string `json:"description,omitempty"`
	Author          *string `json:"author,omitempty"`
	Recomendations  *string `json:"recomendations,omitempty"`
	CommentsSummary *string `json:"comments_summary,omitempty"`
	CommentsTone    *string `json:"comments_tone,omitempty"`
	AiAnalyzeDate   *int    `json:"ai_analyze_date,omitempty"`
}

type MWSTableNewRecord struct {
	Fields MWSTableAddRecordFields `json:"fields"`
}

type MWSTableUpdateRecord struct {
	RecordID string                  `json:"recordId"`
	Fields   MWSTableAddRecordFields `json:"fields"`
}
