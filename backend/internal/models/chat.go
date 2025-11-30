package models

type Message struct {
	Id        int64  `json:"id"`
	ChatId    int64  `json:"chat_id"`
	IsUser    bool   `json:"is_user"`
	Message   string `json:"message"`
	CreatedAt int64  `json:"created_at"`
}
