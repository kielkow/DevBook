package main

import (
	"api/src/config"
	"api/src/router"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"log"
	"net/http"
)

// Function to generate env SECRET_KEY
func init() {
	key := make([]byte, 64)

	if _, error := rand.Read(key); error != nil {
		log.Fatal(error)
	}

	stringBase64 := base64.StdEncoding.EncodeToString(key)
	fmt.Println(stringBase64)
}

func main() {
	config.SetEnv()

	r := router.Generate()

	fmt.Printf("Start DevBook API - Listenning Port: %d", config.Port)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), r))
}
