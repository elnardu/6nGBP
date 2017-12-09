package types

import (
	"gopkg.in/mgo.v2/bson"
)

// User user struct
type User struct {
	ID          bson.ObjectId `bson:"_id"`
	Login       string        `bson:"login"`
	Fullname    string        `bson:"fullname"`
	Admin       bool          `bson:"admin"`
	Points      int           `bson:"points"`
	PointsSpent int           `bson:"pointsSpent"`
	Password    string        `bson:"password"`
}
