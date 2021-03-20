package entry

import (
	"embed"
	"heroku-line-bot/bootstrap"
	"heroku-line-bot/server"
	"os"
)

func Run(f embed.FS) {
	configName := os.Getenv("config")
	if configName == "" {
		configName = "config"
	}

	cfg := bootstrap.LoadConfig(f, configName)
	if err := bootstrap.LoadEnv(cfg); err != nil {
		panic(err)
	}

	server.Init(cfg)

	if err := server.Run(); err != nil {
		panic(err)
	}
}
