package main

import (
	"io"
	"log"
	"net/http"
	"os"

	"server/mutations"
	"server/queries"
	"server/shared"
	"server/shared/objects"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	_ "github.com/lib/pq"
)

func init() {
	log.SetOutput(os.Stdout)
	log.SetFlags(log.Lshortfile)
}

func main() {
	cfg := loadConfig()
	db := GetDB(cfg)
	RunMigration(db)
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
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) { io.WriteString(w, `pong`) })
	http.ListenAndServe(":8080", nil)
}

func loadConfig() *Config {
	cfg, err := GetConfig("/app/config")
	shared.CheckErr(err)
	return cfg
}
