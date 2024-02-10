package config

import "os"

type database struct {
	Name     string
	User     string
	Password string
	Host     string
	Port     string
}

func Database() database {
	return database{
		Name:     os.Getenv("DATABASE_NAME"),
		User:     os.Getenv("DATABASE_USER"),
		Password: os.Getenv("DATABASE_PASSWORD"),
		Host:     os.Getenv("DATABASE_HOST"),
		Port:     os.Getenv("DATABASE_PORT"),
	}
}
