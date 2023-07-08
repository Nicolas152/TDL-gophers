package workspace

import (
	"database/sql"
	"errors"
	"gochat/src/connections/database"
	"gochat/src/helpers/common"
	"gochat/src/models/relationship"
	"strings"
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

// Modelo workspace usado para el cliente
type ClientWorkspace struct {
	WorkspaceKey string
	Name         string
}

// Modelo workspace usado para el servidor
type Workspace struct {
	Id           int
	WorkspaceKey string
	Name         string
	Password     string
	Creator      int
}

func GetWorkspaceByKey(key string) (Workspace, error) {
	conn := database.GetConnection()
	defer conn.Close()

	result := (*conn).QueryRow("SELECT id, name, creator, password FROM workspaces WHERE workflow_key = ?", key)

	var id int
	var name string
	var creator int
	var password string

	if err := result.Scan(&id, &name, &creator, &password); err != nil {
		if err == sql.ErrNoRows {
			return Workspace{}, errors.New("Workspace not found")
		}

		return Workspace{}, err
	}

	return Workspace{id, key, name, password, creator}, nil
}

func Get(userId int) ([]ClientWorkspace, error) {
	workspaces := make([]ClientWorkspace, 0)

	conn := database.GetConnection()
	defer conn.Close()

	results, _ := (*conn).Query("SELECT workflow_key, name FROM workspaces INNER JOIN user_workspace ON workspaces.id = user_workspace.workspace_id WHERE user_workspace.user_id = ?", userId)

	for results.Next() {
		var workspace ClientWorkspace
		results.Scan(&workspace.WorkspaceKey, &workspace.Name)
		workspaces = append(workspaces, workspace)
	}

	return workspaces, nil
}

// Manejo de workspaces
func (workspace Workspace) GetId() int {
	return workspace.Id
}

func (workspace Workspace) Create(userId int) error {
	// Realizo validaciones sobre el workspace
	if err := workspace.Validate(userId); err != nil {
		return err
	}

	conn := database.GetConnection()
	defer conn.Close()

	// Inserto el workspace
	stmt, _ := (*conn).Prepare("INSERT INTO workspaces (workflow_key, name, password, creator) VALUES (?, ?, ?, ?)")
	res, err := stmt.Exec(common.CreateKey(), workspace.Name, workspace.Password, userId)
	if err != nil {
		return err
	}

	// Obtengo el WorkspaceId
	workspaceId, err := res.LastInsertId()
	if err != nil {
		return err
	}

	// Creo la relacion entre el usuario y el workspace
	userWorkflow := relationship.UserWorkspaceRelationship{UserId: userId, WorkspaceId: int(workspaceId)}
	return userWorkflow.Create()
}

func (workspace Workspace) Update(userId int, newWorkspace Workspace) error {
	fields, values := []string{}, []interface{}{}
	// Obtengo los parametros a actualizar
	name, password := newWorkspace.Name, newWorkspace.Password

	if name != "" {
		fields = append(fields, "name = ?")
		values = append(values, name)
	}

	if password != "" {
		fields = append(fields, "password = ?")
		values = append(values, password)
	}

	if len(fields) == 0 {
		return errors.New("No fields to update")
	}

	// Agrego el workspace key
	values = append(values, workspace.WorkspaceKey)

	conn := database.GetConnection()
	defer conn.Close()

	_, err := (*conn).Exec(
		"UPDATE workspaces SET "+strings.Join(fields, ", ")+" WHERE workflow_key = ?",
		values...,
	)
	if err != nil {
		return err
	}

	return nil
}

func (workspace Workspace) Delete(userId int) error {
	conn := database.GetConnection()
	defer conn.Close()

	_, err := (*conn).Exec("DELETE FROM workspaces WHERE workflow_key = ?", workspace.WorkspaceKey)
	if err != nil {
		return err
	}

	return nil
}

func (workspace Workspace) Validate(userId int) error {
	if workspace.Name == "" {
		return errors.New("Name is required")
	}

	// Valido que el usuario no tenga un workspace con el mismo nombre
	conn := database.GetConnection()
	defer conn.Close()

	result := (*conn).QueryRow("SELECT id FROM workspaces WHERE name = ? AND creator = ?", workspace.Name, userId)

	var id int
	err := result.Scan(&id)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return err
	}

	if id > 0 {
		return errors.New("Workspace name already exists")
	}

	return nil
}

func (workspace Workspace) IsPublic() bool {
	return workspace.Password == ""
}

func (workspace Workspace) IsOwner(userId int) bool {
	return workspace.Creator == userId
}

func (workspace Workspace) Authenticate(newWorkspace Workspace) bool {
	return workspace.Password == newWorkspace.Password
}

func (workspace Workspace) Exists() (bool, error) {
	ws, err := GetWorkspaceByKey(workspace.WorkspaceKey)
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
