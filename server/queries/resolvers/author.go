package resolvers

import (
	"database/sql"

	"server/entities"
	"server/shared"

	"github.com/graphql-go/graphql"
)

// TODO: use BIZ or Repository here

// Author
func Author(db *sql.DB) graphql.FieldResolveFn {
	return func(params graphql.ResolveParams) (interface{}, error) {
		id, _ := params.Args["id"].(int)
		author := &entities.Author{}
		err := db.QueryRow("select id, name, email from authors where id = $1", id).Scan(&author.ID, &author.Name, &author.Email)
		shared.CheckErr(err)
		return author, nil
	}
}

// Authors
func Authors(db *sql.DB) graphql.FieldResolveFn {
	return func(params graphql.ResolveParams) (interface{}, error) {
		rows, err := db.Query("SELECT id, name, email FROM authors")
		shared.CheckErr(err)
		var authors []*entities.Author

		for rows.Next() {
			author := &entities.Author{}

			err = rows.Scan(&author.ID, &author.Name, &author.Email)
			shared.CheckErr(err)
			authors = append(authors, author)
		}

		return authors, nil
	}
}
