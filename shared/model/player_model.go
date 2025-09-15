package model

type Player struct {
	PlayerBase `bson:",inline"`
	Name       string `bson:"name"`
	Email      string `bson:"email"`
	Password   string `bson:"password"`
}

func (p *Player) SetDocTypeFrom(v interface{}, docType string) {
	if model, ok := v.(*Player); ok {
		model.PlayerBase.BaseModel.DocumentType = docType
	}
}
