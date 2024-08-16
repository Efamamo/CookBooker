package controller

import "example/htmx/domain"

type IRecipeUsecase interface {
	GetRecipes() (*[]domain.Recipe, error)
	GetRecipeById(id string) (*domain.Recipe, error)
	UpdateRecipe(id string, Title string, Ingredients string, Instructions []string, ImageURL string) error
	AddRecipe(recipe domain.Recipe) (*domain.Recipe, error)
	DeleteRecipe(id string) error
}
