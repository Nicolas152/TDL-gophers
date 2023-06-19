package request

import (
	"bytes"
	"context"
	"encoding/json"
	"gochat/src/helpers/authentication"
	"gochat/src/models/user"
	"io/ioutil"
	"net/http"
)

type ActionType string
type ModelType 	string
type Identifier string


type RequestInterface interface {
	GetUserId() Identifier
	GetParameters() []byte

	HasUserId() bool
	HasParameters() bool
}

type UserRequest struct {
	AccessToken *string                 `json:"access_token"`
	UserId 		*int                    `json:"user_id"`
	Parameters 	*map[string]interface{} `json:"parameters"`
	Context 	*context.Context 		`json:"context"`
}

// Metodo general para cargar la request del cliente
func (request *UserRequest) ReadRequest(r *http.Request) error {
	// Cargo los headers de interes
	accessToken := authentication.GetJWTHeader(r)
	request.AccessToken = &accessToken

	// Cargo los parameters
	request.Parameters = nil
	if r.Body != http.NoBody {
		// Consumo el body y guardo el contenido para despues
		requestBody, _ := ioutil.ReadAll(r.Body)
		body := bytes.NewBuffer(requestBody)
		r.Body.Close()

		// Decodifico el body para obtener los parametros
		var parameters map[string]interface{}
		if err := json.NewDecoder(body).Decode(&parameters); err != nil {
			if !request.HasTokenAccess() {
				// Si no tiene parametros y no tiene token de acceso, entonces es un error
				return err
			}
		}

		// Cargo los parametros de interes
		request.Parameters = &parameters

		// Vuelvo a cargar el body para que el servidor pueda leerlo
		r.Body = ioutil.NopCloser(bytes.NewBuffer(requestBody))
	}

	// Cargo el contexto de interes
	context := r.Context()
	request.Context = &context

	return nil
}

func (request *UserRequest) HasTokenAccess() bool {
	return request.AccessToken != nil
}

func (request *UserRequest) IsTokenAccessValid() bool {
	if !authentication.ValidateJWT(*request.AccessToken) {
		return false
	}

	// Si el token es valido, entonces obtengo el userId
	userId, err := authentication.GetJWTUserId(*request.AccessToken)
	if err != nil {
		return false
	}

	// Valido que el userId exista
	userModel, err := user.GetUserById(userId)
	if err != nil {
		return false
	}

	request.UserId = &userModel.Id
	return true
}

func (request *UserRequest) GetUserId() int {
	if !request.HasUserId() {
		userId := (*request.Context).Value("UserId").(int)
		request.UserId = &userId
	}

	return *request.UserId
}

func (request *UserRequest) HasUserId() bool {
	return request.UserId != nil
}

func (request *UserRequest) GetParameters() []byte {
	data, _ := json.Marshal(request.Parameters)
	return data
}

func (request *UserRequest) HasParameters() bool {
	return request.Parameters != nil
}
