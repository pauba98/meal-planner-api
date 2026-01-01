package repository

import (
	"github.com/google/uuid"
	"github.com/pauba98/meal-planner-api/internal/models"
)

// define la interfaz de acceso a datos
type RecipeRepository interface {
	GetAll() ([]models.Recipe, error)
	Create(recipe models.Recipe) models.Recipe
}

// RecipeMockRepo implementa RecipeRepository con datos en memoria
type RecipeMockRepo struct {
	recipes map[uuid.UUID]models.Recipe
}

// Constructor
func NewRecipeMockRepo() RecipeRepository {
	return &RecipeMockRepo{
		recipes: make(map[uuid.UUID]models.Recipe),
	}
}

// GetAll devuelve las recetas de prueba
func (r *RecipeMockRepo) GetAll() ([]models.Recipe, error) {
	all := make([]models.Recipe, 0, len(r.recipes))
	for _, recipe := range r.recipes {
		all = append(all, recipe)
	}
	return all, nil
}

// GetAll devuelve las recetas de prueba
func (r *RecipeMockRepo) Create(recipe models.Recipe) models.Recipe {
	if recipe.ID == uuid.Nil {
		recipe.ID = uuid.New()
	}
	r.recipes[recipe.ID] = recipe
	return recipe
}
