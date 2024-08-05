package main

import (
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

// 1 - Instalar mux para fazer o Router
func main() {
	fmt.Print("Rodando API")
	r := router.Gerar()

	log.Fatal(http.ListenAndServe(":8000", r))
}
