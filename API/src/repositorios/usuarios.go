package repositorios

import (
	"api/src/models"
	"database/sql"
)

type usuarios struct {
	db *sql.DB
}

// Essa função vai receber o banco recebido pelo controller, com esse banco ele vai
// retornar uma "instancia" de usuarios por meio desse struct o controller vai manipular o tipo usuario no banco
func NovoRepositorioDeUsuarios(db *sql.DB) *usuarios {
	return &usuarios{db}
}

// Insere um usuario no banco de dados
func (repositorio usuarios) Criar(usuario *models.Usuario) (uint64, error) {
	statement, erro := repositorio.db.Prepare("INSERT INTO usuarios (nome, nick, email, senha) VALUES (?, ?, ?, ?)")
	if erro != nil {
		return 0, erro
	}
	defer statement.Close()

	resultado, erro := statement.Exec(usuario.Nome, usuario.Nick, usuario.Email, usuario.Senha)
	if erro != nil {
		return 0, erro
	}
	ultimoIDInserido, erro := resultado.LastInsertId()
	if erro != nil {
		return 0, erro
	}

	return uint64(ultimoIDInserido), nil
}
