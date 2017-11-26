package types

import (
	"gopkg.in/mgo.v2/bson"
)

// Category category structure
type Category struct {
	ID          bson.ObjectId `bson:"_id"`
	Title       string        `bson:"title"`
	Description string        `bson:"description"`
	Points      int           `bson:"points"`
	AddedBy     bson.ObjectId `bson:"addedBy"`
}
