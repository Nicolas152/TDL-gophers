package user

import (
	"fiuba/concurrent/gochat/src/connections/database"
	"fiuba/concurrent/gochat/src/models/workspace"
)

type User struct {
	Id int
	Name string
	Email string
}

func Exists(email string) (bool, error) {
	// Valido que el User exista en DB
	conn := database.GetConnection()
	defer conn.Close()

	result := (*conn).QueryRow("SELECT COUNT(*) FROM users WHERE email = ?", email)

	var count int

	err := result.Scan(&count)
	if err != nil {
		return false, err
	}

	if count == 0 {
		return false, nil
	}

	return true, nil
}

func GetUser(email string) (User, error) {
	// Obtengo el usuario de DB
	conn := database.GetConnection()
	defer conn.Close()

	result := (*conn).QueryRow("SELECT id, name FROM users WHERE email = ?", email)

	var id int
	var name string

	err := result.Scan(&id, &name)
	if err != nil {
		return User{}, err
	}

	return User{id, name, email}, nil
}

func (user User) ValidatePassword(password string) bool {
	// Valido que el password sea correcto
	conn := database.GetConnection()
	defer conn.Close()

	result := (*conn).QueryRow("SELECT COUNT(*) FROM users WHERE id = ? AND password = ?", user.Id, password)

	var count int

	err := result.Scan(&count)
	if err != nil || count == 0 {
		return false
	}

	return true
}

// Manejo de usuarios

func Get() []User {
	// Obtengo todos los usuarios de DB
	users := make([]User, 0)

	conn := database.GetConnection()
	defer conn.Close()

	results, _ := (*conn).Query("SELECT id, name, email FROM users")

	for results.Next() {
		var user User
		results.Scan(&user.Id, &user.Name, &user.Email)
		users = append(users, user)
	}

	return users

}

func Create(email string, name string, password string) (User, error) {
	// Creo el usuario en DB
	conn := database.GetConnection()
	defer conn.Close()

	stmt, _ := (*conn).Prepare("INSERT INTO users(email, name, password) VALUES(?, ?, ?)")
	res, err := stmt.Exec(email, name, password)
	if err != nil {
		return User{}, err
	}

	id, _ := res.LastInsertId()
	return User{int(id), name, email}, nil
}

func (user User) Modify(name string, password string) error {
	// Modifico el usuario en DB
	conn := database.GetConnection()
	defer conn.Close()

	_, err := (*conn).Exec("UPDATE users SET name = ?, password = ? WHERE id = ?", name, password, user.Id)
	if err != nil {
		return err
	}

	return nil
}

func (user User) Delete() error {
	// Elimino el usuario de DB
	conn := database.GetConnection()
	defer conn.Close()

	_, err := (*conn).Exec("DELETE FROM users WHERE id = ?", user.Id)
	if err != nil {
		return err
	}

	return nil
}

// Manejo de workspaces

func (user User) GetWorkspaces() []workspace.Workspace {
	// Obtengo los workspaces del usuario
	workspaces := make([]workspace.Workspace, 0)

	conn := database.GetConnection()
	defer conn.Close()

	results, _ := (*conn).Query("SELECT workspaces.id, workspaces.name FROM workspaces JOIN user_workspace ON workspaces.id = user_workspace.workspace_id WHERE user_workspace.user_id = ?", user.Id)

	for results.Next() {
		var workspace workspace.Workspace
		results.Scan(&workspace.Id, &workspace.Name)
		workspaces = append(workspaces, workspace)
	}

	return workspaces
}

func (user User)  CreateWorkspace(name string, description string, password string) (workspace.Workspace, error) {
	// Creo el workspace
	newWorkspace, err := workspace.Create(name, description, password)
	if err != nil {
		return workspace.Workspace{}, err
	}

	// Creo la relacion entre el usuario y el workspace
	err = newWorkspace.AddUser(user.Id)
	if err != nil {
		return workspace.Workspace{}, err
	}

	return newWorkspace, nil
}

func (user User) ModifyWorkspace(workspace *workspace.Workspace, name string, description string, password string) error {
	// Modifico el workspace
	return workspace.Modify(name, description, password)
}

func (user User) DeleteWorkspace(workspace *workspace.Workspace) error {
	// Elimino el workspace
	return workspace.Delete()
}