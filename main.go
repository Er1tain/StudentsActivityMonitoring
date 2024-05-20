package main

import (
	handlers "EventsList/Handlers"
	"net/http"
)

func main() {

	http.HandleFunc("/post_handler", handlers.Post_handler)

	http.ListenAndServe(":8080", nil)

}
