// Package recipe
package recipe

type (
	Category   int
	Difficulty int
)

const (
	Breakfast Category = iota
	Lunch
	Dinner
	Dessert
	Snacks
	Appetizer
	Beverage
)

const (
	Easy Difficulty = iota
	Medium
	Hard
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
