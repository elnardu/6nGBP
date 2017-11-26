package model

import (
	"log"

	"gopkg.in/mgo.v2"
)

type MyDB struct {
	*mgo.Database
}

var DB *MyDB

// New return new mydb instance
func NewDB(connectString string) {
	session, err := mgo.Dial(connectString)
	if err != nil {
		log.Fatal(err)
	}
	DB = &MyDB{session.DB("6ngbp")}

	DB.InitActions()
	DB.InitCategories()
	DB.InitUsers()
}
