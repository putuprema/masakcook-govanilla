import * as savedRecipesContext from "./contexts/saved-recipes";
import * as navbar from "./shared-components/navbar";

document.addEventListener("DOMContentLoaded", () => {
	console.log("Hello from global.ts");

	savedRecipesContext.initialize();
	navbar.initialize();
});
