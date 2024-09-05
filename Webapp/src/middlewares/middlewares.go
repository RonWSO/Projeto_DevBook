package middlewares

import (
	"encoding/json"
	"log"
	"net/http"
	"webapp/src/cookies"
	"webapp/src/respostas"
)

// Logger escreve informações no terminal
func Logger(proximaFuncao http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("\n %s %s %s", r.Method, r.RequestURI, r.Host)
		proximaFuncao(w, r)
	}
}

// Autenticar verifica a existencia de cookies
func Autenticar(proximaFuncao http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if _, erro := cookies.Ler(r); erro != nil {
			http.Redirect(w, r, "/", http.StatusFound)
			return
		}
		proximaFuncao(w, r)
	}
}

func TesteCookieExpirado(w http.ResponseWriter, response *http.Response, request *http.Request) {
	if response.StatusCode == http.StatusUnauthorized {
		var erroApi respostas.ErroApi
		if err := json.NewDecoder(response.Body).Decode(&erroApi); err == nil {
			if erroApi.Erro == "token expired" {
				cookies.Deletar(w)
				http.Redirect(w, request, "/", http.StatusFound)
			}

		}
	}
}
