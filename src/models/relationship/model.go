package relationship

import (
	"errors"
	"gochat/src/connections/database"
)

type RelationshipInterface interface {
	Create() error
}

type UserWorkspaceRelationship struct {
	Id 				int
	UserId 			int
	WorkspaceId 	int
}

func (relationship UserWorkspaceRelationship) Exists() bool {
	conn := database.GetConnection()
	defer conn.Close()

	result := (*conn).QueryRow("SELECT id FROM user_workspace WHERE user_id = ? AND workspace_id = ?", relationship.UserId, relationship.WorkspaceId)

	var id int
	if err := result.Scan(&id); err != nil {
		return false
	}

	return true
}

func (relationship UserWorkspaceRelationship) Create() error {
	if relationship.Exists() {
		return errors.New("Relationship already exists")
	}

	conn := database.GetConnection()
	defer conn.Close()

	_, err := (*conn).Exec("INSERT INTO user_workspace (user_id, workspace_id) VALUES (?, ?)", relationship.UserId, relationship.WorkspaceId)
	if err != nil {
		return err
	}

	return nil
}

