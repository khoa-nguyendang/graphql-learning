package mutations

import (
	"database/sql"
	"server/mutations/resolvers"
	"server/shared/objects"

	"github.com/graphql-go/graphql"
)

type Mutator struct {
	types        map[string]*graphql.Object
	rootMutation *graphql.Object
}

func (m *Mutator) GetRootMutation() *graphql.Object {
	return m.rootMutation
}

func InitMutator(db *sql.DB, ot map[string]*graphql.Object) *Mutator {
	return &Mutator{
		types:        ot,
		rootMutation: BuildMutator(db, ot),
	}
}

func BuildMutator(db *sql.DB, objectTypes map[string]*graphql.Object) *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "RootMutation",
		Fields: graphql.Fields{
			// Author
			"createAuthor": &graphql.Field{
				Type:        objectTypes[objects.AuthorType],
				Description: "Create new author",
				Args: graphql.FieldConfigArgument{
					"name": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"email": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: resolvers.CreateAuthor(db),
			},
			"updateAuthor": &graphql.Field{
				Type:        objectTypes[objects.AuthorType],
				Description: "Update an author",
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
					"name": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"email": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: resolvers.UpdateAuthor(db),
			},
			"deleteAuthor": &graphql.Field{
				Type:        objectTypes[objects.AuthorType],
				Description: "Delete an author",
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: resolvers.DeleteAuthor(db),
			},
			// Post
			"createPost": &graphql.Field{
				Type:        objectTypes[objects.PostType],
				Description: "Create new post",
				Args: graphql.FieldConfigArgument{
					"title": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"content": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"author_id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
				},
				Resolve: resolvers.CreatePost(db),
			},
			"updatePost": &graphql.Field{
				Type:        objectTypes[objects.PostType],
				Description: "Update a post",
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
					"title": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"content": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"author_id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
				},
				Resolve: resolvers.UpdatePost(db),
			},
			"deletePost": &graphql.Field{
				Type:        objectTypes[objects.PostType],
				Description: "Delete a post",
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: resolvers.DeletePost(db),
			},
		},
	})
}
