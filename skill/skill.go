package skill

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Skill struct {
	ID          primitive.ObjectID `bson:"_id"`
	Name        string             `bson:"name"`
	Description string             `bson:"description"`
	Logo        string             `bson:"logo" `
	Kind        string             `bson:"kind"`
	CreatedAt   time.Time          `bson:"createdAt"`
	UpdatedAt   time.Time          `bson:"updatedAt"`
}

type AddSkillRequest struct {
	Name        string `bson:"name" json:"name"`
	Description string `bson:"description" json:"description"`
	Logo        string `bson:"logo" json:"logo"`
	Kind        string `bson:"kind" json:"kind"`
}
