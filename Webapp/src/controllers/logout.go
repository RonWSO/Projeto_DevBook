package controllers

import (
	"net/http"
	"webapp/src/cookies"
)

// Fazer remove as informações salvas no browser do usuário
func FazerLogout(w http.ResponseWriter, r *http.Request) {
	cookies.Deletar(w)
	http.Redirect(w, r, "/login", http.StatusFound)
}
