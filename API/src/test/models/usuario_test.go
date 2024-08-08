package models

import (
	"api/src/models"
	"reflect"
	"testing"
)

type cenarioDeTeste struct {
	Usuario         models.Usuario
	RetornoEsperado interface{}
}

func TestValidarUsuario(t *testing.T) {
	t.Parallel()

	cenariosDeTeste := []cenarioDeTeste{
		{models.Usuario{Nome: "Cleiton", Nick: "Cleitin", Email: "Cleiton@gmail.com", Senha: "12345678"}, "usuario valido!"},
		{models.Usuario{Nome: "Cleiton", Nick: "Cleitin", Email: "Cleiton@gmail.com", Senha: ""}, "o senha do usuário é obrigatório e não pode estar em branco"},
		{models.Usuario{Nome: "", Nick: "Cleitino", Email: "Cleitoni@gmail.com", Senha: "123456789"}, "o nome do usuário é obrigatório e não pode estar em branco"},
		{models.Usuario{Nome: "Cleiton", Nick: "", Email: "Cleiton@gmail.com", Senha: "12345678"}, "o nick do usuário é obrigatório e não pode estar em branco"},
		{models.Usuario{Nome: "Cleiton", Nick: "Cleitin", Email: "", Senha: "12345678"}, "o email do usuário é obrigatório e não pode estar em branco"},
	}

	for _, caso := range cenariosDeTeste {
		RetornoRecebido := caso.Usuario.TesteValidar()
		if RetornoRecebido != caso.RetornoEsperado {
			t.Errorf("Retorno recebido:%s Retorno esperado:%s", RetornoRecebido, caso.RetornoEsperado)
		}
	}
}
func TestFormatarUsuario(t *testing.T) {
	t.Parallel()

	usuarioPadrao := &models.Usuario{Nome: "Cleiton", Nick: "Cleitin", Email: "Cleiton@gmail.com", Senha: "12345678"}
	cenariosDeTeste := []cenarioDeTeste{
		{models.Usuario{Nome: "         Cleiton        ", Nick: "Cleitin", Email: "Cleiton@gmail.com", Senha: "12345678"}, usuarioPadrao},
		{models.Usuario{Nome: "Cleiton", Nick: "         Cleitin       ", Email: "Cleiton@gmail.com", Senha: "12345678"}, usuarioPadrao},
		{models.Usuario{Nome: "Cleiton", Nick: "Cleitin", Email: "       Cleiton@gmail.com        ", Senha: "12345678"}, usuarioPadrao},
	}

	for _, caso := range cenariosDeTeste {
		RetornoRecebido := caso.Usuario.TesteFormatar()
		if !reflect.DeepEqual(RetornoRecebido, caso.RetornoEsperado) {
			t.Errorf("Usuário formatado incorretamente.\nEsperado: %+v\nRecebido: %+v", caso.RetornoEsperado, RetornoRecebido)
		}
	}
}
