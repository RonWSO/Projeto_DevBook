package main

import (
	"fmt"
	"log"
	"net/http"
	"webapp/src/router"
	"webapp/src/utils"
)

func main() {

	utils.CarregarTemplates()
	r := router.Gerar()

	fmt.Println("Rodando Webapp na porta 3000")
	log.Fatal(http.ListenAndServe(":8080", r))
}
