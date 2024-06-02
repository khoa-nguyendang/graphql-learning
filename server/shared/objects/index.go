package objects

import (
	"database/sql"

	"github.com/graphql-go/graphql"
)

const (
	AuthorType string = "author"
	PostType   string = "post"
)

type Object struct {
	objectTypes map[string]*graphql.Object
}

func NewObject(db *sql.DB) *Object {
	authorType := InitAuthorType()
	postType := InitPostType(db, authorType)
	objectTypes := map[string]*graphql.Object{AuthorType: authorType, PostType: postType}
	return &Object{
		objectTypes: objectTypes,
	}
}

func (sc *Object) GetTypes() map[string]*graphql.Object {
	return sc.objectTypes
}

func BuildField(name string, output graphql.Output, description string, args graphql.FieldConfigArgument, resolver graphql.FieldResolveFn) *graphql.Field {
	return &graphql.Field{
		Type:        output,
		Description: description,
		Args:        args,
		Resolve:     resolver,
	}
}
