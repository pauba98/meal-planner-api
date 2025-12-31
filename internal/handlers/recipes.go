package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/pauba98/meal-planner-api/internal/models"
)

func GetRecipes(w http.ResponseWriter, r *http.Request) {
	recipes := []models.Recipe{
		{ID: uuid.New(), Name: "Pasta rápida", TimeMin: 30, Tags: []string{"rápido"}},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(recipes)
}
