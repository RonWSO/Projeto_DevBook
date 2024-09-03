package models

import (
	"net/http"
	"time"
)

// Representa um perfil de usuario na aplicação
type Usuario struct {
	ID          uint64       `json:"id"`
	Nome        string       `json:"nome"`
	Email       string       `json:"email"`
	Nick        string       `json:"nick"`
	CriadoEm    time.Time    `json:"criadoEm"`
	Seguidores  []Usuario    `json:"seguidores"`
	Seguindo    []Usuario    `json:"seguindo"`
	Publicacoes []Publicacao `json:"publicacoes"`
}

// Faz 4 requisições na API para montar o usuário
func BuscarUsuarioCompleto(usuarioID uint64, r *http.Request) (Usuario, error) {
	canalUsuario := make(chan Usuario)
	canalSeguidores := make(chan []Usuario)
	canalSeguindo := make(chan []Usuario)
	canalPublicacao := make(chan []Publicacao)

	go BuscarDadosDoUsuario(canalUsuario, usuarioID, r)
	go BuscarSeguidores(canalSeguidores, usuarioID, r)
	go BuscarSeguindo(canalSeguindo, usuarioID, r)
	go BuscarPublicacao(canalPublicacao, usuarioID, r)
	return Usuario{}, nil
}

// Função responsável por buscar todos os dados de um Usuário
func BuscarDadosDoUsuario(canal <-chan Usuario, usuarioID uint64, r *http.Request) {

}

// Função responsável por buscar todos os seguidores de um Usuário
func BuscarSeguidores(canal <-chan []Usuario, usuarioID uint64, r *http.Request) {

}

// Função responsável por buscar todos os seguindo de um Usuário
func BuscarSeguindo(canal <-chan []Usuario, usuarioID uint64, r *http.Request) {

}

// Função responsável por buscar todas publicações de um Usuário
func BuscarPublicacao(canal <-chan []Publicacao, usuarioID uint64, r *http.Request) {

}
