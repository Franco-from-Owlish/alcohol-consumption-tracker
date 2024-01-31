package gorm

import (
	cocktaildb "alcohol-consumption-tracker/internal/cocktaildb/redis"
	cocktail "alcohol-consumption-tracker/internal/cocktails"
	"fmt"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var _ cocktail.CocktailsService = (*CocktailsService)(nil)

type CocktailsService struct {
	DB                *gorm.DB
	CocktailDB        cocktaildb.CocktailDbService
	IngredientService cocktail.IngredientsService
}

func NewCocktailsService(db *gorm.DB, datastore *redis.Client) *CocktailsService {
	cocktailDbService := cocktaildb.NewCocktailDbService(datastore)
	ingredientService := NewIngredientsService(db, datastore)
	err := db.SetupJoinTable(&cocktail.Recipe{}, "Ingredients", &cocktail.RecipeIngredient{})
	if err != nil {
		panic(fmt.Sprintf("could not setup join table: %v", err))
	}

	return &CocktailsService{
		DB:                db,
		CocktailDB:        *cocktailDbService,
		IngredientService: ingredientService,
	}
}

// GetAll Returns all cocktails on the menu
func (s *CocktailsService) GetAll() (*[]cocktail.Cocktail, error) {
	var cocktails []cocktail.Cocktail
	err := s.DB.Where("on_menu IS True").Find(&cocktails).Error
	return &cocktails, err
}

// GetRandom Fetches a random cocktails and adds it to the database
func (s *CocktailsService) GetRandom() (*cocktail.Cocktail, error) {
	random, errRdm := s.CocktailDB.GetRandomCocktail()
	if errRdm != nil {
		return nil, errRdm
	}
	randomCocktail := s.newCocktailFromCocktailDB(random)
	errSave := s.DB.Save(&randomCocktail).Error
	return randomCocktail, errSave
}

func (s *CocktailsService) GetByName(name string) (*cocktail.Cocktail, error) {
	var drink cocktail.Cocktail
	err := s.DB.Find(&drink, "name ILIKE ?", name).Error
	return &drink, err
}

func (s *CocktailsService) GetCocktailIngredients(data *cocktail.Cocktail) error {
	return s.DB.Preload("Recipe").Preload("Recipe.Ingredients").
		Find(&data).Error
}

func (s *CocktailsService) GetCocktailRecipe(data *cocktail.Cocktail) ([]map[string]interface{}, error) {
	errRcp := s.DB.Preload("Recipe").Find(&data).Error
	if errRcp != nil {
		return nil, errRcp
	}
	var recipe []map[string]interface{}
	err := s.DB.Table("recipe_ingredients").
		Select("ingredients.name, recipe_ingredients.amount, recipe_ingredients.unit").
		Joins("join ingredients on recipe_ingredients.ingredient_id = ingredients.id").
		Where("recipe_ingredients.recipe_id = ?", data.Recipe.ID).
		Find(&recipe).Error
	if err != nil {
		return nil, err
	}
	fmt.Printf("\nrecipe: %v\n", recipe)
	return recipe, nil
}

func (s *CocktailsService) AddToMenu(name string) error {
	drink, errDrink := s.GetByName(name)
	if errDrink != nil {
		return errDrink
	}
	drink.OnMenu = true
	return s.DB.Save(&drink).Error
}

func (s *CocktailsService) UpdateCocktailAlcoholContent(data *cocktail.Cocktail) error {
	errRcp := s.GetCocktailIngredients(data)
	if errRcp != nil {
		return errRcp
	}

	var sum float32 = 0.00
	for _, ing := range data.Recipe.Ingredients {
		recipeIngredient := cocktail.RecipeIngredient{
			RecipeID:     int(data.Recipe.ID),
			IngredientID: int(ing.ID),
		}
		errRI := s.DB.Find(&recipeIngredient).Error
		if errRI != nil {
			return errRI
		}
		sum += ing.Abv / 100 * recipeIngredient.Amount
	}

	data.TotalAlcohol = sum
	return s.DB.Save(&data).Error
}
