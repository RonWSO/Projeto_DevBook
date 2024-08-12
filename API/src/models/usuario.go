package models

import (
	"api/src/security"
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

	if erro := u.validar(etapa); erro != nil {
		return erro
	}
	if erro := u.formatar(etapa); erro != nil {
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
	if etapa == "cadastro" && u.Senha == "" {
		return string("o senha do usuário é obrigatório e não pode estar em branco")
	}
	return string("usuario valido!")
}

func (u *Usuario) TesteFormatar(etapa string) (*Usuario, error) {
	u.Nome = strings.TrimSpace(u.Nome)
	u.Nick = strings.TrimSpace(u.Nick)
	u.Email = strings.TrimSpace(u.Email)
	if etapa == "cadastro" {
		senhaComHash, erro := security.Hash(u.Senha)
		if erro != nil {
			return &Usuario{}, erro
		}
		u.Senha = string(senhaComHash)
	}

	return u, nil
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

func (u *Usuario) formatar(etapa string) error {
	u.Nome = strings.TrimSpace(u.Nome)
	u.Nick = strings.TrimSpace(u.Nick)
	u.Email = strings.TrimSpace(u.Email)
	if etapa == "cadastro" {
		senhaComHash, erro := security.Hash(u.Senha)
		if erro != nil {
			return erro
		}
		u.Senha = string(senhaComHash)
	}
	return nil
}
