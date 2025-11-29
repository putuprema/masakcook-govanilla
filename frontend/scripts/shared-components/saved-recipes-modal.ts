import * as savedRecipesContext from "../contexts/saved-recipes";
import * as eventTypes from "../types/custom-events";
import { RecipeCategory } from "../types/recipe";

let savedRecipesModalEl: HTMLDivElement | null = null;

export const initialize = () => {
	if (savedRecipesModalEl !== null) return;
	savedRecipesModalEl = document.querySelector(
		'[data-component-type="SavedRecipesModal"]',
	);

	if (savedRecipesModalEl === null) {
		throw new Error(
			"savedRecipesModal.initialize(): Cannot attach to modal: Element not found.",
		);
	}

	// Bind to close button
	savedRecipesModalEl
		.querySelectorAll(
			"#SavedRecipesModal__backdrop, #SavedRecipesModal__closeBtn",
		)
		?.forEach((el) => {
			el.addEventListener("click", toggle);
		});

	updateModal();
	document.addEventListener(eventTypes.savedRecipesChanged, updateModal);
};

export const toggle = () => {
	savedRecipesModalEl?.classList.toggle("hidden");
};

const updateModal = () => {
	initialize();
	const savedRecipes = savedRecipesContext.getSavedRecipes();

	const recipeItemTemplate: HTMLTemplateElement =
		savedRecipesModalEl!.querySelector(
			"#SavedRecipesModal__recipeItemTemplate",
		)!;

	const recipeItems = savedRecipes.map((recipe) => {
		const el = recipeItemTemplate.content.cloneNode(true) as HTMLElement;
		el.id = recipe.ID;

		const img = el.querySelector("img")!;
		img.src = recipe.Image;
		img.alt = recipe.Title;

		el.querySelector("#SavedRecipesModal__recipeItem__title")!.textContent =
			recipe.Title;

		el.querySelector("#SavedRecipesModal__recipeItem__subtitle")!.textContent =
			`${RecipeCategory[recipe.Category]} • ${recipe.CookTime + recipe.PrepTime} min`;

		el.querySelector(
			"#SavedRecipesModal__recipeItem__unsaveBtn",
		)!.addEventListener("click", () => {
			savedRecipesContext.toggleSaveRecipe(recipe);
		});

		return el;
	});

	const emptyStateEl = savedRecipesModalEl!.querySelector(
		"#SavedRecipesModal__emptyState",
	)!;
	const recipeListContainerEl: HTMLDivElement =
		savedRecipesModalEl!.querySelector(
			"#SavedRecipesModal__recipeListContainer",
		)!;
	recipeListContainerEl.replaceChildren(...recipeItems);

	if (recipeItems.length > 0) {
		recipeListContainerEl.classList.remove("hidden");
		emptyStateEl.classList.add("hidden");
	} else {
		recipeListContainerEl.classList.add("hidden");
		emptyStateEl.classList.remove("hidden");
	}
};
