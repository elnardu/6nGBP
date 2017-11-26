package model

import (
	"github.com/elnardu/6nGBP/backend/types"
	"github.com/graphql-go/graphql"
)

// ActionType graphql action type
var ActionType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Action",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.NewNonNull(graphql.ID),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if action, ok := p.Source.(*types.Action); ok == true {
					return action.ID.Hex(), nil
				}
				return nil, nil
			},
		},
		"category": &graphql.Field{
			Type: CategoryType,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if action, ok := p.Source.(*types.Action); ok == true {
					return DB.GetCategoryByID(&action.CategoryID), nil
				}
				return nil, nil
			},
		},
		"points": &graphql.Field{
			Type: graphql.NewNonNull(graphql.Int),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if action, ok := p.Source.(*types.Action); ok == true {
					return action.Points, nil
				}
				return nil, nil
			},
		},
		"user": &graphql.Field{
			Type: graphql.NewNonNull(UserType),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if action, ok := p.Source.(*types.Action); ok == true {
					return DB.GetUserByID(&action.UserID), nil
				}
				return nil, nil
			},
		},
		"comment": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if action, ok := p.Source.(*types.Action); ok == true {
					return action.Comment, nil
				}
				return nil, nil
			},
		},
		"date": &graphql.Field{
			Type: graphql.NewNonNull(graphql.Int),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if action, ok := p.Source.(*types.Action); ok == true {
					return action.Date.Unix(), nil
				}
				return nil, nil
			},
		},
		"addedBy": &graphql.Field{
			Type: graphql.NewNonNull(UserType),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if action, ok := p.Source.(*types.Action); ok == true {
					return DB.GetUserByID(&action.AddedBy), nil
				}
				return nil, nil
			},
		},
	},
})
