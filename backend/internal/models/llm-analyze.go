package models

type AnalyzeData struct {
	Topic           string `json:"topic"`
	CommentsSummary string `json:"comments_summary"`
	CommentsTone    string `json:"comments_tone"`
	Analyze         string `json:"analyze"`
}
