package utils

import "os"

var env = TrimFull(os.Getenv("WATERWAY_ENV"))

const PRODUCTION_ENV = "production"

type appConfig struct {
	IsLocal bool
	Env     string
}

var defaultConfig = &appConfig{
	IsLocal: env != PRODUCTION_ENV,
	Env:     env,
}

func AppConfig() *appConfig {
	return defaultConfig
}
