package request

import (
	"encoding/json"
	"errors"
	"github.com/gorilla/websocket"
)

type ActionType string
type ModelType 	string
type Identifier string

const (
	SignInAction ActionType = "signin"
	LogInAction  ActionType = "login"
	ListAction   ActionType = "list"
	CreateAction ActionType = "create"
	UpdateAction ActionType = "update"
	DeleteAction ActionType = "delete"
)

const (
	WorkspaceModel ModelType = "workspace"
)

type RequestInterface interface {
	GetAction() ActionType
	GetModel() ModelType
	GetId() Identifier
	GetParameters() []byte

	HasId() bool
	HasParameters() bool
}

type UserRequest struct {
	UserId 		*int                    `json:"user_id"`
	Action 		*ActionType             `json:"action"`
	Model 		*ModelType               `json:"model"`
	Id 			*Identifier             `json:"id"`
	Parameters 	*map[string]interface{} `json:"parameters"`
}

func (request *UserRequest) ReadRequest(ws *websocket.Conn) error {
	if err := ws.ReadJSON(&request); err != nil {
		return errors.New("Invalid request")
	}

	// Valido campos mandatorios
	if request.Action == nil {
		return errors.New("Action is required")
	}

	return nil
}

func (request *UserRequest) GetAction() ActionType {
	return *request.Action
}

func (request *UserRequest) GetModel() ModelType {
	return *request.Model
}

func (request *UserRequest) GetId() string {
	return string(*request.Id)
}

func (request *UserRequest) HasId() bool {
	return request.Id != nil
}

func (request *UserRequest) GetParameters() []byte {
	data, _ := json.Marshal(request.Parameters)
	return data
}

func (request *UserRequest) HasParameters() bool {
	return request.Parameters != nil
}

func (request *UserRequest) GetUserId() int {
	return *request.UserId
}

func (request *UserRequest) LoadUserId(userId int) {
	request.UserId = &userId
}