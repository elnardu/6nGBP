package model

import (
	"github.com/elnardu/6nGBP/backend/types"
	"github.com/graphql-go/graphql"
)

// UserType graphql user type
var UserType = graphql.NewObject(graphql.ObjectConfig{
	Name: "User",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.NewNonNull(graphql.ID),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if user, ok := p.Source.(*types.User); ok == true {

					return user.ID.Hex(), nil
				}
				return nil, nil
			},
		},
		"login": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if user, ok := p.Source.(*types.User); ok == true {
					return user.Login, nil
				}
				return nil, nil
			},
		},
		"fullname": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if user, ok := p.Source.(*types.User); ok == true {
					return user.Fullname, nil
				}
				return nil, nil
			},
		},
		"admin": &graphql.Field{
			Type: graphql.NewNonNull(graphql.Boolean),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if user, ok := p.Source.(*types.User); ok == true {
					return user.Admin, nil
				}
				return nil, nil
			},
		},
		"points": &graphql.Field{
			Type: graphql.NewNonNull(graphql.Int),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if user, ok := p.Source.(*types.User); ok == true {
					return user.Points, nil
				}
				return nil, nil
			},
		},
	},
})

func init() {
	// fix to circular references compiler error
	UserType.AddFieldConfig("actions", &graphql.Field{
		Type: graphql.NewNonNull(graphql.NewList(graphql.NewNonNull(ActionType))),
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			if user, ok := p.Source.(*types.User); ok == true {
				return *DB.GetActionsByUserID(&user.ID), nil
			}
			return nil, nil
		},
	})
}
