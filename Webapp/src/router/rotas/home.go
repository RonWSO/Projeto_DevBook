package rotas

import (
	"net/http"
	"webapp/src/controllers"
)

var rotasHome = []Rota{
	{
		URI:                "/home",
		Metodo:             http.MethodGet,
		Funcao:             controllers.CarregarTelaHome,
		RequerAutenticacao: true,
	},
}
