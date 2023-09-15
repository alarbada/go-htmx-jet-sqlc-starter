package internal

import (
	"github.com/alarbada/go-htmx-jet-sqlc-starter/internal/db"
	"github.com/gin-gonic/gin"
)

type Handlers struct {
	// Here all the app state
	conn *db.Connection
}

func (h *Handlers) Dashboard(c *gin.Context) {
	ctx := c.Request.Context()
	todos, err := h.conn.GetAllTodos(ctx)
	c.HTML(200, "/pages/dashboard.tmpl", gin.H{
		"message": "dashboard",
		"todos":   todos,
		"err":     err,
	})
}

func (h *Handlers) Login(c *gin.Context) {
	c.HTML(200, "/pages/login.tmpl", nil)
}
