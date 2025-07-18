package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Aster1cks/workoutlog/internal/database"
	"github.com/gorilla/mux"
)

func (app *Application) home(w http.ResponseWriter, r *http.Request) {
	values, err := app.Workouts.GetAll()
	if err != nil {
		app.serverError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(values)
}

func (app *Application) delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		app.serverError(w, err)
		return
	}

	values, err := app.Workouts.DeleteEntry(id)
	if err != nil {
		app.sqlError(w, err, id)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(values)
}

func (app *Application) add(w http.ResponseWriter, r *http.Request) {
	entry := database.Workoutentry{}
	err := json.NewDecoder(r.Body).Decode(&entry)
	if err != nil {
		app.serverError(w, err)
		return
	}

	id, err := app.Workouts.AddEntry(entry.WorkoutType, entry.Duration, entry.Notes)
	if err != nil {
		app.serverError(w, err)
		return
	}
	w.Header().Set("Content-Type", "text")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "New entry added\nCan be found here: http://localhost:4000/workout/%d", id)
}

func (app *Application) edit(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		app.serverError(w, err)
		return
	}

	entry := database.Workoutentry{}
	err = json.NewDecoder(r.Body).Decode(&entry)
	if err != nil {
		app.serverError(w, err)
		return
	}

	err = app.Workouts.EditEntry(entry.WorkoutType, entry.Duration, entry.Notes, id)
	if err != nil {
		app.sqlError(w, err, id)
		return
	}

	w.Header().Set("Content-Type", "text")
	w.WriteHeader(http.StatusOK)
	link := fmt.Sprintf("http://localhost:4000/workout/%d", id)
	fmt.Fprintf(w, "Eddited entry \nCan be found here: <a href=%s>View Workout</a>", link)
}

func (app *Application) getByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		app.serverError(w, err)
	}

	values, err := app.Workouts.EntryByID(id)
	if err != nil {
		app.sqlError(w, err, id)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(values)
}
