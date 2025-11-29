package postgresql

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

type Message struct {
	Id      int64
	ChatNum int64
	IsUser  bool
	Message string
}

func (s *Storage) AddMessage(mes Message) error {
	req := `INSERT INTO chats(chat_num, is_user, message)
						VALUES($1, $2, $3)`

	_, err := s.Db.Exec(context.Background(), req, mes.ChatNum, mes.IsUser, mes.Message)
	if err != nil {
		fmt.Println("Insert error: ", err)
		return err
	}
	return nil
}

func (s *Storage) TakeByChatNum(chat int64) ([]Message, error) {
	req := `SELECT id, chat_num, is_user, message
			FROM chats
			WHERE chat_num = $1
			ORDER BY id`

	rows, err := s.Db.Query(context.Background(), req, chat)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, err
		} else {
			return nil, err
		}
	}
	defer rows.Close()

	var messages []Message
	for rows.Next() {
		var message Message
		if err := rows.Scan(&message.Id, &message.ChatNum, &message.IsUser, &message.Message); err != nil {
			return nil, err
		}
		messages = append(messages, message)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return messages, nil
}
