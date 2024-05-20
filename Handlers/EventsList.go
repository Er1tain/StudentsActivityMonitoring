package handlers

import (
	"EventsList/Models/Mocks"
	"encoding/json"
	"net/http"
)

func GetListEvents(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		database := Mocks.ConnectToDB()
		catigory := r.URL.Query().Get("catigory")
		list_events := database.GetEvents(catigory)

		data, err := json.Marshal(list_events)
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
