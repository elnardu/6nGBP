package model

import (
	"github.com/elnardu/6nGBP/backend/types"
	"github.com/graphql-go/graphql"
)

// CategoryType graphql category type
var CategoryType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Category",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.NewNonNull(graphql.ID),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if category, ok := p.Source.(*types.Category); ok == true {
					return category.ID.Hex(), nil
				}
				return nil, nil
			},
		},
		"title": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if category, ok := p.Source.(*types.Category); ok == true {
					return category.Title, nil
				}
				return nil, nil
			},
		},
		"description": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if category, ok := p.Source.(*types.Category); ok == true {
					return category.Description, nil
				}
				return nil, nil
			},
		},
		"points": &graphql.Field{
			Type: graphql.NewNonNull(graphql.Int),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if category, ok := p.Source.(*types.Category); ok == true {
					return category.Points, nil
				}
				return nil, nil
			},
		},
	},
})

func init() {
	// fix to circular references compiler error
	CategoryType.AddFieldConfig("addedBy", &graphql.Field{
		Type: graphql.NewNonNull(UserType),
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			if cat, ok := p.Source.(*types.Category); ok == true {
				return DB.GetUserByID(&cat.AddedBy), nil
			}
			return nil, nil
		},
	})
}
