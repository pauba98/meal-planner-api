package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"

	"github.com/pauba98/meal-planner-api/internal/models"
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

func (h *RecipeHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	recipes, err := h.service.GetRecipes()
	if err != nil {
		http.Error(w, "Error getting recipes", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(recipes)
}

func (h *RecipeHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	recipe, err := h.service.GetByID(id)
	if err != nil {
		http.Error(w, "recipe not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(recipe)
}

func (h *RecipeHandler) Create(w http.ResponseWriter, r *http.Request) {
	var recipe models.Recipe
	if err := json.NewDecoder(r.Body).Decode(&recipe); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}
	created, err := h.service.CreateRecipe(recipe)
	if err != nil {
		http.Error(w, "Error creating recipe", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(created)
}

func (h *RecipeHandler) Update(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}
	var recipe models.Recipe
	if err := json.NewDecoder(r.Body).Decode(&recipe); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}
	updated, err := h.service.Update(id, recipe)
	if err != nil {
		http.Error(w, "recipe not found", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updated)
}

func (h *RecipeHandler) Delete(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	if err := h.service.Delete(id); err != nil {
		http.Error(w, "recipe not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
