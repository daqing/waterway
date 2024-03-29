package main

import (
	"fmt"
	"log"
	"os"

	"github.com/daqing/waterway/cli/generator"
	"github.com/daqing/waterway/cli/scaffold"
	"github.com/daqing/waterway/cli/seed"
	"github.com/daqing/waterway/lib/repo"
	"github.com/daqing/waterway/lib/utils"
	"github.com/joho/godotenv"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		showHelp()
		return
	}

	setUpDB()

	switch args[0] {
	case "g":
		generator.Generate(args[1:])
	case "sf":
		scaffold.Generate(args[1:])
	case "seed":
		seed.Generate(args[1:])
	default:
		showHelp()
	}

}

func showHelp() {
	fmt.Println("cli g [what] [params]")
	fmt.Println("cli seed [what] [params]")
	fmt.Println("cli sf [model] [attr:type] [attr:type]")
}

func setUpDB() {
	appConfig := utils.AppConfig()

	if appConfig.Env == "" {
		log.Println("WATERWAY_ENV is not set")
		os.Exit(1)
	}

	if appConfig.IsLocal {
		envFile := ".env"
		err := godotenv.Load(envFile)
		if err != nil {
			log.Printf("Loading env file: %s failed", envFile)
			os.Exit(2)
		}
	}

	repo.Setup()
}
