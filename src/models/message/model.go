package message

import "gochat/src/connections/database"

type Message struct {
	Id      int    `json:"id"`
	ChatId  int    `json:"chatId"`
	Message string `json:"message"`
	UserId  int    `json:"userId"`
}

type MessageInterface interface {
	Get() (Message, error)
	Save() error
	Update() error
	Delete() error
}

func (message Message) Get() (Message, error) {
	return message, nil
}

func (message Message) Save() error {
	conn := database.GetConnection()
	defer conn.Close()

	query := `
	INSERT INTO chat_messages (chat_id, user_id, message) 
	VALUES (?, ?, ?)`

	_, err := (*conn).Exec(query, message.ChatId, message.UserId, message.Message)

	return err
}
