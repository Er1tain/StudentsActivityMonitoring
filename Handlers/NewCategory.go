package handlers

import (
	"EventsList/Models/Mocks"
	"EventsList/Models/Mocks/serialize_struct"
	"encoding/json"
	"net/http"
)

type data_for_init_categories = serialize_struct.Data_for_init_categories

func PostNewCategory(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		data, _ := json.Marshal("Only POST!")
		w.Write(data)
		return
	}

	//Десериализуем json
	decoder := json.NewDecoder(r.Body)
	var request_data data_for_init_categories
	err := decoder.Decode(&request_data)
	if err != nil {
		panic("Fatal error!")
	}

	database := Mocks.ConnectToDB()
	status_operation := database.NewCatigory(request_data)

	w.Header().Set("Content-Type", "application/json")

	//Сериализуемые данные
	type Operation struct {
		Status_operation string
	}

	//Если запрос к БД успешен
	if status_operation {
		operation, _ := json.Marshal(Operation{Status_operation: "succes"})
		w.Write(operation)
		return
	}

	fail, _ := json.Marshal("Sometthing broke(")
	w.Write(fail)
}
