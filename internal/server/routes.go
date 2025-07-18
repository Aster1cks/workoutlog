package server

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (app *Application) Routes() http.Handler {
	mux := mux.NewRouter()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/workout", http.StatusMovedPermanently)
	})
	mux.HandleFunc("/workout", app.home).Methods("GET")
	mux.HandleFunc("/workout/{id}", app.delete).Methods("DELETE")
	mux.HandleFunc("/workout/{id}", app.edit).Methods("PATCH")
	mux.HandleFunc("/workout", app.add).Methods("POST")
	mux.HandleFunc("/workout/{id}", app.getByID).Methods("GET")
	//mux.HandleFunc(/)
	return app.requestLogger(mux)
}
