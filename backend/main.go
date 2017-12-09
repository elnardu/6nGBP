package main

import (
	"log"
	"net/http"
	"os"

	"github.com/auth0/go-jwt-middleware"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/elnardu/6nGBP/backend/model"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
)

func configureSchema() graphql.Schema {
	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query:    model.QueryType,
		Mutation: model.MutationType,
	})
	if err != nil {
		log.Fatal(err)
	}
	return schema
}

// func handler(schema graphql.Schema) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		query, err := ioutil.ReadAll(r.Body)
// 		log.Println(query)
// 		if err != nil {
// 			http.Error(w, err.Error(), http.StatusBadRequest)
// 			return
// 		}
// 		result := graphql.Do(graphql.Params{
// 			Schema:        schema,
// 			RequestString: string(query),
// 		})
// 		w.Header().Set("Content-Type", "application/json")
// 		w.WriteHeader(http.StatusOK)
// 		json.NewEncoder(w).Encode(result)
// 	}
// }

// func addContext(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		authContext := map[string]interface{}{
// 			"key1": 23,
// 			"key2": "another value",
// 		}
// 		ctx := context.WithValue(r.Context(), "!@3", authContext)
// 		next.ServeHTTP(w, r.WithContext(ctx))
// 	})
// }

func main() {
	log.Println("Starting server...")
	log.Println("Database: " + os.Getenv("DBCONNECT"))
	model.NewDB(os.Getenv("DBCONNECT"))
	schema := configureSchema()

	h := handler.New(&handler.Config{
		Schema:   &schema,
		Pretty:   true,
		GraphiQL: true,
	})

	jwtMiddleware := jwtmiddleware.New(jwtmiddleware.Options{
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("SECRET")), nil
		},
		SigningMethod:       jwt.SigningMethodHS256,
		CredentialsOptional: true,
	})

	http.Handle("/api/graphql", jwtMiddleware.Handler(h))
	// http.Handle("/graphql", handler(schema))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
