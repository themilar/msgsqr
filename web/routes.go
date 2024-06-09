package main

import "net/http"

// like urls
func (app *application) routes() http.Handler {
	mux := http.NewServeMux()
	fileServer := http.FileServer(http.Dir("./ui/static"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))
	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/message/detail/", app.messageDetail)
	mux.HandleFunc("/message/create/", app.messageCreate)
	return app.requestLogger(secureHeaders(mux))
}
