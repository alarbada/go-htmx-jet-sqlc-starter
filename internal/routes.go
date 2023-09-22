package internal

import (
	"github.com/alarbada/go-htmx-jet-sqlc-starter/internal/db"
	"github.com/alarbada/go-htmx-jet-sqlc-starter/views"
	"github.com/gin-gonic/gin"
)

func setupHandlers(r *gin.Engine) {
	r.GET("/login", func(c *gin.Context) {
		ctx := c.Request.Context()
		todos, err := db.Conn.GetAllTodos(ctx)
		views.Render(c, "/pages/dashboard.tmpl", gin.H{
			"message": "dashboard",
			"todos":   todos,
			"err":     err,
		})
	})

	r.GET("/", func(c *gin.Context) {
		views.Render(c, "/pages/dashboard.tmpl", gin.H{
			"message": "dashboard",
		})
		c.HTML(200, "/pages/login.tmpl", nil)
	})
}

