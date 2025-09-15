package model

type PlayerInventory struct {
	PlayerBase `bson:",inline"`
	Items      []ItemStaticData `bson:"items"`
}

func (p *PlayerInventory) SetDocTypeFrom(v interface{}, docType string) {
	if model, ok := v.(*PlayerInventory); ok {
		model.PlayerBase.BaseModel.DocumentType = docType
	}
}
