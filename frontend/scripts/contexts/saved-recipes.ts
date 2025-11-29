import * as eventTypes from "../types/custom-events";
import type { Recipe } from "../types/recipe";

interface SavedRecipesContextType {
	savedRecipes: Recipe[];
}

const STORAGE_KEY = "masakcook_saved_recipes";
let isInitialized = false;

const _savedRecipesContext: SavedRecipesContextType = {
	savedRecipes: [],
};
const savedRecipesContext = new Proxy(_savedRecipesContext, {
	get(target, key) {
		return Reflect.get(target, key);
	},
	set(target, key, value) {
		const result = Reflect.set(target, key, value);
		notify();
		return result;
	},
});

const notify = () => {
	document.dispatchEvent(new CustomEvent(eventTypes.savedRecipesChanged));
};

export const getSavedRecipes = () => savedRecipesContext.savedRecipes;

export const initialize = () => {
	if (isInitialized) return;
	const stored = localStorage.getItem(STORAGE_KEY);
	if (stored) {
		const savedRecipeParsed = JSON.parse(stored) as Recipe[];
		savedRecipesContext.savedRecipes = savedRecipeParsed;
	}
	isInitialized = true;
};

export const isSaved = (recipe: Recipe) =>
	savedRecipesContext.savedRecipes.findIndex((el) => el.ID === recipe.ID) !==
	-1;

export const getSavedCount = () => savedRecipesContext.savedRecipes.length;

export const toggleSaveRecipe = (recipe: Recipe) => {
	const savedRecipes = savedRecipesContext.savedRecipes;
	let neww: Recipe[] = [];
	if (
		savedRecipes.findIndex((prevRecipe) => prevRecipe.ID === recipe.ID) !== -1
	) {
		neww = savedRecipes.filter((prevRecipe) => prevRecipe.ID !== recipe.ID);
	} else {
		neww = [...savedRecipes, recipe];
	}
	localStorage.setItem(STORAGE_KEY, JSON.stringify(neww));
	savedRecipesContext.savedRecipes = neww;
};
