import { isSaved, toggleSaveRecipe } from "../contexts/saved-recipes";
import * as eventTypes from "../types/custom-events";
import type { Recipe } from "../types/recipe";

export const initialize = (card: HTMLDivElement) => {
	if (!card.matches('[data-component-type="RecipeCard"]')) {
		return;
	}

	const recipeData = JSON.parse(card.dataset.recipe!) as Recipe;

	const onSavedRecipeChanged = () => {
		const heartDisabledIcon = card.querySelector(
			"#RecipeCard__heartDisabledIcon",
		);
		const heartEnabledIcon = card.querySelector(
			"#RecipeCard__heartEnabledIcon",
		);
		if (isSaved(recipeData)) {
			heartDisabledIcon?.classList.add("hidden");
			heartEnabledIcon?.classList.remove("hidden");
		} else {
			heartDisabledIcon?.classList.remove("hidden");
			heartEnabledIcon?.classList.add("hidden");
		}
	};
	onSavedRecipeChanged();

	// Configure bookmark button logic
	const bookmarkBtn: HTMLButtonElement = card.querySelector(
		'[data-button-instance="RecipeCard__saveRecipeBtn"]',
	)!;
	bookmarkBtn.addEventListener("click", () => {
		toggleSaveRecipe(recipeData);
	});

	// Listen to saved recipe changes
	document.addEventListener(
		eventTypes.savedRecipesChanged,
		onSavedRecipeChanged,
	);
};
