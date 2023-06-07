package main

import (
	"template-go/cmd"
	"template-go/config"
	"template-go/internal/app"
)

func main() {
	//Get mode from terminal
	mode := cmd.Init()

	//After get mode, Get config from config/xx.yaml
	cfg := config.NewConfig(mode.AppMode)

	//Run app
	app.Run(cfg)
}
