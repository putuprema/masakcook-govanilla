package viewmodel

import recipe "masakcook/internal/business"

type IndexViewModel struct {
	RecipeOfTheDay  recipe.Recipe
	TrendingRecipes []recipe.Recipe
	Categories      []string
}
