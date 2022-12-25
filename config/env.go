package config

import "os"

type A struct {
	Port string
}

func AppConfig() A {
	return  A{
		Port: os.Getenv("PORT"),
	}
}