package server

import (
	"github.com/gin-gonic/gin"
)

func (app *Application) requestLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		app.InfoLogger.Printf("%s - %s %s %s", c.ClientIP(), c.Request.Proto, c.Request.Method, c.Request.RequestURI)
		c.Next()
	}
}
