package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/pauba98/meal-planner-api/internal/services"
)

// RecipeHandler conecta HTTP con el service
type RecipeHandler struct {
	service *services.RecipeService
}

// Constructor
func NewRecipeHandler(service *services.RecipeService) *RecipeHandler {
	return &RecipeHandler{service: service}
}

// GetAll maneja GET /recipes
func (h *RecipeHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	recipes, err := h.service.GetRecipes()
	if err != nil {
		http.Error(w, "Error getting recipes", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(recipes)
}
