package message

type LoginRequest struct {
    Email    string `json:"email"`
    Password string `json:"password"`
    Name     string `json:"name,omitempty"`
}

type LoginResponse struct {
    AccessToken string `json:"accessToken"`
    PlayerId    string `json:"playerId"`
    Email       string `json:"email"`
    Name        string `json:"name"`
}
