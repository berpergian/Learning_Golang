package event

const (
    TopicPlayerRegistered = "player.registered"
    TopicPlayerLoggedIn   = "player.loggedin"
)

type PlayerRegistered struct {
    PlayerID string `json:"playerId"`
    Email    string `json:"email"`
    Name     string `json:"name"`
}

type PlayerLoggedIn struct {
    PlayerID string `json:"playerId"`
    Email    string `json:"email"`
}
