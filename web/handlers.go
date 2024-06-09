package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/themilar/msgsqr/internal/models"
)

// like django views

type templateData struct {
	// acts like the context object used in views for dynamic template data
	Message  *models.Message
	Messages []*models.Message
}

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

	app.render(w, http.StatusOK, "home.html", &templateData{
		Messages: messages,
	})

}
func (app *application) messageDetail(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())
	id, err := strconv.Atoi(params.ByName("id"))
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
	app.render(w, http.StatusOK, "detail.html", &templateData{
		Message: message,
	})
}
func (app *application) messageCreate(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPost {
		err := r.ParseForm()
		if err != nil {
			app.clientError(w, http.StatusBadRequest)
			return
		}
		title := r.PostForm.Get("title")
		content := r.PostForm.Get("content")
		id, err := app.messages.Insert(title, content)
		if err != nil {
			app.serverError(w, err)
			return
		}
		http.Redirect(w, r, fmt.Sprintf("/message/detail/%d", id), http.StatusSeeOther)
	}

	app.render(w, http.StatusOK, "create.html", &templateData{})

}
