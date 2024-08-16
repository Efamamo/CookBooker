package usecase

import "example/htmx/domain"

type IRecipeRepo interface {
	FindAll() (*[]domain.Recipe, error)
	FindOne(id string) (*domain.Recipe, error)
	UpdateOne(id string, Title string, Ingredients string, Instructions []string, ImageURL string) error
	Save(newRecipe domain.Recipe) (*domain.Recipe, error)
	DeleteOne(id string) error
}
