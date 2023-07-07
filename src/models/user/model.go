package user

import (
	"errors"
	"fmt"
	"gochat/src/connections/database"
)

type UserInterface interface {
	Authenticate() bool
	GetContext() error
	Create() error
}

type User struct {
	Id       int
	Email    string
	Name     string
	Password string
}

type UserClient struct {
	Id    int
	Email string
	Name  string
}

func GetUserById(id int) (User, error) {
	// Obtengo un usuario por id
	conn := database.GetConnection()
	defer conn.Close()

	result := (*conn).QueryRow("SELECT name, email FROM users WHERE id = ?", id)

	var name string
	var email string

	err := result.Scan(&name, &email)
	if err != nil {
		return User{}, err
	}

	return User{id, email, name, ""}, nil
}

func GetUserIDByEmail(email string) (int, error) {
    conn := database.GetConnection()
    defer conn.Close()

    var userID int
    query := "SELECT id FROM users WHERE email = ?"
    err := conn.QueryRow(query, email).Scan(&userID)
    if err != nil {
        return 0, err
    }

    return userID, nil
}


func Get() []User {
	// Obtengo todos los usuarios de DB
	users := make([]User, 0)

	conn := database.GetConnection()
	defer conn.Close()

	results, _ := (*conn).Query("SELECT id, name, email FROM users")

	for results.Next() {
		var user User

		if err := results.Scan(&user.Id, &user.Name, &user.Email); err == nil {
			users = append(users, user)
		}
	}

	return users

}

// Manejo de usuarios

func (user User) Authenticate() bool {
	var count int

	// Autentico el usuario en DB
	conn := database.GetConnection()
	defer conn.Close()

	// Valido que email y password sea correcto
	result := (*conn).QueryRow("SELECT COUNT(*) FROM users WHERE email = ? AND password = ?", user.Email, user.Password)
	if err := result.Scan(&count); err != nil || count == 0 {
		return false
	}

	return true
}

func (user *User) GetContext() error {
	// Obtengo el contexto del usuario
	conn := database.GetConnection()
	defer conn.Close()

	result := (*conn).QueryRow("SELECT id, name FROM users WHERE email = ?", user.Email)

	var id int
	var name string

	err := result.Scan(&id, &name)
	if err != nil {
		return err
	}

	user.Id = id
	user.Name = name

	return nil
}

func (user User) Create() error {
	// Creo el usuario en DB
	conn := database.GetConnection()
	defer conn.Close()

	stmt, err := (*conn).Prepare("INSERT INTO users(email, name, password) VALUES(?, ?, ?)")
	// handle error
	if err != nil {
		errMsg := fmt.Sprintf("error preparing query: %v", err)
		return errors.New(errMsg)
	}

	_, err = stmt.Exec(user.Email, user.Name, user.Password)
	if err != nil {
		return errors.New("User already exists")
	}

	return nil
}
