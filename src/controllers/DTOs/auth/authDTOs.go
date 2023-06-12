package authDTO

// LoginDTO represents the structure of a login request
type LoginDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// SignInDTO represents the structure of a signup request
type SignInDTO struct {
	Email    string `json:"email"`
	Name     string `json:"name"`
	Password string `json:"password"`
}
