package model

type ItemRarity string

const (
	Common ItemRarity = "Common"
	Rare   ItemRarity = "Rare"
	Epic   ItemRarity = "Epic"
)

type ItemStaticData struct {
	BaseStaticDataModel `bson:",inline"`
	Name                string     `json:"name"`
	Rarity              ItemRarity `json:"rarity"`
}
