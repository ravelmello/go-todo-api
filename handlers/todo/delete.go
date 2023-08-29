package handlers

import (
	"encoding/json"
	"log"
	models "main/models/todo"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func Delete(res http.ResponseWriter, req *http.Request) {

	id, err := strconv.Atoi(chi.URLParam(req, "id"))

	if err != nil {
		log.Printf("Error: %v", err)
		http.Error(res, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	rows, err := models.Delete(int64(id))

	if rows > 1 {
		log.Panicf("Number of rows deleted is wrong ", rows)
	}

	resp := map[string]any{
		"Error":   false,
		"Message": "Deleted successfully",
	}

	res.Header().Add("Content-Type", "application/json")
	json.NewEncoder(res).Encode(resp)

}
