// Package recipe
package recipe

type (
	Category   int
	Difficulty int
)

const (
	CategoryBreakfast Category = iota
	CategoryLunch
	CategoryDinner
	CategoryDessert
	CategorySnacks
	CategoryAppetizer
	CategoryBeverage
)

const (
	DifficultyEasy Difficulty = iota
	DifficultyMedium
	DifficultyHard
)

type Author struct {
	Name   string
	Avatar string
}

type Rating struct {
	Average float64
	Count   int
}

type Recipe struct {
	ID           string
	Title        string
	Description  string
	Image        string
	Ingredients  []string
	Instructions []string
	Servings     int
	Category     Category
	Difficulty   Difficulty
	PrepTime     int // in minutes
	CookTime     int // in minutes
	Author       Author
	Ratings      Rating
	Likes        int
	Saves        int
	CreatedAt    string // ISO date string
}
