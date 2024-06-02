package objects

import (
	"server/entities"

	"github.com/graphql-go/graphql"
)

func InitAuthorType() *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name:        "Author",
		Description: "An author",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.Int),
				Description: "The identifier of the author.",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if author, ok := p.Source.(*entities.Author); ok {
						return author.ID, nil
					}

					return nil, nil
				},
			},
			"name": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.String),
				Description: "The name of the author.",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if author, ok := p.Source.(*entities.Author); ok {
						return author.Name, nil
					}

					return nil, nil
				},
			},
			"email": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.String),
				Description: "The email address of the author.",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if author, ok := p.Source.(*entities.Author); ok {
						return author.Email, nil
					}

					return nil, nil
				},
			},
			"created_at": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.String),
				Description: "The created_at date of the author.",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if author, ok := p.Source.(*entities.Author); ok {
						return author.CreatedAt, nil
					}

					return nil, nil
				},
			},
		},
	})

}
