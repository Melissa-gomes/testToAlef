package config

import "os"

type Env struct {
	APPLICATION_PORT  string
	POSTGRES_DATABASE string
	POSTGRES_USERNAME string
	POSTGRES_PASSWORD string
	POSTGRES_HOST     string
	POSTGRES_PORT     string
}

func LoadEnv() *Env {
	return &Env{
		APPLICATION_PORT:  os.Getenv("APPLICATION_PORT"),
		POSTGRES_DATABASE: os.Getenv("POSTGRES_DATABASE"),
		POSTGRES_USERNAME: os.Getenv("POSTGRES_USERNAME"),
		POSTGRES_PASSWORD: os.Getenv("POSTGRES_PASSWORD"),
		POSTGRES_HOST:     os.Getenv("POSTGRES_HOST"),
		POSTGRES_PORT:     os.Getenv("POSTGRES_PORT"),
	}
}
