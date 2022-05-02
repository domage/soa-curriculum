package main

import (
	"context"
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/go-chi/chi"

	"nuuka.com/go-gql/generated"
	"nuuka.com/go-gql/graph"
)

func main() {
	router := chi.NewRouter()

	srv := handler.New(generated.NewExecutableSchema(graph.New()))
	srv.AddTransport(transport.POST{})
	srv.AroundOperations(func(ctx context.Context, next graphql.OperationHandler) graphql.ResponseHandler {
		rc := graphql.GetOperationContext(ctx)
		rc.DisableIntrospection = false

		return next(ctx)
	})

	router.Handle("/graphql", srv)

	log.Println("Server running at *:4000")
	log.Fatal(http.ListenAndServe(":4000", router))
}
