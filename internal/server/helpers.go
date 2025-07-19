package server

import (
	"errors"
	"fmt"
	"net/http"
	"runtime/debug"

	"github.com/Aster1cks/workoutlog/internal/errdef"
	"github.com/gin-gonic/gin"
)

func (app *Application) serverError(c *gin.Context, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	app.ErrorLogger.Output(2, trace)
	c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
}

// func (app *Application) clientError(w http.ResponseWriter, status int) {
// 	http.Error(w, http.StatusText(status), status)
// }

func (app *Application) sqlError(c *gin.Context, err error, id int) {
	if errors.Is(err, errdef.ErrNoRecord) {
		app.InfoLogger.Printf("Record with ID: %d not in database", id)
		c.JSON(http.StatusNotFound, gin.H{"error": http.StatusText(http.StatusNotFound)})
	} else {
		app.serverError(c, err)
	}
}
