package entities

import (
	"gopkg.in/mgo.v2/bson"
)

type Mobile struct {
	Id       bson.ObjectId `bson:"_id"`
	Name     string        `bson:"name"`
	Price    float64       `bson:"price"`
	Quantity int           `bson:"quantity"`
	Status   bool          `bson:"status"`
}
