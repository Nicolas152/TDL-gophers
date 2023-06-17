package userChannelRelationship

import "gochat/src/connections/database"

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
