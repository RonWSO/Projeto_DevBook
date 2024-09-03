package rotas

import (
	"api/src/controllers"
	"net/http"
)

// Cria um slice do tipo Rota
var rotasUsuarios = []Rota{
	{
		//Rota de Criação de Usuario
		URI:                "/usuario",
		Metodo:             http.MethodPost,
		Funcao:             controllers.CriarUsuario,
		RequerAutenticacao: false,
	},
	{
		//Rota para buscar todos os usuarios
		URI:                "/usuario",
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscarUsuarios,
		RequerAutenticacao: true,
	},
	{
		//Rota para buscar um unico usuario
		URI:                "/usuario/{usuarioId}",
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscarUsuario,
		RequerAutenticacao: true,
	},
	{
		//Rota para atualizar o Usuario
		URI:                "/usuario/{usuarioId}",
		Metodo:             http.MethodPut,
		Funcao:             controllers.AtualizarUsuario,
		RequerAutenticacao: true,
	},
	{
		//Rota para Deletar o usuario
		URI:                "/usuario/{usuarioId}",
		Metodo:             http.MethodDelete,
		Funcao:             controllers.DeletarUsuario,
		RequerAutenticacao: true,
	},
	{
		//Rota para seguir
		URI:                "/usuario/{usuarioId}/seguir",
		Metodo:             http.MethodPost,
		Funcao:             controllers.SeguirUsuario,
		RequerAutenticacao: true,
	},
	{
		//Rota para tirar o seguir
		URI:                "/usuario/{usuarioId}/desseguir",
		Metodo:             http.MethodPost,
		Funcao:             controllers.DesseguirUsuario,
		RequerAutenticacao: true,
	},
	{
		//Rota mostrar os seguidores
		URI:                "/usuario/{usuarioId}/seguidores",
		Metodo:             http.MethodGet,
		Funcao:             controllers.Seguidores,
		RequerAutenticacao: true,
	},
	{
		//Rota para mostrar quem está seguindo o usuário
		URI:                "/usuario/{usuarioId}/seguindo",
		Metodo:             http.MethodGet,
		Funcao:             controllers.Seguindo,
		RequerAutenticacao: true,
	},
	{
		//Rota para atualizar a senha
		URI:                "/usuario/{usuarioId}/atualizarSenha",
		Metodo:             http.MethodPost,
		Funcao:             controllers.AtualizarSenha,
		RequerAutenticacao: true,
	},
}
