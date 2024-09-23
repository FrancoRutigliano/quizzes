package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port      string
	Mongo_Uri string
	User_Db   string
	Pass_Db   string
}

func SetUpConfig() (*Config, error) {
	err := godotenv.Load(".env")
	if err != nil {
		return &Config{}, err
	}

	return &Config{
		Port:      os.Getenv("PORT"),
		Mongo_Uri: os.Getenv("MONGO_URI"),
		User_Db:   os.Getenv("USER_DB"),
		Pass_Db:   os.Getenv("PASS_DB"),
	}, nil
}
