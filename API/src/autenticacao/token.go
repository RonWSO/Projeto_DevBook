package autenticacao

import (
	"api/src/config"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

// Cria um token com as permissões de um usuário
func CriarToken(usuarioID uint64) (string, error) {
	permissoes := jwt.MapClaims{}
	permissoes["authorized"] = true
	permissoes["exp"] = time.Now().Add(time.Hour * 6).Unix()
	permissoes["usuarioId"] = usuarioID
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permissoes)
	tokenAssinado, erro := token.SignedString([]byte(config.SecretKey))
	if erro != nil {
		return "", erro
	}
	return tokenAssinado, nil
}

// Verifica se o token passado na requisição é válido
func ValidarToken(r *http.Request) (*jwt.Token, error) {
	//Retira a string do token do request e transforma em string
	tokenString := extrairToken(r)
	//Transforma a string do token em um token jwt, a função retornar chave de verificação, é uma função que checa a família do token e checa se está correto
	token, erro := jwt.Parse(tokenString, retornarChaveDeVerificacao)
	if erro != nil {
		return nil, erro
	}
	// retorna o map com os claims do token e um book para saber se está true, checa se o token é Válido, esse metodo é utilizável quando se válida o token
	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		//Se estiver tudo ok retorna o token e o nil
		return token, nil
	}
	//Se o token estiver inválido retorna o erro.
	return nil, errors.New("token inválido")
}

// Pega o token de dentro da requisição
func extrairToken(r *http.Request) string {
	token := r.Header.Get("Authorization")
	//O token beared que estamos recebendo vem escrito beared antes, então pegamos a string e separamos em duas palavras para checar e poder retornar apenas a string do token
	if len(strings.Split(token, " ")) == 2 {
		return strings.Split(token, " ")[1]
	}
	return ""
}

// Checa se o token é da família esperada
func retornarChaveDeVerificacao(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("método de assinatura inesperado %v", token.Header["alg"])
	}

	return config.SecretKey, nil
}

// Pega o id do usuário no token, checando primeiramente se o token é válido
func ExtrairUsuarioToken(r *http.Request) (uint64, error) {
	//Valida token utilizando a função
	token, erro := ValidarToken(r)
	if erro != nil {
		return 0, erro
	}

	if permissoes, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		//O ParseUint espera uma string para converter para inteiro e o Mapclaims é uma interface genérica, então precisa converter para float, para converter para string e depois ir para uint64
		usuarioID, erro := strconv.ParseUint(fmt.Sprintf("%.0f", permissoes["usuarioId"]), 10, 64)
		if erro != nil {
			return 0, erro
		}
		return usuarioID, nil
	}
	return 0, errors.New("token inválido")
}
