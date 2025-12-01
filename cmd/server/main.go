// Main entrypoint for the SSR server
package main

import (
	"encoding/json"
	"net/http"
	"sync"

	"masakcook/internal/config"
	"masakcook/internal/data"
	"masakcook/internal/viewmodel"
	"masakcook/views/pages"

	"github.com/a-h/templ"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	cfg := config.Load()
	cfg.LogConfig()

	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Handle("/static/*", http.StripPrefix("/static/", http.FileServer(http.Dir(cfg.StaticDir))))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		vm := &viewmodel.IndexViewModel{
			Categories: data.GetCategories(),
		}

		slotContents := make(chan data.SlotContents)

		var wg sync.WaitGroup
		wg.Add(2)

		// Partial async streaming for Recipe Of The Day section
		go func() {
			defer wg.Done()
			recipeOfTheDay := data.GetRecipeOfTheDay()
			slotContents <- data.SlotContents{Name: "recipeOfTheDay", Contents: pages.RecipeOfTheDaySection(&recipeOfTheDay)}
		}()

		// Partial async streaming for Trending Recipes section
		go func() {
			defer wg.Done()
			trendingRecipes := data.GetTrendingRecipes(12)
			slotContents <- data.SlotContents{Name: "trendingRecipes", Contents: pages.TrendingItemsContainer(trendingRecipes)}
		}()

		// Close the stream once all completed
		go func() {
			wg.Wait()
			close(slotContents)
		}()

		templ.Handler(pages.Index(vm, slotContents), templ.WithStreaming()).ServeHTTP(w, r)
	})

	r.Get("/api/recipes/search", func(w http.ResponseWriter, r *http.Request) {
		keyword := r.URL.Query().Get("keyword")
		category := r.URL.Query().Get("category")

		results := data.SearchRecipes(keyword, category)

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"data":  results,
			"count": len(results),
		})
	})

	if err := http.ListenAndServe(":"+cfg.Port, r); err != nil {
		panic(err.Error())
	}
}
