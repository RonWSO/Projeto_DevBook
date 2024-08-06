package rotas

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Rota define a estrutura padrão de todas as rotas da API
type Rota struct {
	URI                string
	Metodo             string
	Funcao             func(http.ResponseWriter, *http.Request)
	RequerAutenticacao bool
}

// Configurar, coloca todas as rotas dentro do usuario
func Configurar(r *mux.Router) *mux.Router {
	//Acessa o rotas usuarios para pegar todas as struct
	rotas := &rotasUsuarios

	for _, rota := range *rotas {
		//Coloca todas as rotas dentro do router recebido
		r.HandleFunc(rota.URI, rota.Funcao).Methods(rota.Metodo)
	}

	return r
}
