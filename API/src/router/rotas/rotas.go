package rotas

import (
	"api/src/middlewares"
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
	rotas := rotasUsuarios
	rotas = append(rotas, RotaLogin)
	//com reticencias no final da variável o appende percorre o slice de rotas um por um
	rotas = append(rotas, rotasPublicacoes...)
	for _, rota := range rotas {
		if rota.RequerAutenticacao {
			//cria um handlerfunc, onde chama o middlewares logger onde é escrito o log da requisição, passando a função de autenticar, para quando for escrito o log ela seja chamada, Dentro da função de chamada
			r.HandleFunc(rota.URI,
				middlewares.Logger(middlewares.Autenticar(rota.Funcao)),
			).Methods(rota.Metodo)
		} else {
			//Coloca todas as rotas dentro do router recebido
			r.HandleFunc(rota.URI, middlewares.Logger(rota.Funcao)).Methods(rota.Metodo)
		}

	}

	return r
}
