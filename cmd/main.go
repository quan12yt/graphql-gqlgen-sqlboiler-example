package main

import (
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/quan12yt/graphql-sqlboiler-example/cmd/graph"
	"github.com/quan12yt/graphql-sqlboiler-example/cmd/graph/generated"
	"github.com/quan12yt/graphql-sqlboiler-example/db"
	"github.com/quan12yt/graphql-sqlboiler-example/internal/service/users"
	"github.com/volatiletech/sqlboiler/boil"
)

func main() {
	db := db.DBConnection()
	boil.DebugMode = true

	us := users.NewUserServiceImpl(db)

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(graph.NewSchemaConfig(us)))

	http.Handle("/playground", playground.Handler("GraphQL playground", "/graphql"))
	http.Handle("/graphql", srv)

	log.Println("server running on port :8080")
	log.Println("graphql playground running on :8080/playground")

	log.Fatal(http.ListenAndServe(":8080", nil))
}
