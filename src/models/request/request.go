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
type ModelType string
type Identifier string

type RequestInterface interface {
	GetUserId() Identifier
	GetParameters() []byte

	HasUserId() bool
	HasParameters() bool
}

type UserRequest struct {
	AccessToken *string                 `json:"access_token"`
	UserId      *int                    `json:"user_id"`
	Parameters  *map[string]interface{} `json:"parameters"`
	Context     *context.Context        `json:"context"`
}

// General method to load the client's request
func (request *UserRequest) ReadRequest(r *http.Request) error {
	// Cargo los headers de interes
	accessToken := authentication.GetJWTHeader(r)
	request.AccessToken = &accessToken

	// load parameters
	request.Parameters = nil
	if r.Body != http.NoBody {
		// Consume the body and saves the content for later
		requestBody, _ := ioutil.ReadAll(r.Body)
		body := bytes.NewBuffer(requestBody)
		r.Body.Close()

		// Decode the body to obtain the parameters
		var parameters map[string]interface{}
		if err := json.NewDecoder(body).Decode(&parameters); err != nil {
			if !request.HasTokenAccess() {
				// If there are no parameters and no access token, then it's an error
				return err
			}
		}

		// Load the relevant parameters
		request.Parameters = &parameters

		// Re-load the body so that the server can read it
		r.Body = ioutil.NopCloser(bytes.NewBuffer(requestBody))
	}

	// Load the relevant context
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

	// If the token is valid, then obtain the userId
	userId, err := authentication.GetJWTUserId(*request.AccessToken)
	if err != nil {
		return false
	}

	// Validate that the userId exists
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
