package repository

import (
	"errors"

	"github.com/google/uuid"
	"github.com/pauba98/meal-planner-api/internal/models"
)

// define la interfaz de acceso a datos
type RecipeRepository interface {
	GetAll() ([]models.Recipe, error)
	GetById(id uuid.UUID) (models.Recipe, error)
	Create(recipe models.Recipe) models.Recipe
	Update(id uuid.UUID, recipe models.Recipe) (models.Recipe, error)
	Delete(id uuid.UUID) error
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
func (r *RecipeMockRepo) GetById(id uuid.UUID) (models.Recipe, error) {
	recipe, exists := r.recipes[id]
	if !exists {
		return models.Recipe{}, errors.New("recipe not found")
	}
	return recipe, nil
}

// GetAll devuelve las recetas de prueba
func (r *RecipeMockRepo) Create(recipe models.Recipe) models.Recipe {
	if recipe.ID == uuid.Nil {
		recipe.ID = uuid.New()
	}
	r.recipes[recipe.ID] = recipe
	return recipe
}

func (r *RecipeMockRepo) Update(id uuid.UUID, recipe models.Recipe) (models.Recipe, error) {
	_, ok := r.recipes[id]
	if !ok {
		return models.Recipe{}, errors.New("recipe not found")
	}
	recipe.ID = id
	r.recipes[id] = recipe
	return recipe, nil
}

func (r *RecipeMockRepo) Delete(id uuid.UUID) error {
	_, ok := r.recipes[id]
	if !ok {
		return errors.New("recipe not found")
	}
	delete(r.recipes, id)
	return nil
}
