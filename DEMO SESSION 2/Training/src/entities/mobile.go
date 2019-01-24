package entities

import (
	"gopkg.in/mgo.v2/bson"
)

type Mobile struct {
	Id       bson.ObjectId `json:"id" bson:"_id"`
	Name     string        `json:"name" bson:"name"`
	Price    float64       `json:"price" bson:"price"`
	Quantity int           `json:"quantity" bson:"quantity"`
	Status   bool          `json:"status" bson:"status"`
}
