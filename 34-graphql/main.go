package main

import (
	"log"
	"net/http"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
)

var authors = []Author{
	Author{ID: 1, Name: "Author 1"},
	Author{ID: 2, Name: "Author 2"},
}

var books = []Book{
	Book{ID: 1, Title: "Book 1", Author: 1},
	Book{ID: 2, Title: "Book 2", Author: 2},
}

var authorType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Author",
		Fields: graphql.Fields{
			"ID":   &graphql.Field{Type: graphql.Int},
			"Name": &graphql.Field{Type: graphql.String},
		},
	},
)

var bookType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Book",
		Fields: graphql.Fields{
			"ID":    &graphql.Field{Type: graphql.Int},
			"Title": &graphql.Field{Type: graphql.String},
			"Author": &graphql.Field{
				Type: authorType,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					book := p.Source.(Book)
					for _, author := range authors {
						if author.ID == book.Author {
							return author, nil
						}
					}
					return nil, nil
				},
			},
		},
	},
)

var queryType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"book": &graphql.Field{
				Type: bookType,
				Args: graphql.FieldConfigArgument{
					"ID": &graphql.ArgumentConfig{Type: graphql.Int},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					id, _ := p.Args["ID"].(int)
					for _, book := range books {
						if book.ID == id {
							return book, nil
						}
					}
					return nil, nil
				},
			},
		},
	},
)

func main() {
	schema, err := graphql.NewSchema(
		graphql.SchemaConfig{
			Query: queryType,
		},
	)

	if err != nil {
		log.Fatalf("Error creating schema: %v", err)
	}

	h := handler.New(&handler.Config{
		Schema: &schema,
		Pretty: true,
	})

	http.Handle("/graphql", h)
	log.Println("Server started at http://localhost:8080/graphql")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
