package config

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"github.com/ilyakaznacheev/cleanenv"
)

var (
	_, filename, _, _ = runtime.Caller(0)
	Root              = filepath.Join(filepath.Dir(filename), "../")
)

type (
	// Config -.
	Config struct {
		App     App
		HTTP    HTTP
		Log     Log
		OAuth   OAuth
		DB      Database
		RMQ     RMQ
		Metrics Metrics
		Swagger Swagger
	}

	// App -.
	App struct {
		Name    string `env:"APP_NAME,required"`
		Version string `env:"APP_VERSION,required"`
	}

	// HTTP -.
	HTTP struct {
		Port           string `env:"HTTP_PORT,required"`
		UsePreforkMode bool   `env:"HTTP_USE_PREFORK_MODE" env-default:"false"`
	}

	// Log -.
	Log struct {
		Level string `env:"LOG_LEVEL,required"`
	}

	// OAuth -.
	OAuth struct {
		ClientID     string `env:"OAUTH_CLIENT_ID,required"`
		ClientSecret string `env:"OAUTH_CLIENT_SECRET,required"`
	}

	// Database -.
	Database struct {
		PoolMax int    `env:"DB_POOL_MAX" env-default:"10"`
		URL     string `env:"DB_URL,required"`
	}

	// RMQ -.
	RMQ struct {
		ServerExchange string `env:"RMQ_RPC_SERVER,required"`
		ClientExchange string `env:"RMQ_RPC_CLIENT,required"`
		URL            string `env:"RMQ_URL,required"`
	}

	// Metrics -.
	Metrics struct {
		Enabled bool `env:"METRICS_ENABLED" env-default:"true"`
	}

	// Swagger -.
	Swagger struct {
		Enabled bool `env:"SWAGGER_ENABLED" env-default:"false"`
	}
)

// MustLoad returns app config.
func MustLoad() *Config {
	configPath := resolveConfigPath()
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		panic("config on path " + configPath + " does not exist")
	}
	var cfg Config
	err := cleanenv.ReadConfig(configPath, &cfg)
	if err != nil {
		panic(err)
	}
	return &cfg
}

func resolveConfigPath() string {
	mode := os.Getenv("MODE")
	if mode == "" {
		mode = "local"
	}
	currDir, err := os.Getwd()
	if err != nil {
		panic(fmt.Sprintf("Failed to build config path. Error: %s", err))
	}
	return filepath.Join(currDir, ".env."+mode)
}
