package channel

import (
	"encoding/json"
	"errors"
	"gochat/src/models/channel"
	"gochat/src/models/chat"
	"gochat/src/models/workspace"
	"net/http"
)

func GetChannelsByWorkspace(workspaceKey string, userId int) ([]byte, error, int) {

	// validate if workspaceModel exists
	workspaceModel, err := workspace.GetWorkspaceByKey(workspaceKey)
	if err != nil {
		return nil, errors.New("Error validating workspace: " + err.Error()), http.StatusInternalServerError
	}

	if err, statusErr := ChannelValidations(workspaceModel, userId); err != nil {
		return nil, err, statusErr
	}

	channels, err := channel.GetChannelsByWorkspaceId(workspaceModel.Id)
	if err != nil {
		return nil, errors.New("Error getting channels: " + err.Error()), http.StatusInternalServerError
	}

	channelsJson, err := json.Marshal(channels)
	if err != nil {
		return nil, errors.New("Error marshalling channels: " + err.Error()), http.StatusInternalServerError
	}

	return channelsJson, err, 0
}

func GetChannel(id int, workspaceKey string, userId int) ([]byte, error, int) {

	// validate if workspaceModel exists
	workspaceModel, err := workspace.GetWorkspaceByKey(workspaceKey)
	if err != nil {
		return nil, errors.New("Error validating workspace: " + err.Error()), http.StatusInternalServerError
	}

	if err, statusErr := ChannelValidations(workspaceModel, userId); err != nil {
		return nil, err, statusErr
	}

	channelModel := channel.Channel{WorkspaceId: workspaceModel.Id, Id: id}
	channelModel, err = channelModel.Get()
	if err != nil {
		return nil, errors.New("Could not get channel. Reason:" + err.Error()), http.StatusBadRequest
	}

	channelJson, err := json.Marshal(channelModel)
	if err != nil {
		return nil, errors.New("Error marshalling channel: " + err.Error()), http.StatusInternalServerError
	}

	return channelJson, err, 0
}

// ChannelValidations performs validations to determine if the user has access to the workspace.
// It returns an error and a corresponding HTTP status code based on the validation results.
func ChannelValidations(workspaceModel workspace.Workspace, userId int) (error, int) {

	if workspaceModel.Id == 0 {
		return errors.New("Workspace does not exists"), http.StatusBadRequest
	}

	// validate if user is member of workspace
	exists, err2 := workspaceModel.HasMember(userId)
	if err2 != nil {
		return errors.New("Error validating if user is member of workspace: " + err2.Error()), http.StatusInternalServerError
	}
	if !exists {
		return errors.New("User is not member of workspace"), http.StatusUnauthorized
	}

	return nil, 0
}

func CreateChannel(name string, password string, workspaceKey string, userId int) (error, int) {

	// validate if workspaceModel exists
	workspaceModel, err := workspace.GetWorkspaceByKey(workspaceKey)
	if err != nil {
		return errors.New("Error validating workspace: " + err.Error()), http.StatusInternalServerError
	}

	if err, statusErr := ChannelValidations(workspaceModel, userId); err != nil {
		return err, statusErr
	}
	// create channel
	channelModel := channel.Channel{Name: name, Password: password, WorkspaceId: workspaceModel.Id, Creator: userId}
	if err := channelModel.Create(); err != nil {
		return errors.New("Error creating channel: " + err.Error()), http.StatusInternalServerError
	}

	return nil, 0
}

func UpdateChannel(id int, name string, password string, workspaceKey string, userId int) (error, int) {

	// validate if workspaceModel exists
	workspaceModel, err := workspace.GetWorkspaceByKey(workspaceKey)
	if err != nil {
		return errors.New("Error validating workspace: " + err.Error()), http.StatusInternalServerError
	}

	if err, statusErr := ChannelValidations(workspaceModel, userId); err != nil {
		return err, statusErr
	}

	channelModel := channel.Channel{WorkspaceId: workspaceModel.Id, Id: id}
	channelModel, err = channelModel.Get()
	if err != nil {
		return errors.New("Could not update Channel. Reason:" + err.Error()), http.StatusBadRequest
	}

	// check if user is the creator of the channel
	if !channelModel.IsOwner(userId) {
		return errors.New("User is not owner of channel"), http.StatusUnauthorized
	}

	// update channel
	channelModel.Name = name
	channelModel.Password = password
	if err := channelModel.Update(); err != nil {
		return errors.New("Error updating channel: " + err.Error()), http.StatusInternalServerError
	}

	return nil, 0
}

func DeleteChannel(id int, workspaceKey string, userId int) (error, int) {

	// validate if workspaceModel exists
	workspaceModel, err := workspace.GetWorkspaceByKey(workspaceKey)
	if err != nil {
		return errors.New("Error validating workspace: " + err.Error()), http.StatusInternalServerError
	}

	if err, statusErr := ChannelValidations(workspaceModel, userId); err != nil {
		return err, statusErr
	}

	channelModel := channel.Channel{WorkspaceId: workspaceModel.Id, Id: id}
	channelModel, err = channelModel.Get()
	if err != nil {
		return errors.New("Could not delete Channel. Reason:" + err.Error()), http.StatusBadRequest
	}

	// check if user is the creator of the channel
	if !channelModel.IsOwner(userId) {
		return errors.New("User is not owner of channel"), http.StatusUnauthorized
	}

	// delete channel
	if err := channelModel.Delete(); err != nil {
		return errors.New("Error deleting channel: " + err.Error()), http.StatusInternalServerError
	}

	return nil, 0
}

func JoinToChannel(id int, password string, workspaceKey string, userId int) (error, int) {

	// validate if workspaceModel exists
	workspaceModel, err := workspace.GetWorkspaceByKey(workspaceKey)
	if err != nil {
		return errors.New("Error validating workspace: " + err.Error()), http.StatusInternalServerError
	}

	if err, statusErr := ChannelValidations(workspaceModel, userId); err != nil {
		return err, statusErr
	}

	channelModel := channel.Channel{WorkspaceId: workspaceModel.Id, Id: id}
	channelModel, err = channelModel.Get()
	if err != nil {
		return errors.New("Could not join Channel. Reason:" + err.Error()), http.StatusBadRequest
	}

	// check if channel is password protected
	if channelModel.Password != password {
		return errors.New("Invalid password"), http.StatusUnauthorized
	}

	// join channel
	if err := channelModel.Join(userId); err != nil {
		return errors.New("Error joining channel: " + err.Error()), http.StatusInternalServerError
	}

	return nil, 0
}

func MembersOfChannel(id int, workspaceKey string, userId int) ([]byte, error, int) {

	// validate if workspaceModel exists
	workspaceModel, err := workspace.GetWorkspaceByKey(workspaceKey)
	if err != nil {
		return nil, errors.New("Error validating workspace: " + err.Error()), http.StatusInternalServerError
	}

	if err, statusErr := ChannelValidations(workspaceModel, userId); err != nil {
		return nil, err, statusErr
	}

	channelModel := channel.Channel{WorkspaceId: workspaceModel.Id, Id: id}
	channelModel, err = channelModel.Get()
	if err != nil {
		return nil, errors.New("Could not get members of Channel. Reason:" + err.Error()), http.StatusBadRequest
	}

	// only members of private channels can see its members
	// public channels are visible to everyone in the workspace
	if !channelModel.IsMember(userId) && channelModel.Password != "" {
		return nil, errors.New("User is not member of this private channel"), http.StatusUnauthorized
	}

	members, err := channelModel.GetMembers()
	if err != nil {
		return nil, errors.New("Error getting members of channel: " + err.Error()), http.StatusInternalServerError
	}

	membersJson, err := json.Marshal(members)
	if err != nil {
		return nil, errors.New("Error marshalling members: " + err.Error()), http.StatusInternalServerError
	}

	return membersJson, nil, 0
}

func LeaveChannel(id int, workspaceKey string, userId int) (error, int) {

	// validate if workspaceModel exists
	workspaceModel, err := workspace.GetWorkspaceByKey(workspaceKey)
	if err != nil {
		return errors.New("Error validating workspace: " + err.Error()), http.StatusInternalServerError
	}

	if err, statusErr := ChannelValidations(workspaceModel, userId); err != nil {
		return err, statusErr
	}

	channelModel := channel.Channel{WorkspaceId: workspaceModel.Id, Id: id}
	channelModel, err = channelModel.Get()
	if err != nil {
		return errors.New("Could not leave Channel. Reason:" + err.Error()), http.StatusBadRequest
	}

	// leave channel
	if err := channelModel.Leave(userId); err != nil {
		return errors.New("Error leaving channel: " + err.Error()), http.StatusInternalServerError
	}

	return nil, 0
}

func Messages(id int, workspaceKey string, userId int) ([]byte, error, int) {

	// validate if workspaceModel exists
	workspaceModel, err := workspace.GetWorkspaceByKey(workspaceKey)
	if err != nil {
		return nil, errors.New("Error validating workspace: " + err.Error()), http.StatusInternalServerError
	}

	if err, statusErr := ChannelValidations(workspaceModel, userId); err != nil {
		return nil, err, statusErr
	}

	chatId, err := chat.Resolve(id, 0)
	messages, err := chat.GetMessages(chatId)
	if err != nil {
		return nil, errors.New("Could not get messages of Channel. Reason:" + err.Error()), http.StatusBadRequest
	}

	messagesJson, err := json.Marshal(messages)
	if err != nil {
		return nil, errors.New("Error marshalling messages: " + err.Error()), http.StatusInternalServerError
	}

	return messagesJson, nil, 0
}
