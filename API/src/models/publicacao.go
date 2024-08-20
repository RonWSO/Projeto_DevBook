package models

import (
	"errors"
	"strings"
	"time"
)

// Representa uma publicação feita por um usuário
type Publicacao struct {
	ID        uint64     `json:"id,omitempty"`
	Titulo    string     `json:"titulo,omitempty"`
	Conteudo  string     `json:"conteudo,omitempty"`
	AutorId   uint64     `json:"autorId,omitempty"`
	AutorNick string     `json:"autorNick,omitempty"`
	Curtidas  uint64     `json:"curtidas"`
	CriadaEm  *time.Time `json:"criadaEm,omitempty"`
}

func (publicacao *Publicacao) Preparar() error {
	if erro := publicacao.validar(); erro != nil {
		return erro
	}

	publicacao.formatar()
	return nil
}

// Valida se os campos da publicação estão preenchidos
func (publicacao *Publicacao) validar() error {
	if publicacao.Titulo == "" {
		return errors.New("titulo é obrigatório")
	}
	if publicacao.Conteudo == "" {
		return errors.New("conteúdo é obrigatório")
	}
	return nil
}

// Retira espaços antes de pois do conteúdo e do título
func (publicacao *Publicacao) formatar() {
	publicacao.Titulo = strings.TrimSpace(publicacao.Titulo)
	publicacao.Conteudo = strings.TrimSpace(publicacao.Conteudo)

}
