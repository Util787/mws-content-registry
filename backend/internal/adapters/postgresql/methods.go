package postgresql

import (
	"context"
	"fmt"

	"github.com/Util787/mws-content-registry/internal/models"
	"github.com/jackc/pgx/v5"
)

func (s *Storage) SaveMessage(mes models.Message) error {
	req := `INSERT INTO chats(chat_id, is_user, message, created_at)
						VALUES($1, $2, $3, $4)`

	_, err := s.Db.Exec(context.Background(), req, mes.ChatId, mes.IsUser, mes.Message, mes.CreatedAt)
	if err != nil {
		fmt.Println("Insert error: ", err)
		return err
	}
	return nil
}

func (s *Storage) GetChatHistory(chat int64) ([]models.Message, error) {
	req := `SELECT id, chat_id, is_user, message, created_at
			FROM chats
			WHERE chat_id = $1
			ORDER BY created_at DESC` // desc because its unix int64 timestamp

	rows, err := s.Db.Query(context.Background(), req, chat)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, err
		} else {
			return nil, err
		}
	}
	defer rows.Close()

	var messages []models.Message
	for rows.Next() {
		var message models.Message
		if err := rows.Scan(&message.Id, &message.ChatId, &message.IsUser, &message.Message, &message.CreatedAt); err != nil {
			return nil, err
		}
		messages = append(messages, message)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return messages, nil
}
