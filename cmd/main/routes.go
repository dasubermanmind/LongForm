package main

import (
	"github.com/go-chi/chi/v5"
	"net/http"
)

func (app *application) routes() http.Handler {
	mux := chi.NewRouter()
	//mux.Use(SessionLoad)

	mux.Get("/", app.Home)
	mux.Get("/posts", app.Posts)
	mux.Get("/about", app.About)
	mux.Get("/contact", app.Contact)
	mux.Get("/advocacy", app.Advocacy)
	fileServer := http.FileServer(http.Dir("./static"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))
	return mux
}
