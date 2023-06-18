package channel

import (
	"database/sql"
	"gochat/src/connections/database"
	"gochat/src/models/channel/userChannelRelationship"
)

type ChannelInterface interface {
	Get() (Channel, error)
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

// Channel Model to be returned to client
type ClientChannel struct {
	Id          string
	Name        string
	WorkspaceId int
	Creator     int
}

func (channel Channel) Get() (Channel, error) {

	conn := database.GetConnection()
	defer conn.Close()

	query := `
	SELECT id, workspace_id, name, password, creator 
	FROM channels 
	WHERE id = ? AND workspace_id = ?`

	var password sql.NullString
	if err := (*conn).QueryRow(query, channel.Id, channel.WorkspaceId).Scan(&channel.Id, &channel.WorkspaceId, &channel.Name, &password, &channel.Creator); err != nil {
		return channel, err
	}

	// Handle null password
	if password.Valid {
		channel.Password = password.String
	}

	return channel, nil
}

func GetChannelsByWorkspaceId(workspaceId int) ([]ClientChannel, error) {

	channels := make([]ClientChannel, 0)

	conn := database.GetConnection()
	defer conn.Close()

	query := `
	SELECT id, workspace_id, name, creator 
	FROM channels 
	WHERE workspace_id = ?`

	results, err := (*conn).Query(query, workspaceId)

	if err != nil {
		return channels, err
	}

	for results.Next() {
		var channel ClientChannel
		if err := results.Scan(&channel.Id, &channel.WorkspaceId, &channel.Name, &channel.Creator); err == nil {
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

func (channel Channel) Update() error {

	conn := database.GetConnection()
	defer conn.Close()

	query := `
	UPDATE channels 
	SET name = ?, password = ? 
	WHERE id = ?`

	_, err := (*conn).Exec(query, channel.Name, channel.Password, channel.Id)

	if err != nil {
		return err
	}

	return nil
}

func (channel Channel) IsOwner(userId int) bool {
	return channel.Creator == userId
}
