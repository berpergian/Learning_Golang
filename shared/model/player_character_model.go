package model

type PlayerCharacter struct {
	PlayerBase `bson:",inline"`
	Items      []CharacterStaticData `bson:"items"`
}

func (p *PlayerCharacter) SetDocTypeFrom(v interface{}, docType string) {
	if model, ok := v.(*PlayerCharacter); ok {
		model.PlayerBase.BaseModel.DocumentType = docType
	}
}
