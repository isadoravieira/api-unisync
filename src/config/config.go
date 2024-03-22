package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	ConectionDB = ""
	Port        = 0
)

// Load starts environment variables
func Load() {
	var err error

	if err = godotenv.Load(); err != nil {
		log.Fatal("could not load environment variables: %w", err)
	}

	Port, err = strconv.Atoi(os.Getenv("API_PORT"))
	if err != nil {
		Port = 5000
	}

	ConectionDB = fmt.Sprintf(
		"%s:%s@/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)
}
