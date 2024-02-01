export interface Ingredient {
  id: string;
  name: string;
  abv: number;
}

export interface Recipe {
  id: string;
  ingredients: Ingredient[];
}

export interface Cocktail {
  id: string;
  name: string;
  totalAlcohol: number;
  onMenu: boolean;
}

export interface RecipeIngredient {
  amount: number;
  unit: string;
  name: string;
}
