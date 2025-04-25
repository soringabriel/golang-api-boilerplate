package models

import "go.mongodb.org/mongo-driver/bson"

type Model interface {
	ToBson() bson.D
}
