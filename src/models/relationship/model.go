package relationship

import "gochat/src/connections/database"

type RelationshipInterface interface {
	Create() error
}

type UserWorkspaceRelationship struct {
	Id 				int
	UserId 			int
	WorkspaceId 	int
}

func (relationship UserWorkspaceRelationship) Create() error {
	conn := database.GetConnection()
	defer conn.Close()

	_, err := (*conn).Exec("INSERT INTO user_workspace (user_id, workspace_id) VALUES (?, ?)", relationship.UserId, relationship.WorkspaceId)
	if err != nil {
		return err
	}

	return nil
}

