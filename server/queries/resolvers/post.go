package resolvers

import (
	"database/sql"
	"server/entities"
	"server/shared"

	"github.com/graphql-go/graphql"
)

func Post(db *sql.DB) graphql.FieldResolveFn {
	return func(params graphql.ResolveParams) (interface{}, error) {
		id, _ := params.Args["id"].(int)

		post := &entities.Post{}
		err := db.QueryRow("select id, title, content, author_id from posts where id = $1", id).Scan(&post.ID, &post.Title, &post.Content, &post.AuthorID)
		shared.CheckErr(err)

		return post, nil
	}
}

func Posts(db *sql.DB) graphql.FieldResolveFn {
	return func(params graphql.ResolveParams) (interface{}, error) {
		rows, err := db.Query("SELECT id, title, content, author_id FROM posts")
		shared.CheckErr(err)
		var posts []*entities.Post

		for rows.Next() {
			post := &entities.Post{}

			err = rows.Scan(&post.ID, &post.Title, &post.Content, &post.AuthorID)
			shared.CheckErr(err)
			posts = append(posts, post)
		}

		return posts, nil
	}
}

// PostAuthor this is different params field from query, check shared/objects/post.go "author": &graphql.Field
func PostAuthor(db *sql.DB) graphql.FieldResolveFn {
	return func(params graphql.ResolveParams) (interface{}, error) {
		if post, ok := params.Source.(*entities.Post); ok {
			author := &entities.Author{}
			err := db.QueryRow("select id, name, email from authors where id = $1", post.AuthorID).Scan(&author.ID, &author.Name, &author.Email)
			shared.CheckErr(err)
			return author, nil
		}
		return nil, nil
	}
}
