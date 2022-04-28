package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/graphql-go/graphql"
	"github.com/mirogindev/graphqlgo-gin-handler/pkg/handler"
	"log"
)

func StartServer() {
	r := gin.Default()
	r.Use(Midleware3())
	r.Use(Midleware2())
	r.Use(Midleware1())
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	h := handler.New(&handler.Config{
		Schema:     getTestSchema(),
		Pretty:     true,
		Playground: true,
	})
	h.Bind(r, "/graphql")

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func getTestSchema() *graphql.Schema {
	// Schema
	fields := graphql.Fields{
		"hello": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return "world", nil
			},
		},
	}
	rootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: fields}
	schemaConfig := graphql.SchemaConfig{Query: graphql.NewObject(rootQuery)}
	schema, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		log.Fatalf("failed to create new schema, error: %v", err)
	}
	return &schema
}

func Midleware1() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("1 before")
		c.Next()
		fmt.Println("1 after")
	}
}

func Midleware2() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("2 before")

		c.Next()

		fmt.Println("2 after")
	}
}

func Midleware3() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("3 before")

		c.Next()

		fmt.Println("3 after")
	}
}
