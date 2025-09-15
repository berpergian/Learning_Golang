package model

type BaseStaticDataModel struct {
	ID string `json:"id"`
}

type StaticModel interface {
	GetId() string
}

func (b BaseStaticDataModel) GetId() string { return b.ID }
