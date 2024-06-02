package queries

import (
	"database/sql"
	"server/queries/resolvers"
	"server/shared/objects"

	"github.com/graphql-go/graphql"
)

type Querier struct {
	types     map[string]*graphql.Object
	rootQuery *graphql.Object
}

func NewQuerier(db *sql.DB, sc map[string]*graphql.Object) *Querier {
	rootQuery := BuildQuery(db, sc)
	return &Querier{
		types:     sc,
		rootQuery: rootQuery,
	}
}

func (q *Querier) GetRootQuery() *graphql.Object {
	return q.rootQuery
}

func BuildQuery(db *sql.DB, objectTypes map[string]*graphql.Object) *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "RootQuery",
		Fields: graphql.Fields{
			"author": &graphql.Field{
				Type:        objectTypes[objects.AuthorType],
				Description: "Get an author.",
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: resolvers.Author(db),
			},
			"authors": &graphql.Field{
				Type:        graphql.NewList(objectTypes[objects.AuthorType]),
				Description: "List of authors.",
				Resolve:     resolvers.Authors(db),
			},
			"post": &graphql.Field{
				Type:        objectTypes[objects.PostType],
				Description: "Get a post.",
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: resolvers.Post(db),
			},
			"posts": &graphql.Field{
				Type:        graphql.NewList(objectTypes[objects.PostType]),
				Description: "List of posts.",
				Resolve:     resolvers.Posts(db),
			},
		},
	})
}
