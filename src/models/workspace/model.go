package workspace

import (
	"fiuba/concurrent/gochat/src/connections/database"
)

type Workspace struct {
	Id 			int
	Name 		string
	Description string
}


func EmptyWorkspace() Workspace {
	return Workspace{Id: -1}
}

func (workspace Workspace) IsEmpty() bool {
	return workspace.Id == -1
}

func GetWorkspaceById(id int) (Workspace, error) {
	conn := database.GetConnection()
	defer conn.Close()

	result := (*conn).QueryRow("SELECT name, description FROM workspaces WHERE id = ?", id)

	var name string
	var description string

	err := result.Scan(&name, &description)
	if err != nil {
		return Workspace{}, err
	}

	return Workspace{id, name, description}, nil
}

// Manejo de workspaces

func Get() []Workspace {
	workspaces := make([]Workspace, 0)

	conn := database.GetConnection()
	defer conn.Close()

	results, _ := (*conn).Query("SELECT id, name, description FROM workspaces")

	for results.Next() {
		var workspace Workspace
		results.Scan(&workspace.Id, &workspace.Name, &workspace.Description)
		workspaces = append(workspaces, workspace)
	}

	return workspaces
}

func Create(name string, description string, password string) (Workspace, error) {
	conn := database.GetConnection()
	defer conn.Close()

	stmt, _ := (*conn).Prepare("INSERT INTO workspaces (name, description, password) VALUES (?, ?, ?)")
	res, err := stmt.Exec(name, description, password)
	if err != nil {
		return Workspace{}, err
	}

	id, _ := res.LastInsertId()
	return Workspace{int(id), name, description}, err
}

func (workspace Workspace) Modify(name string, description string, password string) error {
	conn := database.GetConnection()
	defer conn.Close()

	_, err := (*conn).Exec("UPDATE workspaces SET name = ?, description = ? WHERE id = ?", name, description, workspace.Id)
	if err != nil {
		return err
	}

	return nil
}

func (workspace Workspace) Delete() error {
	conn := database.GetConnection()
	defer conn.Close()

	_, err := (*conn).Exec("DELETE FROM workspaces WHERE id = ?", workspace.Id)
	if err != nil {
		return err
	}

	return nil
}

// Manejo de relaciones con usuarios

func (workspace Workspace) GetUsers() []int {
	users := make([]int, 0)

	conn := database.GetConnection()
	defer conn.Close()

	results, _ := (*conn).Query("SELECT user_id FROM user_workspace WHERE workspace_id = ?", workspace.Id)

	for results.Next() {
		var userId int
		results.Scan(&userId)
		users = append(users, userId)
	}

	return users
}

func (workspace Workspace) AddUser(userId int) error {
	conn := database.GetConnection()
	defer conn.Close()

	_, err := (*conn).Exec("INSERT INTO user_workspace (user_id, workspace_id) VALUES (?, ?)", userId, workspace.Id)
	if err != nil {
		return err
	}

	return nil
}

func (workspace Workspace) RemoveUser(userId int) error {
	conn := database.GetConnection()
	defer conn.Close()

	_, err := (*conn).Exec("DELETE FROM user_workspace WHERE user_id = ? AND workspace_id = ?", userId, workspace.Id)
	if err != nil {
		return err
	}

	return nil
}