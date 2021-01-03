package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	// APIURL is the API URL
	APIURL = ""

	// Port APP
	Port = 0

	// HashKey to authenticate cookie
	HashKey []byte

	// BlockKey to encrypt cookie data
	BlockKey []byte
)

// SetEnv func set env variables
func SetEnv() {
	var error error

	if error = godotenv.Load(); error != nil {
		log.Fatal(error)
	}

	Port, error = strconv.Atoi(os.Getenv("APP_PORT"))
	if error != nil {
		log.Fatal(error)
	}

	APIURL = os.Getenv("API_URL")
	HashKey = []byte(os.Getenv("HASH_KEY"))
	BlockKey = []byte(os.Getenv("BLOCK_KEY"))
}
