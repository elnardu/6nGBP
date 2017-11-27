package model

import (
	"errors"
	"log"
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/elnardu/6nGBP/backend/types"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func (db *MyDB) GetUserByID(id *bson.ObjectId) *types.User {
	result := types.User{}
	db.C("users").Find(bson.M{"_id": id}).One(&result)
	return &result
}

func (db *MyDB) GetUsers() *[]*types.User {
	result := &[]*types.User{}
	db.C("users").Find(bson.M{}).Sort("-points").Limit(50).All(result)
	return result
}

func (db *MyDB) CreateUser(user *types.User) error {
	user.ID = bson.NewObjectId()
	err := db.C("users").Insert(user)
	if err != nil && err.Error()[:6] == "E11000" {
		return errors.New("You cannot use this login")
	}
	return err
}

type recalcAggregationResults struct {
	Points int `bson:"points"`
}

func (db *MyDB) RecalcPointsForUser(id *bson.ObjectId) error {
	pipe := db.C("actions").Pipe([]bson.M{{"$match": bson.M{"userId": id}},
		{"$group": bson.M{"_id": "$userId", "points": bson.M{"$sum": "$points"}}}})

	results := recalcAggregationResults{}
	err := pipe.One(&results)
	if err != nil {
		return err
	}
	_, err = db.C("users").Upsert(bson.M{"_id": id}, bson.M{"$set": bson.M{"points": results.Points}})
	return err
}

func (db *MyDB) AuthUser(login string, password string) (string, error) {
	if login == "" || password == "" {
		return "", errors.New("Invalid input")
	}
	result := types.User{}
	db.C("users").Find(bson.M{"login": login}).One(&result)
	var tokenString string
	if result.Password != password {
		return "", errors.New("NOOOOOOPE, something wrong with your login or password")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":    result.ID.Hex(),
		"admin": result.Admin,
		"exp":   time.Now().Add(time.Hour * 48).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		log.Fatal(err)
	}
	return tokenString, nil
}

func (db *MyDB) InitUsers() {
	var err error
	col := db.C("users")

	nameIndex := mgo.Index{
		Key: []string{"fullname"},
	}
	err = col.EnsureIndex(nameIndex)
	if err != nil {
		log.Fatal(err)
	}

	pointsIndex := mgo.Index{
		Key: []string{"points"},
	}
	err = col.EnsureIndex(pointsIndex)
	if err != nil {
		log.Fatal(err)
	}

	loginIndex := mgo.Index{
		Key:    []string{"login"},
		Unique: true,
	}
	err = col.EnsureIndex(loginIndex)
	if err != nil {
		log.Fatal(err)
	}
}
