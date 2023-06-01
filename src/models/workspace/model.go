package workspace

import (
	"gochat/src/connections/database"
	"gochat/src/helpers/common"
)

type WorkspaceInterface interface {
	Create() error
	Update() error
	Delete() error
}

type Workspace struct {
	Id 				int
	WorkflowKey 	string
	Name 			string
	Password 		string
}


func GetWorkspaceByKey(key string) (Workspace, error) {
	conn := database.GetConnection()
	defer conn.Close()

	result := (*conn).QueryRow("SELECT id, name FROM workspaces WHERE workflow_key = ?", key)

	var id int
	var name string

	err := result.Scan(&id, &name)
	if err != nil {
		return Workspace{}, err
	}

	return Workspace{id, key, name, ""}, nil
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
func (workspace Workspace) Create(userId int) error {
	conn := database.GetConnection()
	defer conn.Close()

	stmt, _ := (*conn).Prepare("INSERT INTO workspaces (workflow_key, name, password) VALUES (?, ?, ?)")
	_, err := stmt.Exec(common.CreateKey(), workspace.Name, workspace.Password)
	return err
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
