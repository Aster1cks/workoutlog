package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (app *Application) Routes() *gin.Engine {
	r := gin.New()

	r.Use(gin.Recovery(), app.requestLogger())

	r.Any("/", func(ctx *gin.Context) {
		ctx.Redirect(http.StatusMovedPermanently, "/workout")
	})

	r.Handle("GET", "/workout", app.home)
	r.Handle("DELETE", "/workout/:id", app.delete)
	r.Handle("PATCH", "/workout/:id", app.edit)
	r.Handle("POST", "/workout", app.add)
	r.Handle("GET", "/workout/:id", app.getByID)
	//mux.HandleFunc(/)
	return r
}
