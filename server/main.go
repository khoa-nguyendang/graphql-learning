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
	http.Handle("/graphql", disableCors(h))
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) { io.WriteString(w, `pong`) })
	http.ListenAndServe(":8080", nil)
}

func loadConfig() *Config {
	cfg, err := GetConfig("/app/config")
	shared.ErrorHandling(err)
	return cfg
}

func EnableCors(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST,OPTIONS")
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}
}

func disableCors(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Authorization, Content-Type, Content-Length, Accept-Encoding")

		// I added this for another handler of mine,
		// but I do not think this is necessary for GraphQL's handler
		if r.Method == "OPTIONS" {
			w.Header().Set("Access-Control-Max-Age", "86400")
			w.WriteHeader(http.StatusOK)
			return
		}

		h.ServeHTTP(w, r)
	})
}
