package handlers

import (
	"encoding/json"
	"log"
	models "main/models/todo"
	"net/http"
	"strconv"
	"github.com/go-chi/chi/v5"
)

func GetOne(res http.ResponseWriter, req *http.Request) {

	id, err := strconv.Atoi(chi.URLParam(req, "id"))

	if err != nil {
		log.Printf("Error: %v", err)
		http.Error(res, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	todo, err := models.Get(int64(id))

	if err != nil {
		log.Printf("Error: %v", err)
		http.Error(res, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	res.Header().Add("Content-Type", "application/json")
	json.NewEncoder(res).Encode(todo)
}
