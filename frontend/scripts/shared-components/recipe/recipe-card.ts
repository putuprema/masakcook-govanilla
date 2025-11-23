import { isSaved, toggleSaveRecipe } from "../../contexts/saved-recipes";
import * as eventTypes from "../../types/custom-events";
import type { Recipe } from "../../types/recipe";

let recipeCards: NodeListOf<HTMLDivElement>;

export const initialize = () => {
	recipeCards = document.querySelectorAll<HTMLDivElement>(
		'[data-component-type="RecipeCard"]',
	);
	recipeCards.forEach((card) => {
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
	});
};
