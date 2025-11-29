import * as recipeCard from "../shared-components/recipe-card";

document.addEventListener("DOMContentLoaded", () => {
	// Bind to all recipe cards in the Trending Now section
	const recipeCards = document.querySelectorAll<HTMLDivElement>(
		'[data-component-type="RecipeCard"]',
	);
	recipeCards.forEach((card) => {
		recipeCard.initialize(card);
	});
});
