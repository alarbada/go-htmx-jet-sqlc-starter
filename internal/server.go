package internal

import (
	"os"

	"github.com/CloudyKit/jet/v6"
	"github.com/alarbada/go-htmx-jet-sqlc-starter/internal/db"
	jetr "github.com/alarbada/jet-html-renderer"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/render"
	"github.com/gookit/goutil/dump"
	"github.com/joho/godotenv"
)

type AppConfig struct {
	APP_PORT     string
	IsProduction bool
	SQliteFile   string
}

func NewAppConfig() AppConfig {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	config := AppConfig{
		APP_PORT:     getEnvOr("APP_PORT", ":8080"),
		IsProduction: getEnvOr("APP_ENV", "development") == "production",
		SQliteFile:   getRequiredEnv("SQLITE_FILE"),
	}

	return config
}

func getRequiredEnv(envName string) string {
	env := os.Getenv(envName)
	if env == "" {
		panic("Required env " + envName + " is not set")
	}
	return env
}

func getEnvOr(envName string, defaultValue string) string {
	env := os.Getenv(envName)
	if env == "" {
		return defaultValue
	}
	return env
}

type jetTemplate interface {
	AddGlobal(string, any) *jet.Set
}

func setupGlobals(renderer jetTemplate) {
	dump.Config(dump.OptionFunc(func(opts *dump.Options) {
		opts.MaxDepth = 4
		opts.IndentLen = 3
		opts.BytesAsString = true
	}))

	renderer.AddGlobal("dump", func(i any) string {
		dump.P(i)
		return ""
	})
}

func setupTemplates(config AppConfig) render.HTMLRender {
	if config.IsProduction {
		gin.SetMode(gin.ReleaseMode)

		renderer := jetr.New(
			jet.NewOSFileSystemLoader("./templates"),
		)

		setupGlobals(renderer.Set)
		return renderer
	}

	renderer := jetr.New(
		jet.NewOSFileSystemLoader("./templates"),
		jet.InDevelopmentMode(),
	)

	setupGlobals(renderer.Set)
	return renderer
}

func StartServer() {
	config := NewAppConfig()

	r := gin.Default()

	r.HTMLRender = setupTemplates(config)

	conn, err := db.Connect(config.SQliteFile)
	if err != nil {
		panic(err)
	}

	handlers := Handlers{conn}

	r.Static("/public", "./public")

	r.GET("/login", handlers.Login)
	r.GET("/", handlers.Dashboard)

	r.Run(":" + config.APP_PORT)
}
