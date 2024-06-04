package resolvers

import (
	"database/sql"
	"server/entities"
	"server/shared"
	"time"

	"github.com/graphql-go/graphql"
)

func CreatePost(db *sql.DB) graphql.FieldResolveFn {
	return func(params graphql.ResolveParams) (interface{}, error) {
		title, _ := params.Args["title"].(string)
		content, _ := params.Args["content"].(string)
		authorId, _ := params.Args["author_id"].(int)
		createdAt := time.Now()

		var lastInsertId int
		err := db.QueryRow("INSERT INTO posts(title, content, author_id, created_at) VALUES($1, $2, $3, $4) returning id;", title, content, authorId, createdAt).Scan(&lastInsertId)
		shared.ErrorHandling(err)

		newPost := &entities.Post{
			ID:        lastInsertId,
			Title:     title,
			Content:   content,
			AuthorID:  authorId,
			CreatedAt: createdAt,
		}

		return newPost, nil
	}
}

func UpdatePost(db *sql.DB) graphql.FieldResolveFn {
	return func(params graphql.ResolveParams) (interface{}, error) {
		id, _ := params.Args["id"].(int)
		title, _ := params.Args["title"].(string)
		content, _ := params.Args["content"].(string)
		authorId, _ := params.Args["author_id"].(int)

		stmt, err := db.Prepare("UPDATE posts SET title = $1, content = $2, author_id = $3 WHERE id = $4")
		shared.ErrorHandling(err)

		_, err2 := stmt.Exec(title, content, authorId, id)
		shared.ErrorHandling(err2)

		newPost := &entities.Post{
			ID:       id,
			Title:    title,
			Content:  content,
			AuthorID: authorId,
		}

		return newPost, nil
	}
}

func DeletePost(db *sql.DB) graphql.FieldResolveFn {
	return func(params graphql.ResolveParams) (interface{}, error) {
		id, _ := params.Args["id"].(int)

		stmt, err := db.Prepare("DELETE FROM posts WHERE id = $1")
		shared.ErrorHandling(err)

		_, err2 := stmt.Exec(id)
		shared.ErrorHandling(err2)

		return nil, nil
	}
}
