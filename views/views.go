package views

import (
	"errors"
	"fmt"
	"strings"

	"github.com/CloudyKit/jet/v6"
	"github.com/gin-gonic/gin"
	"github.com/gookit/goutil/dump"
)

var set *jet.Set

func Setup(isProduction bool) {
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
}

func Render(c *gin.Context, name string, data any) error {
	if strings.ContainsRune(name, '#') {
		splitted := strings.Split(name, "#")
		if len(splitted) != 2 {
			return errors.New("only one '#' is allowed in the template name as a template fragment")
		}

		templatePath := splitted[0]
		blockName := splitted[1]

		expr := fmt.Sprintf(`{{ import "%s" }} {{ yield %s() . }}`, templatePath, blockName)
		t, err := set.Parse(templatePath, expr)
		if err != nil {
			return fmt.Errorf("failed to parse expr '%s': %w", expr, err)
		}

		return t.Execute(c.Writer, nil, data)
	}

	template, err := set.GetTemplate(name)
	if err != nil {
		return err
	}

	return template.Execute(c.Writer, nil, data)
}
