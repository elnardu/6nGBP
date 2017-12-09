package model

import (
	"errors"
	"log"

	"github.com/elnardu/6nGBP/backend/types"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func (db *MyDB) GetCategoryByID(id *bson.ObjectId) *types.Category {
	result := types.Category{}
	db.C("categories").Find(bson.M{"_id": id}).One(&result)
	return &result
}
func (db *MyDB) GetCategories() *[]*types.Category {
	result := &[]*types.Category{}
	db.C("categories").Find(bson.M{}).Sort("title").All(result)
	return result
}
func (db *MyDB) CreateCategory(category *types.Category) error {
	if category.Title == "" {
		return errors.New("Title cannot be empty")
	}

	if category.Description == "" {
		return errors.New("Description cannot be empty")
	}

	category.ID = bson.NewObjectId()
	err := db.C("categories").Insert(category)
	if err != nil && err.Error()[:6] == "E11000" {
		return errors.New("You cannot use this title")
	}
	return err
}

func (db *MyDB) InitCategories() {
	titleIndex := mgo.Index{
		Key:    []string{"title"},
		Unique: true,
	}
	err := db.C("categories").EnsureIndex(titleIndex)
	if err != nil {
		log.Fatal(err)
	}
}
