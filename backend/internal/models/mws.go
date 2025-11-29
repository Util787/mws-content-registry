package models

// Response — общий ответ MWS Tables API
type Response struct {
	Code    int    `json:"code"`
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    Data   `json:"data"`
}

// Data — контейнер для записей и информации о пагинации
type Data struct {
	Records  []Record `json:"records"`
	PageNum  int      `json:"pageNum"`
	PageSize int      `json:"pageSize"`
	Total    int      `json:"total"`
}

// Record — отдельная запись в таблице
type Record struct {
	RecordID string `json:"recordId"`
	Fields   Fields `json:"fields"`
}

// Fields — поля записи таблицы
type Fields struct {
	ID              int      `json:"id"`              
	URL             string   `json:"url"`             
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
	AIAnalyzeDate   int64    `json:"ai_analyze_date"` 
}
