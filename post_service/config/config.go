package config

import (
	"github.com/spf13/cast"
	"os"
)

type Config struct {
	Environment        string
	PostgresHost       string
	PostgresPort       int
	PostgresDatabase   string
	PostgresUser       string
	PostgresPassword   string
	LogLevel           string
	RPCPort            string
	UserServiceHost    string
	UserServicePort    int
	CommentServiceHost string
	CommentServicePort int
	MongoDatabase      string
	MongoHost          string
	MongoPort          int
}

func Load() Config {
	c := Config{}

	c.Environment = cast.ToString(getOrReturnDefault("ENVIRONMENT", "develop"))

	c.PostgresHost = cast.ToString(getOrReturnDefault("POSTGRES_HOST", "db"))
	c.PostgresPort = cast.ToInt(getOrReturnDefault("POSTGRES_PORT", 5432))
	c.PostgresDatabase = cast.ToString(getOrReturnDefault("POSTGRES_DATABASE", "ekzamen4db"))
	c.PostgresUser = cast.ToString(getOrReturnDefault("POSTGRES_USER", "postgres"))
	c.PostgresPassword = cast.ToString(getOrReturnDefault("POSTGRES_PASSWORD", "123"))

	c.MongoHost = cast.ToString(getOrReturnDefault("MONGODB_HOST", "mongodb"))
	c.MongoPort = cast.ToInt(getOrReturnDefault("MONGODB_PORT", 27017))
	c.MongoDatabase = cast.ToString(getOrReturnDefault("MONGO_DATABASE", "ekzamen4db"))

	c.CommentServiceHost = cast.ToString(getOrReturnDefault("TEMPLATE_COMMENT_SERVICE_HOST", "comment_service"))
	c.CommentServicePort = cast.ToInt(getOrReturnDefault("TEMPLATE_COMMENT_SERVICE_PORT", "4040"))

	c.UserServiceHost = cast.ToString(getOrReturnDefault("POST_SERVICE_HOST", "user_service"))
	c.UserServicePort = cast.ToInt(getOrReturnDefault("POST_SERVICE_PORT", "5050"))

	c.LogLevel = cast.ToString(getOrReturnDefault("LOG_LEVEL", "debug"))
	c.RPCPort = cast.ToString(getOrReturnDefault("RPC_PORT", ":8080"))
	return c
}

func getOrReturnDefault(key string, defaultValue interface{}) interface{} {
	_, exists := os.LookupEnv(key)
	if exists {
		return os.Getenv(key)
	}
	return defaultValue
}
