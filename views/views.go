package views

import (
	"github.com/CloudyKit/jet/v6"
	"github.com/gin-gonic/gin"
	"github.com/gookit/goutil/dump"
)

type Views struct {
	set *jet.Set
}

func New(isProduction bool) Views {
	var set *jet.Set
	if isProduction {
		set = jet.NewSet(jet.NewOSFileSystemLoader("./views"))
	} else {
		set = jet.NewSet(jet.NewOSFileSystemLoader("./views"), jet.InDevelopmentMode())
	}

	dump.Config(dump.OptionFunc(func(opts *dump.Options) {
		opts.MaxDepth = 4
		opts.IndentLen = 3
		opts.BytesAsString = true
	}))

	set.AddGlobal("dump", func(i any) string {
		dump.P(i)
		return ""
	})

	return Views{set}
}

func (t *Views) Render(c *gin.Context, name string, data any) error {
	template, err := t.set.GetTemplate(name)
	if err != nil {
		return err
	}

	return template.Execute(c.Writer, nil, data)
}
