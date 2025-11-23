import * as savedRecipesContext from "../contexts/saved-recipes";
import * as eventTypes from "../types/custom-events";

let navbarEl: HTMLElement | null = null;

export const initialize = () => {
	if (navbarEl !== null) return;
	navbarEl = document.querySelector("nav");

	if (navbarEl === null) {
		throw new Error(
			"navbar.initialize(): Cannot attach to navbar: Element not found.",
		);
	}

	updateNavbar();
	document.addEventListener(eventTypes.savedRecipesChanged, updateNavbar);
};

const updateNavbar = () => {
	initialize();
	const bookmarkBadge: HTMLSpanElement = navbarEl!.querySelector(
		"#nav__bookmarkBadge",
	)!;
	const savedCount = savedRecipesContext.getSavedCount();

	if (savedCount > 0) {
		bookmarkBadge.classList.remove("hidden");
		bookmarkBadge.textContent = savedCount > 99 ? "99+" : savedCount.toString();
	} else {
		bookmarkBadge.classList.add("hidden");
		bookmarkBadge.textContent = "0";
	}
};
