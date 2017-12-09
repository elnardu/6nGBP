package model

import (
	"errors"
	"time"

	"github.com/elnardu/6nGBP/backend/types"
	"github.com/graphql-go/graphql"
	"gopkg.in/mgo.v2/bson"
)

// MutationType graphql mutation type
var MutationType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Mutation",
	Fields: graphql.Fields{
		"createUser": &graphql.Field{
			Type: graphql.NewNonNull(UserType),
			Args: graphql.FieldConfigArgument{
				"login": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"fullname": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"password": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				user := &types.User{
					Login:       p.Args["login"].(string),
					Fullname:    p.Args["fullname"].(string),
					Password:    p.Args["password"].(string), // yep, I am lazy
					Points:      0,
					PointsSpent: 0,
					Admin:       false,
				}
				err := DB.CreateUser(user)
				return user, err
			},
		},
		"createCategory": &graphql.Field{
			Type:        graphql.NewNonNull(CategoryType),
			Description: "admin only",
			Args: graphql.FieldConfigArgument{
				"title": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"description": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"points": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.Int),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				token, err := ValidateAndParseToken(p.Context)
				if err != nil {
					return nil, err
				}
				if token.Admin != true {
					return nil, errors.New("You are not admin")
				}

				category := &types.Category{
					Title:       p.Args["title"].(string),
					Description: p.Args["description"].(string),
					Points:      p.Args["points"].(int),
					AddedBy:     token.ID,
				}
				err = DB.CreateCategory(category)
				return category, err
			},
		},
		"createAction": &graphql.Field{
			Type:        graphql.NewNonNull(ActionType),
			Description: "admin only",
			Args: graphql.FieldConfigArgument{
				"user": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.ID),
				},
				"category": &graphql.ArgumentConfig{
					Type: graphql.ID,
				},
				"points": &graphql.ArgumentConfig{
					Type: graphql.Int,
				},
				"comment": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				token, err := ValidateAndParseToken(p.Context)
				if err != nil {
					return nil, err
				}

				if token.Admin != true {
					return nil, errors.New("You are not admin")
				}

				var categoryID bson.ObjectId
				categoryIDString, ok := p.Args["category"].(string)
				if ok == true {
					categoryID = bson.ObjectIdHex(categoryIDString)
				} else {
					categoryID = bson.ObjectIdHex("000000000000000000000000")
				}

				points, ok := p.Args["points"].(int)
				if ok != true {
					points = 0
				}

				comment, ok := p.Args["comment"].(string)
				if ok != true {
					comment = ""
				}

				action := &types.Action{
					UserID:     bson.ObjectIdHex(p.Args["user"].(string)),
					CategoryID: categoryID,
					Points:     points,
					Comment:    comment,
					AddedBy:    token.ID,
					Date:       time.Now(),
				}
				err = DB.CreateAction(action)
				return action, err
			},
		},
	},
})
