import * as savedRecipesContext from "./contexts/saved-recipes";
import * as navbar from "./shared-components/navbar";
import * as savedRecipesModal from "./shared-components/recipe/saved-recipes-modal";

document.addEventListener("DOMContentLoaded", () => {
	console.log("Hello from global.ts");

	savedRecipesContext.initialize();
	navbar.initialize();
	savedRecipesModal.initialize();
});
