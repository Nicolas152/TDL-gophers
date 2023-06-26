package message

import "gochat/src/connections/database"

type Message struct {
	Id           int    `json:"id"`
	WorkspaceKey string `json:"workspaceKey"`
	ChannelKey   int    `json:"channelKey"`
	Message      string `json:"message"`
	UserId       int    `json:"userId"`
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
	INSERT INTO channel_messages (channel_id, user_id, message) 
	VALUES (?, ?, ?)`

	_, err := (*conn).Exec(query, message.ChannelKey, message.UserId, message.Message)

	return err
}
