package config

import (
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
)

type (
	Config struct {
		Server *Server `validate:"required"`
		Bitkub *Bitkub `validate:"required"`
	}

	Server struct {
		Port         int           `validate:"required"`
		AllowOrigins []string      `validate:"required"`
		BodyLimit    int           `validate:"required"`
		TimeOut      time.Duration `validate:"required"`
	}

	Bitkub struct {
		BitkubApiKey    string `validate:"required"`
		BitkubApiSecret string `validate:"required"`
	}
)

var (
	once           sync.Once
	configInstance *Config
)

func ConfigGetting() *Config {
	once.Do(func() {
		_ = godotenv.Load()

		configInstance = &Config{
			Server: &Server{},
		}

		// Server
		port, err := strconv.Atoi(getEnv("SERVER_PORT", "8080"))
		if err != nil {
			panic("SERVER_PORT is not set")
		}
		bodyLimit, err := strconv.Atoi(getEnv("SERVER_BODY_LIMIT", "10485760")) // 10MB
		if err != nil {
			panic("SERVER_BODY_LIMIT in not set")
		}
		timeOut, err := time.ParseDuration(getEnv("SERVER_TIMEOUT", "30s"))
		if err != nil {
			panic("SERVER_TIMEOUT is not set")
		}
		configInstance.Server = &Server{
			Port:         port,
			AllowOrigins: strings.Split(getEnv("SERVER_ALLOW_ORIGINS", "*"), ","),
			BodyLimit:    bodyLimit,
			TimeOut:      timeOut,
		}

		configInstance.Bitkub = &Bitkub{
			BitkubApiKey:    getEnv("BITKUB_API_KEY", ""),
			BitkubApiSecret: getEnv("BITKUB_API_SECRET", ""),
		}
	})

	validating := validator.New()
	if err := validating.Struct(configInstance); err != nil {
		panic(err)
	}

	return configInstance
}

func getEnv(key, defaultValue string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}
	return defaultValue
}
