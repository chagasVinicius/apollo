package main

import (
	"fmt"
	"net/http"
	"github.com/julienschmidt/httprouter"

	"handlers"
)

func main() {
	router := httprouter.New()

	//Post
	router.POST("/music/classify/:name", handlers.MusicClassifyHandler)

	fmt.Println("Starting server on :8080")

	http.ListenAndServe(":8080", router)
}
