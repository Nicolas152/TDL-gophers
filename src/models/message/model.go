package message

type Message struct {
	WorkspaceKey 	string	`json:"workspaceKey"`
	ChannelKey   	string	`json:"channelKey"`
	Message  		string	`json:"message"`
}
