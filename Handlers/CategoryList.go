package handlers

import (
	"EventsList/Models/Mocks"
	"encoding/json"
	"net/http"
)

func GetListCategories(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		database := Mocks.ConnectToDB()
		list_catigories := database.GetCatigories()

		data, err := json.Marshal(list_catigories)
		if err != nil {
			panic("Fatal error!")
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(data)
		return
	}

	data, _ := json.Marshal("Only GET!")
	w.Write(data)
}
