package main

import (
	"net/http"

	"server/mutations"
	"server/queries"
	"server/shared/objects"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	_ "github.com/lib/pq"
)

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	db := GetDB()
	defer db.Close()
	objSvc := objects.NewObject(db)
	types := objSvc.GetTypes()
	querier := queries.NewQuerier(db, types)
	mutator := mutations.InitMutator(db, types)
	schema, _ := graphql.NewSchema(graphql.SchemaConfig{
		Query:    querier.GetRootQuery(),
		Mutation: mutator.GetRootMutation(),
	})

	h := handler.New(&handler.Config{
		Schema:   &schema,
		Pretty:   true,
		GraphiQL: true,
	})

	// serve HTTP
	http.Handle("/graphql", h)
	http.ListenAndServe(":28080", nil)
}
