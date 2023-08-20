package user

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID             primitive.ObjectID `bson:"_id" `
	Username       string             `bson:"username" `
	Password       string             `bson:"password" `
	Email          string             `bson:"email" `
	EmployeeID     string             `bson:"employee_id" `
	Firstname      string             `bson:"firstname" `
	Lastname       string             `bson:"lastname" `
	JobRole        string             `bson:"job_role" `
	Bio            string             `bson:"introduction" `
	Organization   string             `bson:"organization" `
	SoftSkill      []MySkill          `bson:"soft_skill" `
	TechnicalSkill []MySkill          `bson:"technical_skill" `
}

type MySkill struct {
	ID   primitive.ObjectID `bson:"_id" json:"id"`
	Rank int                `bson:"rank" json:"rank"`
}
