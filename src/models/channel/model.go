package channel

import (
	"database/sql"
	"gochat/src/connections/database"
	"gochat/src/models/channel/userChannelRelationship"
)

type ChannelInterface interface {
	Get() Channel
	Create() error
	Update() error
	Delete() error

	IsPublic() bool
	Authenticate() bool
	Exists() (bool, error)
	HasMember() (bool, error)
}

type Channel struct {
	Id          int
	WorkspaceId int
	Name        string
	Password    string
	Creator     int
}

func GetChannelsByWorkspaceId(workspaceId int) ([]Channel, error) {

	channels := make([]Channel, 0)

	conn := database.GetConnection()
	defer conn.Close()

	query := `
	SELECT id, workspace_id, name, password, creator 
	FROM channels 
	WHERE workspace_id = ?`

	results, err := (*conn).Query(query, workspaceId)

	if err != nil {
		return channels, err
	}

	for results.Next() {
		var channel Channel
		var password sql.NullString
		if err := results.Scan(&channel.Id, &channel.WorkspaceId, &channel.Name, &password, &channel.Creator); err == nil {
			// Handle null password
			if password.Valid {
				channel.Password = password.String
			}
			channels = append(channels, channel)
		} else {
			println(err.Error())
		}
	}
	return channels, nil
}

func (channel Channel) Create() error {

	conn := database.GetConnection()
	defer conn.Close()

	query := `
	INSERT INTO channels (workspace_id, name, password, creator) 
	VALUES (?, ?, ?, ?)`

	res, err := (*conn).Exec(query, channel.WorkspaceId, channel.Name, channel.Password, channel.Creator)

	if err != nil {
		return err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return err
	}

	// create channel member
	relationship := userChannelRelationship.UserChannelRelationship{UserId: channel.Creator, ChannelId: int(id)}
	return relationship.Create()
}
