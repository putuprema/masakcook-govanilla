// Main entrypoint for the SSR server
package main

import (
	"encoding/json"
	"net/http"

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
			RecipeOfTheDay:  data.GetRecipeOfTheDay(),
			TrendingRecipes: data.GetTrendingRecipes(12),
			Categories:      data.GetCategories(),
		}
		templ.Handler(pages.Index(vm)).ServeHTTP(w, r)
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
