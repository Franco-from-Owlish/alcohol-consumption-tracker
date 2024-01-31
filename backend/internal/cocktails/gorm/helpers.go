package gorm

import (
	"alcohol-consumption-tracker/internal/cocktaildb"
	cocktail "alcohol-consumption-tracker/internal/cocktails"
	"fmt"
)

func (s *CocktailsService) newCocktailFromCocktailDB(data *cocktaildb.Cocktail) *cocktail.Cocktail {
	c := cocktail.Cocktail{
		Name:         data.Name,
		TotalAlcohol: 0,
		OnMenu:       false,
	}
	s.DB.Save(&c)
	r := cocktail.Recipe{CocktailID: c.ID}
	s.DB.Save(&r)
	s.updateRecipeFromIngredientMap(&r, data.Ingredients)
	return &c
}

func (s *CocktailsService) updateRecipeFromIngredientMap(recipe *cocktail.Recipe, data map[string]cocktaildb.IngredientMeasurement) {
	var ingredients []cocktail.Ingredient
	for name, measurement := range data {
		ingredient, err := s.newIngredientFromIngredientMeasurement(name)
		if err != nil {
			fmt.Printf("\nCould not save ingredient %s: %v\n", name, err)
		}
		s.DB.Save(&ingredient)
		ingredients = append(ingredients, ingredient)
		s.DB.Save(&cocktail.RecipeIngredient{
			RecipeID:     int(recipe.ID),
			IngredientID: int(ingredient.ID),
			Amount:       measurement.Value,
			Unit:         measurement.Unit,
		})
	}
	s.DB.Save(&recipe)
}

func (s *CocktailsService) newIngredientFromIngredientMeasurement(
	name string) (cocktail.Ingredient, error) {
	ingredient, errFetch := s.CocktailDB.GetIngredient(name)
	if errFetch != nil {
		return cocktail.Ingredient{}, errFetch
	}
	return cocktail.Ingredient{
		Name: ingredient.Name,
		Abv:  ingredient.Abv * 100,
	}, nil
}
