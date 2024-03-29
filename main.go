package main

import (
	"fmt"
	"log"
	"os"

	"github.com/daqing/waterway/config"
	"github.com/daqing/waterway/lib/repo"
	"github.com/daqing/waterway/lib/utils"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
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
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	repo.Setup()

	app := NewApp()

	r := app.Router()

	r.Use(cors.Default())

	config.Routes(r)

	port := utils.GetEnvMust("WATERWAY_PORT")

	fmt.Printf("Waterway running at: http://127.0.0.1:%s\n", port)
	app.Run(":" + port)
}

type App struct {
	r *gin.Engine
}

func NewApp() *App {
	router := gin.New()

	router.Static("/public", "./public")

	router.Use(static.Serve("/", static.LocalFile("./public", false)))
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	return &App{
		r: router,
	}
}

func (a *App) Router() *gin.Engine {
	return a.r
}

func (a *App) Run(addr string) {
	a.r.Run(addr)
}
