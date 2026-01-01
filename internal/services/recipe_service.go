package services

import (
	"github.com/pauba98/meal-planner-api/internal/models"
	"github.com/pauba98/meal-planner-api/internal/repository"
)

// RecipeService maneja la lógica de negocio
type RecipeService struct {
	repo repository.RecipeRepository
}

// Constructor
func NewRecipeService(repo repository.RecipeRepository) *RecipeService {
	return &RecipeService{repo: repo}
}

// GetRecipes devuelve todas las recetas
func (s *RecipeService) GetRecipes() ([]models.Recipe, error) {
	return s.repo.GetAll()
}

func (s *RecipeService) CreateRecipe(recipe models.Recipe) (models.Recipe, error) {
	created := s.repo.Create(recipe)
	return created, nil
}
