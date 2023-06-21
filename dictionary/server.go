package main

import (
	"dictionary/controller"
	"dictionary/repository"
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"log"
	"net/http"
)

type Server struct {
	Repository *repository.Repository
}

func (s *Server) Listen() {

	router := chi.NewRouter()
	router.Use(middleware.Logger)

	router.Use(func(handler http.Handler) http.Handler {
		return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
			writer.Header().Set("Content-Type", "application/json")
			handler.ServeHTTP(writer, request)
		})
	})

	controller := controller.Controller{Repository: s.Repository}

	router.Get("/hello", func(writer http.ResponseWriter, request *http.Request) {

		response := map[string]string{
			"message": "Hello World!",
		}
		json.NewEncoder(writer).Encode(response)
	})

	router.Post("/definition", controller.Create)
	router.Get("/definition/{word}", controller.Get)

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal(err)
	}
}
