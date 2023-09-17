package internal

import (
	"os"

	"github.com/alarbada/go-htmx-jet-sqlc-starter/internal/db"
	"github.com/alarbada/go-htmx-jet-sqlc-starter/views"
	"github.com/gin-gonic/gin"
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

func StartServer() {
	config := NewAppConfig()

	r := gin.Default()

	conn, err := db.Connect(config.SQliteFile)
	if err != nil {
		panic(err)
	}

	t := views.NewTemplates(config.IsProduction)

	handlers := Handlers{conn, t}

	r.Static("/public", "./public")

	r.GET("/login", handlers.Login)
	r.GET("/", handlers.Dashboard)

	r.Run(":" + config.APP_PORT)
}
