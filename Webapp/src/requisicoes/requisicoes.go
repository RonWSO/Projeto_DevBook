package requisicoes

import (
	"io"
	"net/http"
	"webapp/src/cookies"
)

// FazerRequisicaoComAutenticacao é utilizada para por o token na requisição
func FazerRequisicaoComAutenticacao(r *http.Request, metodo, url string, dados io.Reader) (*http.Response, error) {
	//Cria a requisição que irá ser feita para a API
	request, erro := http.NewRequest(metodo, url, dados)
	if erro != nil {
		return nil, erro
	}

	//Lê os cookies da requisição recebida do WebApp
	cookie, _ := cookies.Ler(r)
	//Adiciona o token recebido a requisição
	request.Header.Add("Authorization", "Bearer "+cookie["token"])

	//Cria o cliente para fazer a requisição para a API
	client := &http.Client{}
	//Faz a requisição e pega a resposta
	response, erro := client.Do(request)
	if erro != nil {
		return nil, erro
	}
	//Retorna a resposta
	return response, nil
}
