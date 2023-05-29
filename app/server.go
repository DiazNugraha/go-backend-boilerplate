package app

import (
	"backend/app/controllers"
	"backend/app/helpers"
	"flag"
	"log"

	"github.com/joho/godotenv"
)

func Run() {
	var server = controllers.Server{}
	var appconfig = controllers.AppConfig{}
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	appconfig.AppName = helpers.GetEnv("APP_NAME", "secret_chat")
	appconfig.AppEnv = helpers.GetEnv("APP_ENV", "development")
	appconfig.AppPort = helpers.GetEnv("APP_PORT", "8000")
	appconfig.AppURL = helpers.GetEnv("APP_URL", "http://localhost:8000")

	flag.Parse()
	arg := flag.Arg(0)
	if arg != "" {
		server.InitializeCommand(appconfig)
	} else {
		server.Initialize(appconfig)
		server.Run()
	}
}
