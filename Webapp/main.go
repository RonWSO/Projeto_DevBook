package main

import (
	"encoding/hex"
	"fmt"
	"log"
	"net/http"
	"webapp/src/config"
	"webapp/src/cookies"
	"webapp/src/router"
	"webapp/src/utils"

	"github.com/gorilla/securecookie"
)

func init() {
	config.Carregar()
	haskKey := securecookie.GenerateRandomKey(64)
	fmt.Println(hex.EncodeToString(haskKey))
	blockKey := securecookie.GenerateRandomKey(64)
	fmt.Println(hex.EncodeToString(blockKey))
}

func main() {
	cookies.Configurar()
	utils.CarregarTemplates()
	r := router.Gerar()

	fmt.Printf("Rodando Webapp na porta %d\n", config.Porta)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Porta), r))
}
