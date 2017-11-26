package model

import (
	"github.com/graphql-go/graphql"
	"gopkg.in/mgo.v2/bson"
)

// QueryType graphql query
var QueryType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Query",
	Fields: graphql.Fields{
		"user": &graphql.Field{
			Type: UserType,
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Description: "User ID",
					Type:        graphql.NewNonNull(graphql.ID),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				id := bson.ObjectId(p.Args["id"].(string))
				return DB.GetUserByID(&id), nil
			},
		},
		"me": &graphql.Field{
			Type: UserType,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				token, err := ValidateAndParseToken(p.Context)
				if err != nil {
					return nil, err
				}
				return DB.GetUserByID(&token.ID), nil
			},
		},
		"users": &graphql.Field{
			Type: graphql.NewNonNull(graphql.NewList(graphql.NewNonNull(UserType))),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return *DB.GetUsers(), nil
			},
		},
		"auth": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.String),
			Description: "Returns [JWT](https://jwt.io/) token",
			Args: graphql.FieldConfigArgument{
				"login": &graphql.ArgumentConfig{
					Description: "Login",
					Type:        graphql.NewNonNull(graphql.String),
				},
				"password": &graphql.ArgumentConfig{
					Description: "Password",
					Type:        graphql.NewNonNull(graphql.String),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				login := p.Args["login"].(string)
				password := p.Args["password"].(string)
				return DB.AuthUser(login, password)
			},
		},
		"actions": &graphql.Field{
			Type: graphql.NewNonNull(graphql.NewList(graphql.NewNonNull(ActionType))),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return *DB.GetActions(), nil
			},
		},
		"categories": &graphql.Field{
			Type: graphql.NewNonNull(graphql.NewList(graphql.NewNonNull(CategoryType))),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return *DB.GetCategories(), nil
			},
		},
	},
})
