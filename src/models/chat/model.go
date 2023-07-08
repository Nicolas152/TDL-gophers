package chat

import (
	"database/sql"
	"errors"
	"gochat/src/connections/database"
)

type Chat struct {
	Id        int
	ChannelId int
	DmId      int
}

func Create(channelId int, dmId int) (int, error) {
	conn := database.GetConnection()
	defer conn.Close()

	// Verificar si channelId o dmId son igual a cero y asignar NULL en su lugar
	var channelIDValue, dmIDValue interface{}
	if channelId == 0 {
		channelIDValue = nil // Asignar NULL
	} else {
		channelIDValue = channelId
	}

	if dmId == 0 {
		dmIDValue = nil // Asignar NULL
	} else {
		dmIDValue = dmId
	}

	query := `
	INSERT INTO chats (channel_id, dm_id) 
	VALUES (?, ?)`

	result, err := (*conn).Exec(query, channelIDValue, dmIDValue)

	if err != nil {
		return 0, err
	}

	id, _ := result.LastInsertId()

	return int(id), nil
}

func Resolve(channelId int, dmId int) (int, error) {
	conn := database.GetConnection()
	defer conn.Close()

	// Verificar si channelId o dmId son igual a cero y asignar NULL en su lugar
	var query string
	var row *sql.Row
	if channelId != 0 {
		query = ` SELECT id FROM chats 
				  WHERE channel_id = ? AND dm_id IS NULL`
		row = (*conn).QueryRow(query, channelId)
	} else if dmId != 0 {
		query = ` SELECT id FROM chats 
		WHERE channel_id is NULL AND dm_id = ?`
		row = (*conn).QueryRow(query, dmId)
	} else {
		return 0, errors.New("Does not exist a chat with the given parameters")
	}

	var id int
	err := row.Scan(&id)

	if err != nil {
		return 0, err
	}

	return id, nil
}
