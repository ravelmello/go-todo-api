package handlers

import (
	"encoding/json"
	"log"
	models "main/models/todo"
	"net/http"
)

func List(res http.ResponseWriter, req *http.Request) {
	todos, err := models.GetAll()

	if err != nil {
		log.Printf("Error: %v", err)
		http.Error(res, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	res.Header().Add("Content-Type", "application/json")
	json.NewEncoder(res).Encode(todos)

}
