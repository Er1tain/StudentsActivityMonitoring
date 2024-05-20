package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func Post_handler(w http.ResponseWriter, r *http.Request) {

	type User struct {
		Id_user  string
		Nickname string
		Age      int
	}

	decoder := json.NewDecoder(r.Body)
	var t User
	err := decoder.Decode(&t)
	if err != nil {
		panic(err)
	}

	fmt.Println(t.Id_user)

}
