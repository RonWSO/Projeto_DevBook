package banco

import (
	"api/src/config"
	"database/sql"

	//pacote de conexão usando mysql
	_ "github.com/go-sql-driver/mysql" //Driver
)

// Conecta e abre a conexão com o banco
func Conectar() (*sql.DB, error) {
	//Abre conexão usando as variaveis de ambiente
	db, erro := sql.Open("mysql", config.StringConexaoBanco)
	if erro != nil {
		return nil, erro
	}
	//Ping para checar se a conexão está aberta
	if erro = db.Ping(); erro != nil {
		db.Close()
		return nil, erro
	}
	return db, nil
}
