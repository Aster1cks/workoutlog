package server

import (
	"net/http"
	"strconv"

	"github.com/Aster1cks/workoutlog/internal/database"
	"github.com/gin-gonic/gin"
)

func (app *Application) home(c *gin.Context) {
	values, err := app.Workouts.GetAll()
	if err != nil {
		app.serverError(c, err)
		return
	}

	c.JSON(http.StatusOK, values)
}

func (app *Application) delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		app.serverError(c, err)
		return
	}

	values, err := app.Workouts.DeleteEntry(id)
	if err != nil {
		app.sqlError(c, err, id)
		return
	}

	c.JSON(http.StatusOK, values)
}

func (app *Application) add(c *gin.Context) {
	entry := database.Workoutentry{}
	err := c.BindJSON(&entry)
	if err != nil {
		app.serverError(c, err)
		return
	}

	id, err := app.Workouts.AddEntry(entry.WorkoutType, entry.Duration, entry.Notes)
	if err != nil {
		app.serverError(c, err)
		return
	}

	c.String(http.StatusOK, "New entry added\nCan be found here: http://localhost:4000/workout/%d", id)
}

func (app *Application) edit(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		app.serverError(c, err)
		return
	}
	entry := database.Workoutentry{}

	err = c.BindJSON(&entry)
	if err != nil {
		app.serverError(c, err)
		return
	}

	err = app.Workouts.EditEntry(entry.WorkoutType, entry.Duration, entry.Notes, id)
	if err != nil {
		app.sqlError(c, err, id)
		return
	}

	c.String(http.StatusOK, "Entry eddited\nCan be found here: http://localhost:4000/workout/%d", id)

}

func (app *Application) getByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		app.serverError(c, err)
	}

	values, err := app.Workouts.EntryByID(id)
	if err != nil {
		app.sqlError(c, err, id)
		return
	}

	c.JSON(http.StatusOK, values)
}
