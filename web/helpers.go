package main

import (
	"fmt"
	"html/template"
	"net/http"
	"runtime/debug"
)

func (app *application) serverError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	app.errorLog.Print(trace)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}
func (app *application) clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}
func (app *application) notFound(w http.ResponseWriter) {
	app.clientError(w, http.StatusNotFound)
}
func (app *application) render(w http.ResponseWriter, status int, page string, data *templateData) {
	files := []string{
		"./ui/templates/base.html",
		fmt.Sprintf("./ui/templates/pages/%s", page),
	}
	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(w, err)
	}
	ts, err = ts.ParseGlob("./ui/templates/partials/*.html")
	if err != nil {
		app.serverError(w, err)
	}
	w.WriteHeader(status)
	err = ts.ExecuteTemplate(w, "base", data)
	if err != nil {
		app.serverError(w, err)
	}
}
