package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"

	"github.com/chagasVinicius/apollo/app/apollo/handlers"
	"github.com/chagasVinicius/apollo/web"
	"github.com/chagasVinicius/apollo/domain/database"
)

func main() {
	router := httprouter.New()

	web.Load()
	database.Load()

	//GET
	router.GET("/hello/:name", handlers.Hello)

	//Post
	router.POST("/categories/:name", handlers.AddCategory)


	fmt.Println("Starting server on :8080 - Test")

	http.ListenAndServe(":8080", router)
}
