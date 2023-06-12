package workspace

import (
	"database/sql"
	"errors"
	"gochat/src/connections/database"
	"gochat/src/helpers/common"
	"gochat/src/models/relationship"
)

type WorkspaceInterface interface {
	GetId() int
	Create() error
	Update() error
	Delete() error

	IsPublic() bool
	Authenticate() bool
	Exists() (bool, error)
	HasMember() (bool, error)
}

type Workspace struct {
	Id          int
	WorkflowKey string
	Name        string
	Password    string
	Creator     int
}

func GetWorkspaceByKey(key string) (Workspace, error) {
	conn := database.GetConnection()
	defer conn.Close()

	result := (*conn).QueryRow("SELECT id, name, creator, password FROM workspaces WHERE workflow_key = ?", key)

	var id int
	var name string
	var creator int
	var password string

	err := result.Scan(&id, &name, &creator, &password)
	if err != nil {
		return Workspace{}, err
	}

	return Workspace{id, key, name, password, creator}, nil
}

func Get() []Workspace {
	workspaces := make([]Workspace, 0)

	conn := database.GetConnection()
	defer conn.Close()

	results, _ := (*conn).Query("SELECT workflow_key, name, password FROM workspaces")

	for results.Next() {
		var workspace Workspace
		results.Scan(&workspace.WorkflowKey, &workspace.Name, &workspace.Password)
		workspaces = append(workspaces, workspace)
	}

	return workspaces
}

// Manejo de workspaces
func (workspace Workspace) GetId() int {
	return workspace.Id
}

func (workspace Workspace) Create(userId int) error {
	conn := database.GetConnection()
	defer conn.Close()

	// Inserto el workspace
	stmt, _ := (*conn).Prepare("INSERT INTO workspaces (workflow_key, name, password, creator) VALUES (?, ?, ?, ?)")
	res, err := stmt.Exec(common.CreateKey(), workspace.Name, workspace.Password, userId)
	if err != nil {
		return err
	}

	workspaceId, err := res.LastInsertId()
	if err != nil {
		return err
	}

	// Creo la relacion entre el usuario y el workspace
	userWorkflow := relationship.UserWorkspaceRelationship{UserId: userId, WorkspaceId: int(workspaceId)}
	return userWorkflow.Create()
}

func (workspace Workspace) Update(userId int, newWorkspace Workspace) error {
	conn := database.GetConnection()
	defer conn.Close()

	_, err := (*conn).Exec("UPDATE workspaces SET name = ?, password = ? WHERE workflow_key = ?", newWorkspace.Name, newWorkspace.Password, workspace.WorkflowKey)
	if err != nil {
		return err
	}

	return nil
}

func (workspace Workspace) Delete(userId int) error {
	conn := database.GetConnection()
	defer conn.Close()

	_, err := (*conn).Exec("DELETE FROM workspaces WHERE workflow_key = ?", workspace.WorkflowKey)
	if err != nil {
		return err
	}

	return nil
}

func (workspace Workspace) IsPublic() bool {
	return workspace.Password == ""
}

func (workspace Workspace) Authenticate(newWorkspace Workspace) bool {
	return workspace.Password == newWorkspace.Password
}

func (workspace Workspace) Exists() (bool, error) {
	ws, err := GetWorkspaceByKey(workspace.WorkflowKey)
	return ws.Id > 0, err
}

func (workspace Workspace) HasMember(userId int) (bool, error) {
	conn := database.GetConnection()
	defer conn.Close()

	// print userId, workspace.Id
	println(userId, workspace.Id)

	result := (*conn).QueryRow("SELECT id FROM user_workspace WHERE user_id = ? AND workspace_id = ?", userId, workspace.Id)

	var id int
	err := result.Scan(&id)

	// Dont return error if no rows found
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return false, err
	}

	return id > 0, nil
}
