package repositorios

import (
	"api/src/models"
	"database/sql"
	"fmt"
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
func (repositorio usuarios) Criar(usuario models.Usuario) (uint64, error) {
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

// Buscar traz todos os usuarios que atendem ao filtro de nome ou nick
func (repositorio usuarios) Buscar(nomeOuNick string) ([]models.Usuario, error) {
	nomeOuNick = fmt.Sprintf("%%%s%%", nomeOuNick)

	linhas, erro := repositorio.db.Query("SELECT id, nome, nick, email, criadoEm FROM usuarios WHERE nome LIKE ? OR nick LIKE ?", nomeOuNick, nomeOuNick)
	if erro != nil {
		return nil, erro
	}

	defer linhas.Close()

	var usuarios []models.Usuario

	for linhas.Next() {
		var usuario models.Usuario

		if erro = linhas.Scan(
			&usuario.ID, &usuario.Nome, &usuario.Nick, &usuario.Email, &usuario.CriadoEm); erro != nil {
			return nil, erro
		}

		usuarios = append(usuarios, usuario)
	}

	return usuarios, nil
}

// Buscar traz um usuario que atende ao filtro de id
func (repositorio usuarios) BuscarPorID(ID uint64) (models.Usuario, error) {
	linha, erro := repositorio.db.Query("SELECT id, nome, nick, email, criadoEm FROM usuarios WHERE id = ?", ID)
	if erro != nil {
		return models.Usuario{}, erro
	}
	defer linha.Close()

	usuario := models.Usuario{}

	if linha.Next() {
		if erro = linha.Scan(
			&usuario.ID, &usuario.Nome, &usuario.Nick, &usuario.Email, &usuario.CriadoEm); erro != nil {
			return models.Usuario{}, erro
		}
	}

	return usuario, nil
}

// Altera as informações de um usuario no banco de dados, exceto senha
func (repositorio usuarios) Atualizar(ID uint64, u models.Usuario) error {
	statement, erro := repositorio.db.Prepare("UPDATE usuarios SET nome = ?, nick = ?, email = ? WHERE id = ?")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro := statement.Exec(u.Nome, u.Nick, u.Email, ID); erro != nil {
		return erro
	}

	return nil
}

// Exclui um usuario no banco de dados
func (repositorio usuarios) Excluir(ID uint64) error {
	statement, erro := repositorio.db.Prepare("DELETE FROM usuarios WHERE id = ?")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro := statement.Exec(ID); erro != nil {
		return erro
	}

	return nil
}

// Busca usuario por email e retorna email e senha com hash
func (repositorio usuarios) BuscarPorEmail(email string) (models.Usuario, error) {
	linha, erro := repositorio.db.Query("SELECT id, senha FROM usuarios WHERE email = ?", email)
	if erro != nil {
		return models.Usuario{}, erro
	}
	defer linha.Close()

	var usuario models.Usuario
	if linha.Next() {
		if erro = linha.Scan(&usuario.ID, &usuario.Senha); erro != nil {
			return models.Usuario{}, erro
		}

	}

	return usuario, nil
}
