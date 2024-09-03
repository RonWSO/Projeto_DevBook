package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
	"webapp/src/config"
	"webapp/src/requisicoes"
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

	var (
		usuario     Usuario
		seguidores  []Usuario
		seguindo    []Usuario
		publicacoes []Publicacao
	)

	for i := 0; i < 4; i++ {
		select {
		case usuarioCarregado := <-canalUsuario:
			if usuarioCarregado.ID == 0 {
				return Usuario{}, errors.New("erro ao buscar o usuário")
			}
			usuario = usuarioCarregado

		case seguidoresCarregados := <-canalSeguidores:
			if seguidoresCarregados == nil {
				return Usuario{}, errors.New("erro ao buscar seguidores")
			}
			seguidores = seguidoresCarregados

		case seguindoCarregados := <-canalSeguindo:
			if seguindoCarregados == nil {
				return Usuario{}, errors.New("erro ao buscar seguindo")
			}
			seguindo = seguindoCarregados

		case publicacoesCarregados := <-canalPublicacao:
			if publicacoesCarregados == nil {
				return Usuario{}, errors.New("erro ao buscar publicacoes")
			}
			publicacoes = publicacoesCarregados

		}
	}
	usuario.Seguidores = seguidores
	usuario.Seguindo = seguindo
	usuario.Publicacoes = publicacoes

	return usuario, nil
}

// Função responsável por buscar todos os dados de um Usuário
// Chama o a API pra buscar os dados base do usuário
func BuscarDadosDoUsuario(canal chan<- Usuario, usuarioID uint64, r *http.Request) {
	url := fmt.Sprintf("%s/usuario/%d", config.APIURL, usuarioID)
	response, erro := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodGet, url, nil)
	if erro != nil {
		canal <- Usuario{}
		return
	}

	defer response.Body.Close()

	var usuario Usuario
	if erro = json.NewDecoder(response.Body).Decode(&usuario); erro != nil {
		canal <- Usuario{}
	}

	canal <- usuario
}

// Função responsável por buscar todos os seguidores de um Usuário
func BuscarSeguidores(canal chan<- []Usuario, usuarioID uint64, r *http.Request) {
	url := fmt.Sprintf("%s/usuario/%d/seguidores", config.APIURL, usuarioID)
	response, erro := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodGet, url, nil)
	if erro != nil {
		canal <- nil
		return
	}
	defer response.Body.Close()

	var seguidores []Usuario
	if erro = json.NewDecoder(response.Body).Decode(&seguidores); erro != nil {
		canal <- nil
		return
	}
	canal <- seguidores
}

// Função responsável por buscar todos os seguindo de um Usuário
func BuscarSeguindo(canal chan<- []Usuario, usuarioID uint64, r *http.Request) {
	url := fmt.Sprintf("%s/usuario/%d/seguindo", config.APIURL, usuarioID)
	response, erro := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodGet, url, nil)
	if erro != nil {
		canal <- nil
		return
	}
	defer response.Body.Close()

	var seguindo []Usuario
	if erro = json.NewDecoder(response.Body).Decode(&seguindo); erro != nil {
		canal <- nil
		return
	}
	canal <- seguindo
}

// Função responsável por buscar todas publicações de um Usuário
func BuscarPublicacao(canal chan<- []Publicacao, usuarioID uint64, r *http.Request) {
	url := fmt.Sprintf("%s/usuario/%d/publicacoes", config.APIURL, usuarioID)
	response, erro := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodGet, url, nil)
	if erro != nil {
		canal <- nil
		return
	}
	defer response.Body.Close()
	var publicacoes []Publicacao
	if erro = json.NewDecoder(response.Body).Decode(&publicacoes); erro != nil {
		canal <- nil
		return
	}
	canal <- publicacoes
}
