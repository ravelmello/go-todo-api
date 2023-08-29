package main

import (
	"fmt"
	"main/configs"
	"main/handlers/todo"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	err := configs.Load()

	if err != nil {
		panic(err)
	}

	router := chi.NewRouter()

	router.Post("/create", handlers.Create)
	router.Put("/{id}", handlers.Update)
	router.Delete("/{id}", handlers.Delete)
	router.Get("/", handlers.List)
	router.Get("/{id}", handlers.GetOne)

	notStarted := http.ListenAndServe(fmt.Sprintf(":%s", configs.GetServerPort()), router)
	if notStarted != nil {
		panic("The server not started ok")
	}

}
