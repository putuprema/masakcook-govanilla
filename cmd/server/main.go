// Main entrypoint for the SSR server
package main

import (
	"net/http"

	"masakcook/internal/config"
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
		templ.Handler(pages.Index()).ServeHTTP(w, r)
	})

	if err := http.ListenAndServe(":"+cfg.Port, r); err != nil {
		panic(err.Error())
	}
}
