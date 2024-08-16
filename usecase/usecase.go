package usecase

import "example/htmx/domain"

type RecipeUsecase struct {
	RecipeRepository IRecipeRepo
}

func (ru RecipeUsecase) GetRecipes() (*[]domain.Recipe, error) {
	recipes, err := ru.RecipeRepository.FindAll()

	if err != nil {
		return nil, err
	}

	return recipes, nil
}

func (ru RecipeUsecase) GetRecipeById(id string) (*domain.Recipe, error) {
	recipe, err := ru.RecipeRepository.FindOne(id)

	if err != nil {
		return nil, err
	}

	return recipe, nil
}

func (ru RecipeUsecase) UpdateRecipe(id string, Title string, Ingredients string, Instructions []string, ImageURL string) error {
	err := ru.RecipeRepository.UpdateOne(id, Title, Ingredients, Instructions, ImageURL)

	if err != nil {
		return err
	}
	return nil
}

func (ru RecipeUsecase) AddRecipe(recipe domain.Recipe) (*domain.Recipe, error) {
	recipes, err := ru.RecipeRepository.Save(recipe)

	if err != nil {
		return nil, err
	}

	return recipes, nil
}

func (ru RecipeUsecase) DeleteRecipe(id string) error {
	err := ru.RecipeRepository.DeleteOne(id)
	if err != nil {
		return err
	}
	return nil
}
