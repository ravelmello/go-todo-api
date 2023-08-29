package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	models "main/models/todo"
	"net/http"
)

func Create(res http.ResponseWriter, req *http.Request) {
	var todo models.Todo

	err := json.NewDecoder(req.Body).Decode(&todo)

	if err != nil {
		log.Printf("Error: %v", err)
		http.Error(res, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	id, err := models.Insert(todo)

	var response map[string]any

	if err != nil {
		response = map[string]any{
			"error":         true,
			"error_message": fmt.Sprintf("Error when insert %v", err.Error()),
		}
	} else {
		response = map[string]any{
			"error":         false,
			"error_message": fmt.Sprintf("Inserted %d", id),
		}
	}

	res.Header().Add("Content-Type", "application/json")
	json.NewEncoder(res).Encode(response)

}
