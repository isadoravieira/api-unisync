package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/isadoravieira/api-unisync/src/cmd/router"
	"github.com/isadoravieira/api-unisync/src/config"
)

func main() {
	config.Load()

	r := router.Generate()

	fmt.Printf("Listening at the door: %d", config.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), r))
}
