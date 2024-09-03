package rotas

import (
	"net/http"
	"webapp/src/controllers"
)

var rotasUsuarios = []Rota{
	{
		URI:                "/criar-usuario",
		Metodo:             http.MethodGet,
		Funcao:             controllers.CarregarPaginaDeCadastroDeUsuario,
		RequerAutenticacao: false,
	},
	{
		URI:                "/usuarios",
		Metodo:             http.MethodPost,
		Funcao:             controllers.CadastroDeUsuario,
		RequerAutenticacao: false,
	},
	{
		URI:                "/buscar-usuario",
		Metodo:             http.MethodGet,
		Funcao:             controllers.CarregarPaginaDePesquisaUsuarios,
		RequerAutenticacao: true,
	},
	{
		URI:                "/usuario/{idUsuario}",
		Metodo:             http.MethodGet,
		Funcao:             controllers.CarregarPaginaDeUsuarios,
		RequerAutenticacao: true,
	},
}
