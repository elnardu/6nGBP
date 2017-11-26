package model

import (
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
	db.C("categories").Find(bson.M{}).Sort("title").Limit(20).All(result)
	return result
}
func (db *MyDB) CreateCategory(category *types.Category) error {
	category.ID = bson.NewObjectId()
	err := db.C("categories").Insert(category)
	return err
}

func (db *MyDB) InitCategories() {
	titleIndex := mgo.Index{
		Key: []string{"title"},
	}
	err := db.C("categories").EnsureIndex(titleIndex)
	if err != nil {
		log.Fatal(err)
	}
}
