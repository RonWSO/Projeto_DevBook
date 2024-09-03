package cookies

import (
	"net/http"
	"time"
	"webapp/src/config"

	"github.com/gorilla/securecookie"
)

// Cria a variavel de secureCookie
var s *securecookie.SecureCookie

// Configurar utiliza as variaveis de ambiente para a criação do SecureCookie
func Configurar() {
	s = securecookie.New(config.HashKey, config.BlockKey)
}

// Salvar registra as informações de autenticação
func Salvar(w http.ResponseWriter, ID, token string) error {
	//Cria um mapa json com o id e token retornados pela api
	dados := map[string]string{
		"id":    ID,
		"token": token,
	}
	//Cria um mapa json com o id e token retornados pela api
	dadosCodificados, erro := s.Encode("dados", dados)
	if erro != nil {
		return erro
	}
	//Cria um mapa json com o id e token retornados pela api
	http.SetCookie(w, &http.Cookie{
		Name:     "dados",
		Value:    dadosCodificados,
		Path:     "/",
		HttpOnly: true,
	})
	return nil
}

func Ler(r *http.Request) (map[string]string, error) {
	cookie, erro := r.Cookie("dados")
	if erro != nil {
		return nil, erro
	}
	valores := make(map[string]string)
	if erro = s.Decode("dados", cookie.Value, &valores); erro != nil {
		return nil, erro
	}

	return valores, nil
}

// Deleta deleta os cookies no usuáio
func Deletar(w http.ResponseWriter) {
	http.SetCookie(w, &http.Cookie{
		Name:     "dados",
		Value:    "",
		Path:     "/",
		HttpOnly: true,
		Expires:  time.Unix(0, 0),
	})
}
