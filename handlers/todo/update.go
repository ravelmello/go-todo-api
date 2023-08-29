package handlers

import (
	"encoding/json"
	"log"
	models "main/models/todo"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func Update(res http.ResponseWriter, req *http.Request) {

	var todo models.Todo

	id, err := strconv.Atoi(chi.URLParam(req, "id"))

	if err != nil {
		log.Printf("Error: %v", err)
		http.Error(res, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	rows, err := models.Update(int64(id), todo)
	if err != nil {
		log.Printf("Error: %v", err)
		http.Error(res, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	if rows > 1 {
		log.Panicf("Number of rows updated is wrong ", rows)
	}

	resp := map[string]any{
		"Error":   false,
		"Message": "Updated successfully",
	}

	res.Header().Add("Content-Type", "application/json")
	json.NewEncoder(res).Encode(resp)

}
