export interface Recipe {
	ID: string;
	Title: string;
	Description: string;
	Image: string;
	Ingredients: string[];
	Instructions: string[];
	Servings: number;
	Category: number;
	Difficulty: number;
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
