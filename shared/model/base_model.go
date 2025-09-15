package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BaseModel struct {
	ID           primitive.ObjectID `bson:"_id"`
	IsActive     bool               `bson:"isActive"`
	CreatedAt    time.Time          `bson:"createdAt"`
	UpdatedAt    time.Time          `bson:"updatedAt"`
	DocumentType string             `bson:"documentType"`
}
