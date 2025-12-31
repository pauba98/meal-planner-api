package repository

import (
	"github.com/google/uuid"
	"github.com/pauba98/meal-planner-api/internal/models"
)

// define la interfaz de acceso a datos
type RecipeRepository interface {
	GetAll() ([]models.Recipe, error)
}

// RecipeMockRepo implementa RecipeRepository con datos en memoria
type RecipeMockRepo struct{}

// Constructor
func NewRecipeMockRepo() RecipeRepository {
	return &RecipeMockRepo{}
}

// GetAll devuelve las recetas de prueba
func (r *RecipeMockRepo) GetAll() ([]models.Recipe, error) {
	recipes := []models.Recipe{
		{ID: uuid.New(), Name: "Pasta rápida", TimeMin: 30, Tags: []string{"rápido"}},
	}
	return recipes, nil
}
