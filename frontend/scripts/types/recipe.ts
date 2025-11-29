export enum RecipeCategory {
	Breakfast,
	Lunch,
	Dinner,
	Dessert,
	Snacks,
	Appetizer,
	Beverage,
}

export enum RecipeDifficulty {
	Easy,
	Medium,
	Hard,
}

export interface Recipe {
	ID: string;
	Title: string;
	Description: string;
	Image: string;
	Ingredients: string[];
	Instructions: string[];
	Servings: number;
	Category: RecipeCategory;
	Difficulty: RecipeDifficulty;
	PrepTime: number;
	CookTime: number;
	Author: {
		Name: string;
		Avatar: string;
	};
	Ratings: {
		Average: number;
		Count: number;
	};
	Likes: number;
	Saves: number;
	CreatedAt: string;
}
