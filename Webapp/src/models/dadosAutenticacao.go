package models

//Struct que contém todos os dados de autenticação
type DadosAutenticacao struct {
	ID    string `json:"id"`
	Token string `json:"token"`
}
