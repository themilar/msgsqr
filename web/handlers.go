package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/themilar/msgsqr/internal/models"
)

// like django views
func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	messages, err := app.messages.Latest()
	if err != nil {
		app.serverError(w, err)
		return
	}
	for _, message := range messages {
		fmt.Fprintf(w, "%+v\n", message)
	}
	// files := []string{
	// 	"./ui/templates/base.html",
	// 	"./ui/templates/pages/home.html",
	// 	"./ui/templates/partials/nav.html",
	// }
	// t, err := template.ParseFiles(files...)

	// if err != nil {
	// 	app.errorLog.Println(err)
	// 	app.serverError(w, err)
	// 	return
	// }
	// err = t.ExecuteTemplate(w, "base", nil)
	// if err != nil {
	// 	app.errorLog.Print(err)
	// 	app.serverError(w, err)
	// }
}
func (app *application) messageDetail(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}
	message, err := app.messages.Get(id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			app.notFound(w)
		} else {
			app.serverError(w, err)
		}
		return
	}
	fmt.Fprintf(w, "%+v...", message)
}
func (app *application) messageCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}
	title := "O snail"
	content := "O snail\nClimb Mount Fuji,\nBut slowly, slowly!\n\n– Kobayashi Issa"
	id, err := app.messages.Insert(title, content)
	if err != nil {
		app.serverError(w, err)
		return
	}
	http.Redirect(w, r, fmt.Sprintf("/message/detail?id=%d", id), http.StatusSeeOther)
}
