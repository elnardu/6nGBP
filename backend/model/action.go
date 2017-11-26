package model

import (
	"errors"
	"log"

	"github.com/elnardu/6nGBP/backend/types"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func (db *MyDB) GetActionByID(id *bson.ObjectId) *types.Action {
	result := types.Action{}
	db.C("actions").Find(bson.M{"_id": id}).One(&result)
	return &result
}
func (db *MyDB) GetActions() *[]*types.Action {
	result := &[]*types.Action{}
	db.C("actions").Find(bson.M{}).Sort("date").Limit(20).All(result)
	return result
}
func (db *MyDB) GetActionsByUserID(id *bson.ObjectId) *[]*types.Action {
	result := &[]*types.Action{}
	db.C("actions").Find(bson.M{"userId": id}).Sort("date").Limit(20).All(result)
	return result
}
func (db *MyDB) CreateAction(action *types.Action) error {
	action.ID = bson.NewObjectId()
	if action.Points == 0 && action.CategoryID.Hex() == "000000000000000000000000" {
		return errors.New("You must specify 'points' or 'category'")
	}
	if action.CategoryID.Hex() != "000000000000000000000000" {
		cat := DB.GetCategoryByID(&action.CategoryID)
		if cat == nil {
			return errors.New("This category does not exist")
		}
		action.Points = cat.Points
	}
	err := db.C("actions").Insert(action)
	DB.RecalcPointsForUser(&action.UserID)
	return err
}

func (db *MyDB) InitActions() {
	var err error
	col := db.C("actions")

	dateIndex := mgo.Index{
		Key: []string{"date"},
	}
	err = col.EnsureIndex(dateIndex)
	if err != nil {
		log.Fatal(err)
	}
}
