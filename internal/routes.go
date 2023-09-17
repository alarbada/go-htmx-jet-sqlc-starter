package internal

import (
	"github.com/alarbada/go-htmx-jet-sqlc-starter/internal/db"
	"github.com/alarbada/go-htmx-jet-sqlc-starter/views"
	"github.com/gin-gonic/gin"
)

type Handlers struct {
	conn *db.Connection
	views.Views
}

func (h *Handlers) Dashboard(c *gin.Context) {
	ctx := c.Request.Context()
	todos, err := h.conn.GetAllTodos(ctx)
	h.Render(c, "/pages/dashboard.tmpl", gin.H{
		"message": "dashboard",
		"todos":   todos,
		"err":     err,
	})
}

func (h *Handlers) Login(c *gin.Context) {
	h.Render(c, "/pages/dashboard.tmpl", gin.H{
		"message": "dashboard",
	})
	c.HTML(200, "/pages/login.tmpl", nil)
}
