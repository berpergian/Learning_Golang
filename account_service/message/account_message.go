package message

type RegisterRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
	Name     string `json:"name" validate:"required"`
}

type RegisterResponse struct {
	AccessToken string `json:"accessToken"`
	PlayerId    string `json:"playerId"`
	Email       string `json:"email"`
	Name        string `json:"name"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type LoginResponse struct {
	AccessToken string `json:"accessToken"`
}
