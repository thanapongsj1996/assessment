package config

import "os"

type AppEnv struct {
	Port        string
	DatabaseUrl string
}

func AppConfig() AppEnv {
	return AppEnv{
		Port:        os.Getenv("PORT"),
		DatabaseUrl: os.Getenv("DATABASE_URL"),
	}
}
