package data

import (
	"strings"

	recipe "masakcook/internal/business"
)

var mockRecipes = []recipe.Recipe{
	{
		ID:          "1",
		Title:       "Classic Margherita Pizza",
		Description: "A timeless Italian classic with fresh tomatoes, mozzarella, and basil. Simple ingredients come together to create an unforgettable flavor experience.",
		Image:       "https://images.unsplash.com/photo-1604068549290-dea0e4a305ca?w=800&h=600&fit=crop",
		Ingredients: []string{"Pizza dough", "San Marzano tomatoes", "Fresh mozzarella", "Fresh basil", "Extra virgin olive oil", "Salt"},
		Instructions: []string{
			"Preheat oven to 475°F (245°C)",
			"Roll out pizza dough on a floured surface",
			"Spread tomato sauce evenly over the dough",
			"Add torn mozzarella pieces",
			"Bake for 12-15 minutes until golden",
			"Top with fresh basil and drizzle with olive oil",
		},
		PrepTime:   15,
		CookTime:   15,
		Servings:   4,
		Category:   recipe.CategoryDinner,
		Difficulty: recipe.DifficultyEasy,
		Author:     recipe.Author{Name: "Chef Antonio", Avatar: ""},
		Ratings:    recipe.Rating{Average: 4.8, Count: 156},
		Likes:      324,
		Saves:      256,
		CreatedAt:  "2025-09-07",
	},
	{
		ID:          "2",
		Title:       "Fluffy Buttermilk Pancakes",
		Description: "Light, fluffy pancakes that melt in your mouth. Perfect for a lazy Sunday breakfast with maple syrup and fresh berries.",
		Image:       "https://images.unsplash.com/photo-1567620905732-2d1ec7ab7445?w=800&h=600&fit=crop",
		Ingredients: []string{"All-purpose flour", "Buttermilk", "Eggs", "Baking powder", "Sugar", "Salt", "Butter", "Vanilla extract"},
		Instructions: []string{
			"Mix dry ingredients in a large bowl",
			"Whisk wet ingredients in a separate bowl",
			"Combine wet and dry ingredients gently",
			"Heat griddle to medium heat",
			"Pour 1/4 cup batter for each pancake",
			"Flip when bubbles form, cook until golden",
		},
		PrepTime:   10,
		CookTime:   20,
		Servings:   4,
		Category:   recipe.CategoryBreakfast,
		Difficulty: recipe.DifficultyEasy,
		Author:     recipe.Author{Name: "Sarah Mitchell", Avatar: ""},
		Ratings:    recipe.Rating{Average: 4.9, Count: 203},
		Likes:      512,
		Saves:      428,
		CreatedAt:  "2025-10-14",
	},
	{
		ID:          "3",
		Title:       "Spicy Thai Basil Chicken",
		Description: "An authentic Thai street food favorite with aromatic holy basil, chilies, and a perfect balance of sweet, salty, and spicy flavors.",
		Image:       "https://images.unsplash.com/photo-1604908177453-7462950a6a3f?w=800&h=600&fit=crop",
		Ingredients: []string{"Chicken thigh", "Holy basil", "Thai chilies", "Garlic", "Soy sauce", "Oyster sauce", "Fish sauce", "Sugar", "Jasmine rice"},
		Instructions: []string{
			"Mince garlic and slice chilies",
			"Cut chicken into bite-sized pieces",
			"Heat oil in wok over high heat",
			"Stir-fry garlic and chilies until fragrant",
			"Add chicken, cook until done",
			"Add sauces and basil, toss quickly",
			"Serve over jasmine rice with fried egg",
		},
		PrepTime:   15,
		CookTime:   10,
		Servings:   2,
		Category:   recipe.CategoryLunch,
		Difficulty: recipe.DifficultyMedium,
		Author:     recipe.Author{Name: "Nong Pim", Avatar: ""},
		Ratings:    recipe.Rating{Average: 4.7, Count: 178},
		Likes:      445,
		Saves:      389,
		CreatedAt:  "2025-10-02",
	},
	{
		ID:          "4",
		Title:       "Decadent Chocolate Lava Cake",
		Description: "Rich, gooey chocolate cake with a molten center. An impressive dessert that's surprisingly easy to make.",
		Image:       "https://images.unsplash.com/photo-1624353365286-3f8d62daad51?w=800&h=600&fit=crop",
		Ingredients: []string{"Dark chocolate", "Butter", "Eggs", "Sugar", "Flour", "Vanilla extract", "Cocoa powder"},
		Instructions: []string{
			"Preheat oven to 425°F (220°C)",
			"Melt chocolate and butter together",
			"Whisk eggs and sugar until thick",
			"Fold in chocolate mixture",
			"Add flour gently",
			"Pour into buttered ramekins",
			"Bake for 12-14 minutes",
			"Let rest 1 minute, then invert onto plates",
		},
		PrepTime:   15,
		CookTime:   14,
		Servings:   4,
		Category:   recipe.CategoryDessert,
		Difficulty: recipe.DifficultyMedium,
		Author:     recipe.Author{Name: "Pastry Chef Marie", Avatar: ""},
		Ratings:    recipe.Rating{Average: 4.9, Count: 245},
		Likes:      678,
		Saves:      612,
		CreatedAt:  "2025-08-17",
	},
	{
		ID:          "5",
		Title:       "Fresh Guacamole",
		Description: "Creamy, zesty, and perfect for any occasion. This authentic guacamole is chunky and full of flavor.",
		Image:       "https://images.unsplash.com/photo-1553621042-f6e147245754?w=800&h=600&fit=crop",
		Ingredients: []string{"Ripe avocados", "Lime juice", "Red onion", "Tomato", "Cilantro", "Jalapeño", "Salt", "Garlic"},
		Instructions: []string{
			"Halve and pit avocados",
			"Scoop flesh into bowl",
			"Mash to desired consistency",
			"Dice onion, tomato, and jalapeño",
			"Chop cilantro",
			"Mix everything together",
			"Season with lime juice and salt",
			"Serve immediately with chips",
		},
		PrepTime:   10,
		CookTime:   0,
		Servings:   6,
		Category:   recipe.CategoryAppetizer,
		Difficulty: recipe.DifficultyEasy,
		Author:     recipe.Author{Name: "Carlos Rodriguez", Avatar: ""},
		Ratings:    recipe.Rating{Average: 4.8, Count: 189},
		Likes:      534,
		Saves:      467,
		CreatedAt:  "2025-10-10",
	},
	{
		ID:          "6",
		Title:       "Loaded Nachos Supreme",
		Description: "Crispy tortilla chips piled high with melted cheese, seasoned beef, and all your favorite toppings.",
		Image:       "https://images.unsplash.com/photo-1513456852971-30c0b8199d4d?w=800&h=600&fit=crop",
		Ingredients: []string{"Tortilla chips", "Ground beef", "Cheddar cheese", "Black beans", "Jalapeños", "Sour cream", "Salsa", "Guacamole"},
		Instructions: []string{
			"Brown ground beef with taco seasoning",
			"Spread chips on baking sheet",
			"Layer: chips, beef, beans, cheese",
			"Bake at 375°F for 10 minutes",
			"Top with jalapeños, sour cream, salsa",
			"Add guacamole and serve hot",
		},
		PrepTime:   15,
		CookTime:   15,
		Servings:   4,
		Category:   recipe.CategorySnacks,
		Difficulty: recipe.DifficultyEasy,
		Author:     recipe.Author{Name: "Mike Johnson", Avatar: ""},
		Ratings:    recipe.Rating{Average: 4.7, Count: 167},
		Likes:      589,
		Saves:      501,
		CreatedAt:  "2025-10-17",
	},
	{
		ID:          "7",
		Title:       "Homemade Iced Coffee",
		Description: "Smooth, cold-brewed coffee with just the right amount of sweetness. Better than any coffee shop.",
		Image:       "https://images.unsplash.com/photo-1517487881594-2787fef5ebf7?w=800&h=600&fit=crop",
		Ingredients: []string{"Coarse ground coffee", "Cold water", "Milk or cream", "Simple syrup", "Ice"},
		Instructions: []string{
			"Combine coffee grounds and cold water",
			"Steep in refrigerator for 12-24 hours",
			"Strain through cheesecloth or filter",
			"Dilute concentrate with water to taste",
			"Serve over ice with milk and syrup",
		},
		PrepTime:   5,
		CookTime:   0,
		Servings:   4,
		Category:   recipe.CategoryBeverage,
		Difficulty: recipe.DifficultyEasy,
		Author:     recipe.Author{Name: "Coffee Lover Emma", Avatar: ""},
		Ratings:    recipe.Rating{Average: 4.8, Count: 234},
		Likes:      723,
		Saves:      645,
		CreatedAt:  "2025-10-19",
	},
	{
		ID:          "8",
		Title:       "Crispy Chicken Parmesan",
		Description: "Breaded chicken cutlets topped with marinara and melted mozzarella. An Italian-American favorite.",
		Image:       "https://images.unsplash.com/photo-1632778149955-e80f8ceca2e8?w=800&h=600&fit=crop",
		Ingredients: []string{"Chicken breast", "Breadcrumbs", "Parmesan", "Eggs", "Marinara sauce", "Mozzarella", "Basil", "Olive oil"},
		Instructions: []string{
			"Pound chicken breasts thin",
			"Set up breading station: flour, egg, breadcrumbs",
			"Bread chicken cutlets",
			"Fry in olive oil until golden",
			"Top with marinara and mozzarella",
			"Broil until cheese melts",
			"Garnish with fresh basil",
		},
		PrepTime:   20,
		CookTime:   25,
		Servings:   4,
		Category:   recipe.CategoryDinner,
		Difficulty: recipe.DifficultyMedium,
		Author:     recipe.Author{Name: "Nonna Rosa", Avatar: ""},
		Ratings:    recipe.Rating{Average: 4.8, Count: 198},
		Likes:      456,
		Saves:      412,
		CreatedAt:  "2025-10-15",
	},
	{
		ID:          "9",
		Title:       "Banana Bread",
		Description: "Moist, sweet banana bread that's perfect for breakfast or an afternoon snack with coffee.",
		Image:       "https://images.unsplash.com/photo-1626645738196-c2a7c87a8f58?w=800&h=600&fit=crop",
		Ingredients: []string{"Ripe bananas", "Flour", "Sugar", "Eggs", "Butter", "Baking soda", "Salt", "Vanilla", "Walnuts (optional)"},
		Instructions: []string{
			"Preheat oven to 350°F (175°C)",
			"Mash bananas in large bowl",
			"Mix in melted butter",
			"Add sugar, egg, and vanilla",
			"Sprinkle baking soda and salt over mixture",
			"Stir in flour gently",
			"Pour into greased loaf pan",
			"Bake for 60 minutes",
		},
		PrepTime:   15,
		CookTime:   60,
		Servings:   8,
		Category:   recipe.CategorySnacks,
		Difficulty: recipe.DifficultyEasy,
		Author:     recipe.Author{Name: "Grandma Betty", Avatar: ""},
		Ratings:    recipe.Rating{Average: 4.9, Count: 267},
		Likes:      834,
		Saves:      789,
		CreatedAt:  "2025-10-07",
	},
	{
		ID:          "10",
		Title:       "Beef Tacos",
		Description: "Seasoned ground beef in crispy or soft tortillas with fresh toppings. A family favorite.",
		Image:       "https://images.unsplash.com/photo-1551504734-5ee1c4a1479b?w=800&h=600&fit=crop",
		Ingredients: []string{"Ground beef", "Taco seasoning", "Tortillas", "Lettuce", "Tomatoes", "Cheese", "Sour cream", "Salsa"},
		Instructions: []string{
			"Brown ground beef in skillet",
			"Add taco seasoning and water",
			"Simmer until thickened",
			"Warm tortillas",
			"Chop lettuce and tomatoes",
			"Assemble tacos with desired toppings",
		},
		PrepTime:   10,
		CookTime:   15,
		Servings:   6,
		Category:   recipe.CategoryDinner,
		Difficulty: recipe.DifficultyEasy,
		Author:     recipe.Author{Name: "Maria Garcia", Avatar: ""},
		Ratings:    recipe.Rating{Average: 4.8, Count: 223},
		Likes:      621,
		Saves:      534,
		CreatedAt:  "2025-10-13",
	},
	{
		ID:          "11",
		Title:       "Chocolate Chip Cookies",
		Description: "Chewy in the middle, crispy on the edges, and loaded with chocolate chips. The ultimate cookie.",
		Image:       "https://images.unsplash.com/photo-1499636136210-6f4ee915583e?w=800&h=600&fit=crop",
		Ingredients: []string{"Flour", "Butter", "Brown sugar", "White sugar", "Eggs", "Vanilla", "Baking soda", "Chocolate chips", "Salt"},
		Instructions: []string{
			"Cream butter and sugars",
			"Beat in eggs and vanilla",
			"Mix dry ingredients separately",
			"Combine wet and dry ingredients",
			"Fold in chocolate chips",
			"Chill dough for 30 minutes",
			"Bake at 350°F for 10-12 minutes",
		},
		PrepTime:   15,
		CookTime:   12,
		Servings:   24,
		Category:   recipe.CategoryDessert,
		Difficulty: recipe.DifficultyEasy,
		Author:     recipe.Author{Name: "Cookie Monster Chris", Avatar: ""},
		Ratings:    recipe.Rating{Average: 4.9, Count: 312},
		Likes:      921,
		Saves:      856,
		CreatedAt:  "2025-10-21",
	},
	{
		ID:          "12",
		Title:       "Mango Smoothie Bowl",
		Description: "Thick, creamy mango smoothie topped with fresh fruit and granola. A refreshing tropical breakfast.",
		Image:       "https://images.unsplash.com/photo-1590301157890-4810ed352733?w=800&h=600&fit=crop",
		Ingredients: []string{"Frozen mango", "Banana", "Coconut milk", "Honey", "Granola", "Fresh berries", "Coconut flakes", "Chia seeds"},
		Instructions: []string{
			"Blend frozen mango, banana, and coconut milk",
			"Blend until thick and smooth",
			"Pour into bowl",
			"Top with granola, berries, coconut, and chia seeds",
			"Drizzle with honey",
			"Serve immediately",
		},
		PrepTime:   10,
		CookTime:   0,
		Servings:   2,
		Category:   recipe.CategoryBreakfast,
		Difficulty: recipe.DifficultyEasy,
		Author:     recipe.Author{Name: "Healthy Hannah", Avatar: ""},
		Ratings:    recipe.Rating{Average: 4.7, Count: 145},
		Likes:      456,
		Saves:      398,
		CreatedAt:  "2025-10-18",
	},
}

// GetRecipeOfTheDay returns the recipe of the day based on current date
func GetRecipeOfTheDay() recipe.Recipe {
	// Use day of month to select a recipe (1-12 maps to recipe 0-11)
	day := 22 // Simulated day
	index := (day - 1) % len(mockRecipes)
	return mockRecipes[index]
}

// GetTrendingRecipes returns the most popular recipes sorted by saves + likes
func GetTrendingRecipes(limit int) []recipe.Recipe {
	// Create a copy to avoid modifying original
	recipes := make([]recipe.Recipe, len(mockRecipes))
	copy(recipes, mockRecipes)

	// Sort by popularity (saves + likes)
	for i := 0; i < len(recipes)-1; i++ {
		for j := i + 1; j < len(recipes); j++ {
			iPopularity := recipes[i].Saves + recipes[i].Likes
			jPopularity := recipes[j].Saves + recipes[j].Likes
			if jPopularity > iPopularity {
				recipes[i], recipes[j] = recipes[j], recipes[i]
			}
		}
	}

	// Return limited results
	if limit > len(recipes) {
		limit = len(recipes)
	}
	return recipes[:limit]
}

// SearchRecipes filters recipes by keyword and category
func SearchRecipes(keyword, category string) []recipe.Recipe {
	var results []recipe.Recipe

	keyword = strings.ToLower(keyword)

	for _, r := range mockRecipes {
		// Check category filter
		if category != "" && category != "All" {
			categoryStr := CategoryToString(r.Category)
			if !strings.EqualFold(categoryStr, category) {
				continue
			}
		}

		// Check keyword filter
		if keyword != "" {
			titleMatch := strings.Contains(strings.ToLower(r.Title), keyword)
			descMatch := strings.Contains(strings.ToLower(r.Description), keyword)
			authorMatch := strings.Contains(strings.ToLower(r.Author.Name), keyword)

			if !titleMatch && !descMatch && !authorMatch {
				continue
			}
		}

		results = append(results, r)
	}

	return results
}

// GetCategories returns all available categories
func GetCategories() []string {
	return []string{
		"Breakfast",
		"Lunch",
		"Dinner",
		"Dessert",
		"Snacks",
		"Appetizer",
		"Beverage",
	}
}

// CategoryToString converts category enum to string
func CategoryToString(cat recipe.Category) string {
	switch cat {
	case recipe.CategoryBreakfast:
		return "Breakfast"
	case recipe.CategoryLunch:
		return "Lunch"
	case recipe.CategoryDinner:
		return "Dinner"
	case recipe.CategoryDessert:
		return "Dessert"
	case recipe.CategorySnacks:
		return "Snacks"
	case recipe.CategoryAppetizer:
		return "Appetizer"
	case recipe.CategoryBeverage:
		return "Beverage"
	default:
		return "Unknown"
	}
}

// DifficultyToString converts difficulty enum to string
func DifficultyToString(diff recipe.Difficulty) string {
	switch diff {
	case recipe.DifficultyEasy:
		return "Easy"
	case recipe.DifficultyMedium:
		return "Medium"
	case recipe.DifficultyHard:
		return "Hard"
	default:
		return "Unknown"
	}
}
