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
	fmt.Println(config.ConectionDB)
	fmt.Println("aqui sera onde inicializa a api")

	r := router.Generate()

	log.Fatal(http.ListenAndServe(":8080", r))
}
