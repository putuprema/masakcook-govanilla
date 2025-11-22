// Type definitions matching the Go structs
interface Recipe {
	ID: string;
	Title: string;
	Description: string;
	Image: string;
	Ingredients: string[];
	Instructions: string[];
	Servings: number;
	Category: number;
	Difficulty: number;
	PrepTime: number;
	CookTime: number;
	Author: {
		Name: string;
		Avatar: string;
	};
	Ratings: {
		Average: number;
		Count: number;
	};
	Likes: number;
	Saves: number;
	CreatedAt: string;
}

interface SearchResponse {
	data: Recipe[];
	count: number;
}

// Enum mappings
const CategoryMap: Record<number, string> = {
	0: "Breakfast",
	1: "Lunch",
	2: "Dinner",
	3: "Dessert",
	4: "Snacks",
	5: "Appetizer",
	6: "Beverage",
};

const DifficultyMap: Record<number, string> = {
	0: "Easy",
	1: "Medium",
	2: "Hard",
};

// State
let currentKeyword = "";
let currentCategory = "All";
let debounceTimer: number | null = null;
let hasSearched = false;

// DOM Elements
const searchInput = document.getElementById("search-input") as HTMLInputElement;
const categoryFilters = document.querySelectorAll<HTMLButtonElement>(".category-filter");
const resultsGrid = document.getElementById("results-grid");
const resultsCount = document.getElementById("results-count");
const initialState = document.getElementById("initial-state");
const loadingState = document.getElementById("loading-state");
const emptyState = document.getElementById("empty-state");

// Initialize
function init() {
	if (!searchInput || !resultsGrid || !loadingState || !emptyState || !initialState || !resultsCount) {
		console.error("Required DOM elements not found");
		return;
	}

	// Search input handler
	searchInput.addEventListener("input", handleSearchInput);

	// Category filter handlers
	categoryFilters.forEach((button) => {
		button.addEventListener("click", handleCategoryClick);
	});

	// Don't perform search on initial load - show initial state instead
}

// Handle search input with debouncing
function handleSearchInput(e: Event) {
	const input = e.target as HTMLInputElement;
	currentKeyword = input.value;

	// Clear existing timer
	if (debounceTimer !== null) {
		clearTimeout(debounceTimer);
	}

	// Set new timer
	debounceTimer = window.setTimeout(() => {
		performSearch();
	}, 300);
}

// Handle category filter click
function handleCategoryClick(e: Event) {
	const button = e.currentTarget as HTMLButtonElement;
	const category = button.dataset.category || "All";

	// Update active state
	categoryFilters.forEach((btn) => {
		if (btn === button) {
			btn.classList.remove("bg-gray-100", "text-gray-700", "hover:bg-gray-200");
			btn.classList.add("bg-blue-600", "text-white", "hover:bg-blue-700");
		} else {
			btn.classList.remove("bg-blue-600", "text-white", "hover:bg-blue-700");
			btn.classList.add("bg-gray-100", "text-gray-700", "hover:bg-gray-200");
		}
	});

	currentCategory = category;
	performSearch();
}

// Perform search API call
async function performSearch() {
	if (!resultsGrid || !loadingState || !emptyState || !initialState || !resultsCount) return;

	hasSearched = true;

	// Show loading state
	showLoading();

	try {
		// Build query params
		const params = new URLSearchParams();
		if (currentKeyword) params.append("keyword", currentKeyword);
		if (currentCategory && currentCategory !== "All") {
			params.append("category", currentCategory);
		}

		// Fetch results
		const response = await fetch(`/api/recipes/search?${params.toString()}`);
		if (!response.ok) {
			throw new Error("Search request failed");
		}

		const data: SearchResponse = await response.json();

		// Update results count
		updateResultsCount(data.count);

		// Update UI with results
		renderResults(data.data);
	} catch (error) {
		console.error("Search error:", error);
		showEmpty();
	}
}

// Update results count display
function updateResultsCount(count: number) {
	if (!resultsCount) return;

	resultsCount.textContent = `Found ${count} recipe${count !== 1 ? "s" : ""}`;
	resultsCount.classList.remove("hidden");
}

// Show loading state
function showLoading() {
	if (!resultsGrid || !loadingState || !emptyState || !initialState || !resultsCount) return;

	initialState.classList.add("hidden");
	resultsGrid.classList.add("hidden");
	emptyState.classList.add("hidden");
	resultsCount.classList.add("hidden");
	loadingState.classList.remove("hidden");
}

// Show empty state
function showEmpty() {
	if (!resultsGrid || !loadingState || !emptyState || !initialState || !resultsCount) return;

	initialState.classList.add("hidden");
	resultsGrid.classList.add("hidden");
	loadingState.classList.add("hidden");
	emptyState.classList.remove("hidden");
}

// Show results
function showResults() {
	if (!resultsGrid || !loadingState || !emptyState || !initialState) return;

	initialState.classList.add("hidden");
	loadingState.classList.add("hidden");
	emptyState.classList.add("hidden");
	resultsGrid.classList.remove("hidden");
}

// Render search results
function renderResults(recipes: Recipe[]) {
	if (!resultsGrid) return;

	// Clear existing results
	resultsGrid.innerHTML = "";

	// Show empty state if no results
	if (recipes.length === 0) {
		showEmpty();
		return;
	}

	// Show results
	showResults();

	// Render each recipe card
	recipes.forEach((recipe) => {
		const card = createRecipeCard(recipe);
		resultsGrid.appendChild(card);
	});
}

// Create recipe card HTML element (matching Next.js design)
function createRecipeCard(recipe: Recipe): HTMLElement {
	const totalTime = recipe.PrepTime + recipe.CookTime;
	const category = CategoryMap[recipe.Category] || "Unknown";
	const difficulty = DifficultyMap[recipe.Difficulty] || "Unknown";

	const card = document.createElement("div");
	card.className =
		"group relative flex flex-col overflow-hidden rounded-lg border border-gray-200 bg-white shadow-sm transition-shadow hover:shadow-lg";

	card.innerHTML = `
		<div class="relative h-48 w-full overflow-hidden bg-gray-200">
			<img
				src="${recipe.Image}"
				alt="${recipe.Title}"
				class="w-full h-full object-cover group-hover:scale-110 transition-transform duration-300"
				loading="lazy"
			/>
			<button
				class="absolute top-3 right-3 p-2 bg-white/90 hover:bg-white rounded-full transition-colors duration-200 save-button"
				data-recipe-id="${recipe.ID}"
				aria-label="Save recipe"
			>
				<svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="w-5 h-5 text-gray-600">
					<path stroke-linecap="round" stroke-linejoin="round" d="M21 8.25c0-2.485-2.099-4.5-4.688-4.5-1.935 0-3.597 1.126-4.312 2.733-.715-1.607-2.377-2.733-4.313-2.733C5.1 3.75 3 5.765 3 8.25c0 7.22 9 12 9 12s9-4.78 9-12Z"></path>
				</svg>
			</button>
		</div>
		<div class="flex flex-1 flex-col p-4">
			<div class="mb-2">
				<span class="inline-block rounded-full bg-blue-100 px-3 py-1 text-xs font-medium text-blue-800">
					${category}
				</span>
			</div>
			<h3 class="mb-2 line-clamp-2 text-lg font-semibold text-gray-900">
				${recipe.Title}
			</h3>
			<p class="mb-4 line-clamp-2 flex-1 text-sm text-gray-600">
				${recipe.Description}
			</p>
			<div class="mb-4 flex items-center gap-4 text-sm text-gray-500">
				<div class="flex items-center gap-1">
					<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="currentColor" class="w-4 h-4">
						<path fill-rule="evenodd" d="M12 2.25c-5.385 0-9.75 4.365-9.75 9.75s4.365 9.75 9.75 9.75 9.75-4.365 9.75-9.75S17.385 2.25 12 2.25ZM12.75 6a.75.75 0 0 0-1.5 0v6c0 .414.336.75.75.75h4.5a.75.75 0 0 0 0-1.5h-3.75V6Z" clip-rule="evenodd"></path>
					</svg>
					<span>${totalTime} min</span>
				</div>
				<div class="rounded bg-gray-100 px-2 py-1 text-xs font-medium text-gray-700">
					${difficulty}
				</div>
			</div>
			<a
				href="/recipe/${recipe.ID}"
				class="flex w-full items-center justify-center gap-2 rounded-md bg-gray-900 px-4 py-2 text-sm font-medium text-white transition-colors hover:bg-gray-800"
			>
				<svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="w-5 h-5">
					<path stroke-linecap="round" stroke-linejoin="round" d="M2.036 12.322a1.012 1.012 0 0 1 0-.639C3.423 7.51 7.36 4.5 12 4.5c4.638 0 8.573 3.007 9.963 7.178.07.207.07.431 0 .639C20.577 16.49 16.64 19.5 12 19.5c-4.638 0-8.573-3.007-9.963-7.178Z"></path>
					<path stroke-linecap="round" stroke-linejoin="round" d="M15 12a3 3 0 1 1-6 0 3 3 0 0 1 6 0Z"></path>
				</svg>
				<span>View Recipe</span>
			</a>
		</div>
	`;

	return card;
}

// Initialize when DOM is ready
if (document.readyState === "loading") {
	document.addEventListener("DOMContentLoaded", init);
} else {
	init();
}
