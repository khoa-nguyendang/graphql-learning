package resolvers

import (
	"database/sql"
	"server/entities"
	"server/shared"
	"time"

	"github.com/graphql-go/graphql"
)

func CreateAuthor(db *sql.DB) graphql.FieldResolveFn {
	return func(params graphql.ResolveParams) (interface{}, error) {
		name, _ := params.Args["name"].(string)
		email, _ := params.Args["email"].(string)
		createdAt := time.Now()

		var lastInsertId int
		err := db.QueryRow("INSERT INTO authors(name, email, created_at) VALUES($1, $2, $3) returning id;", name, email, createdAt).Scan(&lastInsertId)
		shared.CheckErr(err)

		newAuthor := &entities.Author{
			ID:        lastInsertId,
			Name:      name,
			Email:     email,
			CreatedAt: createdAt,
		}

		return newAuthor, nil
	}
}

func UpdateAuthor(db *sql.DB) graphql.FieldResolveFn {
	return func(params graphql.ResolveParams) (interface{}, error) {
		id, _ := params.Args["id"].(int)
		name, _ := params.Args["name"].(string)
		email, _ := params.Args["email"].(string)

		stmt, err := db.Prepare("UPDATE authors SET name = $1, email = $2 WHERE id = $3")
		shared.CheckErr(err)

		_, err2 := stmt.Exec(name, email, id)
		shared.CheckErr(err2)

		newAuthor := &entities.Author{
			ID:    id,
			Name:  name,
			Email: email,
		}
		return newAuthor, nil
	}
}

func DeleteAuthor(db *sql.DB) graphql.FieldResolveFn {
	return func(params graphql.ResolveParams) (interface{}, error) {
		id, _ := params.Args["id"].(int)

		stmt, err := db.Prepare("DELETE FROM authors WHERE id = $1")
		shared.CheckErr(err)

		_, err2 := stmt.Exec(id)
		shared.CheckErr(err2)

		return nil, nil
	}
}
