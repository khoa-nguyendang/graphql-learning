package resolvers

import (
	"database/sql"
	"server/entities"
	"server/shared"

	"github.com/graphql-go/graphql"
)

// PostAuthor a field resolver for Post's Author field
func PostAuthor(db *sql.DB) graphql.FieldResolveFn {
	return func(params graphql.ResolveParams) (interface{}, error) {
		if post, ok := params.Source.(*entities.Post); ok {
			author := &entities.Author{}
			//TODO: fix N+1 problem
			err := db.QueryRow("select id, name, email from authors where id = $1", post.AuthorID).Scan(&author.ID, &author.Name, &author.Email)
			shared.ErrorHandling(err)
			return author, nil
		}
		return nil, nil
	}
}
