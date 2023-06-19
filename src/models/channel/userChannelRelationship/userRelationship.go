package userChannelRelationship

import (
	"gochat/src/connections/database"
	"gochat/src/models/user"
)

type RelationshipInterface interface {
	Create() error
}

type UserChannelRelationship struct {
	Id        int
	UserId    int
	ChannelId int
}

func (relationship UserChannelRelationship) Create() error {
	conn := database.GetConnection()
	defer conn.Close()

	_, err := (*conn).Exec("INSERT INTO user_channels (user_id, channel_id) VALUES (?, ?)", relationship.UserId, relationship.ChannelId)
	if err != nil {
		return err
	}

	return nil
}

func (relationship UserChannelRelationship) Delete() error {
	conn := database.GetConnection()
	defer conn.Close()

	_, err := (*conn).Exec("DELETE FROM user_channels WHERE user_id = ? AND channel_id = ?", relationship.UserId, relationship.ChannelId)
	if err != nil {
		return err
	}
	return nil
}

func (relationship UserChannelRelationship) Exists() bool {
	conn := database.GetConnection()
	defer conn.Close()

	var count int
	(*conn).QueryRow("SELECT COUNT(*) FROM user_channels WHERE user_id = ? AND channel_id = ?", relationship.UserId, relationship.ChannelId).Scan(&count)

	return count > 0
}

func (relationship UserChannelRelationship) GetMembers() ([]user.UserClient, error) {
	conn := database.GetConnection()
	defer conn.Close()

	query := `
	SELECT users.id, users.name, users.email 
	FROM users
	INNER JOIN user_channels ON users.id = user_channels.user_id
	WHERE user_channels.channel_id = ?`

	results, _ := (*conn).Query(query, relationship.ChannelId)

	users := make([]user.UserClient, 0)

	for results.Next() {
		var user user.UserClient
		if err := results.Scan(&user.Id, &user.Name, &user.Email); err == nil {
			users = append(users, user)
		}
	}

	return users, nil
}
