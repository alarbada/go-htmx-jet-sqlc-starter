package internal

import (
	"os"

	"github.com/alarbada/go-htmx-jet-sqlc-starter/internal/db"
	"github.com/alarbada/go-htmx-jet-sqlc-starter/views"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

type DbConfig struct {
	User, Password, Host, Port, Name string
}

type AppConfig struct {
	APP_PORT     string
	IsProduction bool
	DbConfig     DbConfig
}

func NewAppConfig() AppConfig {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	dbConfig := DbConfig{
		User:     getEnvOr("DB_USER", "postgres"),
		Password: getEnvOr("DB_PASSWORD", "postgres"),
		Host:     getEnvOr("DB_HOST", "localhost"),
		Port:     getEnvOr("DB_PORT", "5432"),
		Name:     getEnvOr("DB_NAME", "postgres"),
	}

	config := AppConfig{
		APP_PORT:     getEnvOr("APP_PORT", ":8080"),
		IsProduction: getEnvOr("APP_ENV", "development") == "production",
		DbConfig:     dbConfig,
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

	err := db.Connect(
		config.DbConfig.User,
		config.DbConfig.Password,
		config.DbConfig.Host,
		config.DbConfig.Port,
		config.DbConfig.Name,
	)
	if err != nil {
		panic(err)
	}

	views.Setup(config.IsProduction)

	r.Static("/public", "./public")

	setupHandlers(r)

	r.Run(":" + config.APP_PORT)
}
