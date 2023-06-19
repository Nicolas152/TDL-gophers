package channelDTO

// ChannelDTO represents the structure of a create channel request
type ChannelDTO struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}
