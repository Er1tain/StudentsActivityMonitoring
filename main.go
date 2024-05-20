package main

import (
	handlers "EventsList/Handlers"
	"net/http"
)

func main() {
	http.HandleFunc("getListCategories/", handlers.GetListCategories)

	http.HandleFunc("getListEvents/", handlers.GetListEvents)

	http.HandleFunc("getListStudents/", handlers.GetListStudents)

	http.HandleFunc("newCategory/", handlers.PostNewCategory)

	http.HandleFunc("newEvent/", handlers.PostNewEvent)

	http.HandleFunc("/DeleteEvent", handlers.DeleteEvent)

	http.HandleFunc("/DeleteCategory", handlers.DeleteCategory)

	http.ListenAndServe(":8080", nil)
}
