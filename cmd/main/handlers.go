package main

import "net/http"

// Home
func (app *application) Home(w http.ResponseWriter, r *http.Request) {
	if err := app.renderTemplate(w, r, "home", &templateData{}); err != nil {
		app.errorLog.Println(err)
	}
}

// Posts
func (app *application) Posts(w http.ResponseWriter, r *http.Request) {
	if err := app.renderTemplate(w, r, "posts", &templateData{}); err != nil {
		app.errorLog.Println(err)
	}
}

// About
func (app *application) About(w http.ResponseWriter, r *http.Request) {
	if err := app.renderTemplate(w, r, "about", &templateData{}); err != nil {
		app.errorLog.Println(err)
	}
}

// Contact
func (app *application) Contact(w http.ResponseWriter, r *http.Request) {
	if err := app.renderTemplate(w, r, "contact", &templateData{}); err != nil {
		app.errorLog.Println(err)
	}
}

// Advocacy
func (app *application) Advocacy(w http.ResponseWriter, r *http.Request) {
	if err := app.renderTemplate(w, r, "advocacy", &templateData{}); err != nil {
		app.errorLog.Println(err)
	}
}
