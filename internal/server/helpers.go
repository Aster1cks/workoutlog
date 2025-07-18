package server

import (
	"errors"
	"fmt"
	"github.com/Aster1cks/workoutlog/internal/errdef"
	"net/http"
	"runtime/debug"
)

func (app *Application) serverError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	app.ErrorLogger.Output(2, trace)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func (app *Application) clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

func (app *Application) sqlError(w http.ResponseWriter, err error, id int) {
	if errors.Is(err, errdef.ErrNoRecord) {
		app.InfoLogger.Printf("Record with ID: %d not in database", id)
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
	} else {
		app.serverError(w, err)
	}
}
