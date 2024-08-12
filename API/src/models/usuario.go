package models

import (
	"errors"
	"strings"
	"time"

	"github.com/badoux/checkmail"
)

// Representa a estrutura de um usuario no banco de dados
type Usuario struct {
	ID    uint64 `json:"id,omitempty"`
	Nome  string `json:"nome,omitempty"`
	Nick  string `json:"nick,omitempty"`
	Email string `json:"email,omitempty"`
	Senha string `json:"senha,omitempty"`
	//Deixa o Criado em como um ponteiro para que o valor possa ser nil
	CriadoEm *time.Time `json:"CriadoEm,omitempty"`
}

// Prepara a instancia de usuario para ser inserida no banco
func (u *Usuario) Preparar(etapa string) error {

	u.formatar()

	if erro := u.validar(etapa); erro != nil {
		return erro
	}

	return nil
}

func (u *Usuario) TesteValidar(etapa string) string {
	if u.Nome == "" {
		return string("o nome do usuário é obrigatório e não pode estar em branco")
	}
	if u.Nick == "" {
		return string("o nick do usuário é obrigatório e não pode estar em branco")
	}
	if u.Email == "" {
		return string("o email do usuário é obrigatório e não pode estar em branco")
	}
	if u.Senha == "" {
		return string("o senha do usuário é obrigatório e não pode estar em branco")
	}
	return string("usuario valido!")
}

func (u *Usuario) TesteFormatar(etapa string) *Usuario {
	u.Nome = strings.TrimSpace(u.Nome)
	u.Nick = strings.TrimSpace(u.Nick)
	u.Email = strings.TrimSpace(u.Email)

	return u
}

func (u *Usuario) validar(etapa string) error {
	if u.Nome == "" {
		return errors.New("o nome do usuário é obrigatório e não pode estar em branco")
	}
	if u.Nick == "" {
		return errors.New("o nick do usuário é obrigatório e não pode estar em branco")
	}
	if u.Email == "" {
		return errors.New("o email do usuário é obrigatório e não pode estar em branco")
	}

	if erro := checkmail.ValidateFormat(u.Email); erro != nil {
		return errors.New("o email do usuário não está em formato válido")
	}

	if etapa == "cadastro" && u.Senha == "" {
		return errors.New("o senha do usuário é obrigatório e não pode estar em branco")
	}
	return nil
}

func (u *Usuario) formatar() {
	u.Nome = strings.TrimSpace(u.Nome)
	u.Nick = strings.TrimSpace(u.Nick)
	u.Email = strings.TrimSpace(u.Email)
}
