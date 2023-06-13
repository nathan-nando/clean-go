package main

import (
	"template-go/app"
	"template-go/cmd"
	"template-go/config"
)

func main() {
	//Get mode from terminal
	mode := cmd.Init()

	//After get mode, Get config from config/xx.yaml
	cfg := config.NewConfig(mode.AppMode)

	//Send dbMode to app
	dbMode := mode.DbMode

	//Setup app
	api := app.New(cfg, dbMode)

	//Run server
	app.Run(api)
}
