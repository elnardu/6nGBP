package types

import "time"
import "gopkg.in/mgo.v2/bson"

// Action user action structure
type Action struct {
	ID         bson.ObjectId `bson:"_id"`
	CategoryID bson.ObjectId `bson:"categoryId"`
	Points     int           `bson:"points"`
	Comment    string        `bson:"comment"`
	Date       time.Time     `bson:"date"`
	UserID     bson.ObjectId `bson:"userId"`
	AddedBy    bson.ObjectId `bson:"addedBy"`
}
