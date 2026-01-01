package main

import (
	"log"
	"net/http"

	"github.com/pauba98/meal-planner-api/internal/handlers"
	"github.com/pauba98/meal-planner-api/internal/repository"
	"github.com/pauba98/meal-planner-api/internal/services"

	"github.com/go-chi/chi/v5"
)

func main() {

	repo := repository.NewRecipeMockRepo()
	service := services.NewRecipeService(repo)
	handler := handlers.NewRecipeHandler(service)

	r := chi.NewRouter()
	r.Get("/recipes", handler.GetAll)
	// r.Get("/recipes/{id}", handler.GetByID)
	r.Post("/recipe", handler.Create)
	// r.Put("/recipe/{id}", handler.Update)
	// r.Delete("/recipes/{id}", handler.Delete)

	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))

	http.ListenAndServe(":8080", r)
}
