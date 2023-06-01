package authentication

import (
	"errors"
	"gochat/src/models/user"
)

type AuthenticationInterface interface {
	SignIn() error
	LogIn() error
	Access() error
}

type UserCredentials struct {
	Email		*string `json:"email"`
	Name 		*string `json:"name"`
	Password 	*string `json:"password"`
}

type AuthenticationFunction func() error


func (credentials *UserCredentials) SignIn() error {
	// Son necesarios el email, el nombre y el password
	if credentials.Email == nil || credentials.Name == nil || credentials.Password == nil {
		return errors.New("Email, name and password are required to sign in to GoChat")
	}

	// Creo el usuario
	userModel := user.User{Email: *credentials.Email, Name: *credentials.Name, Password: *credentials.Password}
	if err := userModel.Create(); err != nil {
		return err
	}

	return nil
}

func (credentials *UserCredentials) LogIn() error {
	// Son necesarios el email y el password
	if credentials.Email == nil || credentials.Password == nil {
		return errors.New("Email and password are required to log in to GoChat")
	}

	// Valido las credenciales del usuario
	userModel := user.User{Email: *credentials.Email, Password: *credentials.Password}
	if authenticated := userModel.Authenticate(); !authenticated {
		return errors.New("Invalid email or password")
	}

	return nil
}

func (credentials *UserCredentials) Access() error {
	// Valido las credenciales del usuario
	if credentials.Email == nil {
		return errors.New("Email is required to access to GoChat")
	}

	return nil
}