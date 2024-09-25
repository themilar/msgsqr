package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// like urls
func (app *application) routes() http.Handler {
	router := httprouter.New()
	router.NotFound = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		app.notFound(w)
	})
	fileServer := http.FileServer(http.Dir("./ui/static"))
	router.Handler(http.MethodGet, "/static/*filepath", http.StripPrefix("/static", fileServer))
	router.HandlerFunc(http.MethodGet, "/", app.home)
	router.HandlerFunc(http.MethodGet, "/message/detail/:id", app.messageDetail)
	router.HandlerFunc(http.MethodPost, "/message/delete/:id", app.messageDelete)
	router.HandlerFunc(http.MethodGet, "/message/delete/:id", app.messageDelete)
	router.HandlerFunc(http.MethodGet, "/message/create", app.messageCreate)
	router.HandlerFunc(http.MethodPost, "/message/create", app.messageCreate)
	// router.HandlerFunc(http.MethodPost, "/snippet/create", 	app.messageCreatePost)

	return app.requestLogger(secureHeaders(router))
}
