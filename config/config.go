package config

import (
	"time"

	"github.com/eridiumdev/yandex-praktikum-go-diploma/pkg/http"
)

type (
	Config struct {
		App    App
		Logger Logger
		Server http.ServerConfig
	}
	App struct {
		ShutdownTimeout time.Duration
	}
	Logger struct {
		Level  string
		Pretty bool
	}
)

func Load() (*Config, error) {
	cfg := &Config{
		App: App{
			ShutdownTimeout: time.Second * 3,
		},
		Logger: Logger{
			Level:  "debug",
			Pretty: true,
		},
		Server: http.ServerConfig{
			Addr: ":8080",
		},
	}

	return cfg, nil
}
