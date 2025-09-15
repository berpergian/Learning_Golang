package model

type CharacterStaticData struct {
	BaseStaticDataModel `bson:",inline"`
	Name                string `json:"name"`
	Level               int    `json:"level"`
}
