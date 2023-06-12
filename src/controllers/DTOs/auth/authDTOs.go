package authDTO

// LoginDTO represents the structure of a login request
type LoginDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// SignupDTO represents the structure of a signup request
type SignupDTO struct {
	Email    string `json:"email"`
	Name     string `json:"name"`
	Password string `json:"password"`
}
