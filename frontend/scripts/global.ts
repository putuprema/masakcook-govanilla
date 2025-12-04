import * as savedRecipesContext from "./contexts/saved-recipes";
import * as navbar from "./shared-components/navbar";
import * as savedRecipesModal from "./shared-components/saved-recipes-modal";

console.log("Hello from global.ts");

savedRecipesContext.initialize();
navbar.initialize();
savedRecipesModal.initialize();
