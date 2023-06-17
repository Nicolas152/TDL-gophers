package channel

import (
	"database/sql"
	"gochat/src/connections/database"
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
	WorkspaceId string
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
