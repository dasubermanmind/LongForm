package main

import (
	"io/ioutil"
	"net/http"
)

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

type Bitcoin struct {
	ID     string
	Rank   string
	Symbol string
	Name   string
	supply string
}

// Advocacy
func (app *application) Advocacy(w http.ResponseWriter, r *http.Request) {
	//var btc Bitcoin
	response, err := http.Get("") // TODO: Need to find a reliable API
	if err != nil {
		app.errorLog.Println("", err)
		return
	}
	app.infoLog.Println(response)

	//retVal = Bitcoin {
	//	ID: response
	//}

	defer response.Body.Close()

	bcBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return
	}

	data := make(map[string]interface{})

	if err != nil {
		return
	}
	data["btc"] = bcBody

	if err := app.renderTemplate(w, r, "advocacy", &templateData{
		Data: data,
	}); err != nil {
		app.errorLog.Println(err)
	}
}
