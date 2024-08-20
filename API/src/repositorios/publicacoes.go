package repositorios

import (
	"api/src/models"
	"database/sql"
	"errors"
	"fmt"
)

// Publicações representa um repositorio de publicações
type Publicacoes struct {
	db *sql.DB
}

// Novo repositório de publicações
func NovoRepositorioDePublicacoes(db *sql.DB) *Publicacoes {
	return &Publicacoes{db}
}

// Criar insere uma publicação no banco de dados
func (repositorio Publicacoes) Criar(publicacao models.Publicacao) (uint64, error) {
	statement, erro := repositorio.db.Prepare("INSERT INTO publicacoes (titulo, conteudo, autor_id) VALUES (?,?,?)")
	if erro != nil {
		return 0, nil
	}

	defer statement.Close()

	resultado, erro := statement.Exec(publicacao.Titulo, publicacao.Conteudo, publicacao.AutorId)
	if erro != nil {
		return 0, nil
	}

	ultimoIdInserido, erro := resultado.LastInsertId()
	if erro != nil {
		return 0, nil
	}
	return uint64(ultimoIdInserido), nil
}

// trás uma unica publicação do banco de dados
func (repositorio Publicacoes) BuscarPorID(publicacaoID uint64) (models.Publicacao, error) {
	linha, erro := repositorio.db.Query("SELECT p.*, u.nick FROM publicacoes p INNER JOIN usuarios u ON u.id = p.autor_id WHERE p.id = ?", publicacaoID)
	if erro != nil {
		return models.Publicacao{}, erro
	}
	defer linha.Close()
	var publicacao models.Publicacao

	if linha.Next() {
		if erro = linha.Scan(
			&publicacao.ID,
			&publicacao.Titulo,
			&publicacao.Conteudo,
			&publicacao.AutorId,
			&publicacao.Curtidas,
			&publicacao.CriadaEm,
			&publicacao.AutorNick,
		); erro != nil {
			return models.Publicacao{}, erro
		}
	} else {
		return models.Publicacao{}, errors.New("publicação não pode ser localizada")
	}
	return publicacao, nil
}

// Busca publicações de quem o usuário segue e do próprio usuário
func (repositorio Publicacoes) Buscar(usuarioID uint64) ([]models.Publicacao, error) {

	linhas, erro := repositorio.db.Query(`SELECT p.*, u.nick FROM publicacoes p 
	INNER JOIN usuarios u ON u.id = p.autor_id
	WHERE u.id = ? OR p.autor_id IN 
	(SELECT usuario_id FROM seguidores 
	WHERE seguidor_id = ?)`, usuarioID, usuarioID)
	if erro != nil {
		return nil, erro
	}
	defer linhas.Close()

	var publicacoes []models.Publicacao

	for linhas.Next() {
		var publicacao models.Publicacao

		if erro = linhas.Scan(
			&publicacao.ID,
			&publicacao.Titulo,
			&publicacao.Conteudo,
			&publicacao.AutorId,
			&publicacao.Curtidas,
			&publicacao.CriadaEm,
			&publicacao.AutorNick,
		); erro != nil {
			return nil, erro
		}
		fmt.Println(publicacao)
		publicacoes = append(publicacoes, publicacao)
	}
	return publicacoes, nil
}

func (repositorio Publicacoes) BuscarPorUsuario(usuarioID uint64) ([]models.Publicacao, error) {
	linhas, erro := repositorio.db.Query("SELECT p.*, u.nick FROM publicacoes p INNER JOIN usuarios u ON u.id = p.autor_id WHERE p.autor_id = ?", usuarioID)
	if erro != nil {
		return nil, erro
	}
	defer linhas.Close()

	var publicacoes []models.Publicacao

	for linhas.Next() {
		var publicacao models.Publicacao

		if erro = linhas.Scan(
			&publicacao.ID,
			&publicacao.Titulo,
			&publicacao.Conteudo,
			&publicacao.AutorId,
			&publicacao.Curtidas,
			&publicacao.CriadaEm,
			&publicacao.AutorNick,
		); erro != nil {
			return nil, erro
		}
		publicacoes = append(publicacoes, publicacao)
	}
	return publicacoes, nil
}

// Atualiza a publicação do usuário
func (repositorio Publicacoes) Atualizar(publicacaoID uint64, publicacao models.Publicacao) error {
	statement, erro := repositorio.db.Prepare("UPDATE publicacoes SET titulo = ?, conteudo = ? WHERE id = ?")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro := statement.Exec(publicacao.Titulo, publicacao.Conteudo, publicacaoID); erro != nil {
		return erro
	}
	return nil
}

// Deleta a publicação do usuário
func (repositorio Publicacoes) Deletar(publicacaoID uint64) error {
	statement, erro := repositorio.db.Prepare("DELETE FROM publicacoes WHERE publicacoes.id = ?")
	if erro != nil {
		return erro
	}
	defer statement.Close()
	if _, erro = statement.Exec(publicacaoID); erro != nil {
		return erro
	}
	return nil
}

// Adiciona uma curtida na publicação
func (repositorio Publicacoes) Curtir(publicacaoID uint64) error {
	statement, erro := repositorio.db.Prepare("UPDATE publicacoes SET curtidas = curtidas + 1 WHERE id = ?")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(publicacaoID); erro != nil {
		return erro
	}
	return nil
}

// Diminui uma curtida na publicação
func (repositorio Publicacoes) Descurtir(publicacaoID uint64) error {
	statement, erro := repositorio.db.Prepare("UPDATE publicacoes  SET curtidas = CASE WHEN curtidas > 0 THEN curtidas - 1 ELSE 0 END WHERE id = ?")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(publicacaoID); erro != nil {
		return erro
	}
	return nil
}
