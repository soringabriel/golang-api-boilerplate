package user_model

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const CollectionName = "users"

type User struct {
	Id    primitive.ObjectID `bson:"_id" json:"-"`
	Email string             `bson:"email" json:"email"`
	Name  *string            `bson:"name" json:"name"`
}

func (user *User) ToBson() bson.D {
	newDocument := bson.D{
		{Key: "email", Value: user.Email},
	}
	if user.Name != nil {
		newDocument = append(newDocument, bson.E{Key: "name", Value: user.Name})
	}

	return newDocument
}
