package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"webapp/src/config"
	"webapp/src/respostas"
)

// Chama a API para cadastrar um usuário no banco de dados
func CadastroDeUsuario(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	usuario, erro := json.Marshal(map[string]string{
		"nome":  r.FormValue("nome"),
		"email": r.FormValue("email"),
		"nick":  r.FormValue("nick"),
		"senha": r.FormValue("senha"),
	})
	fmt.Println("21")
	if erro != nil {
		respostas.JSON(w, http.StatusBadRequest, respostas.ErroApi{Erro: "Dados enviados não podem ser processados"})
		return
	}
	fmt.Println("26")
	url := fmt.Sprintf("%s/usuario", config.APIURL)
	fmt.Println(url)
	response, erro := http.Post(url, "application/json", bytes.NewBuffer(usuario))
	if erro != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroApi{Erro: erro.Error()})
		return
	}
	defer response.Body.Close()
	fmt.Println("34")
	if response.StatusCode >= 400 {
		respostas.TratarRespostaErro(w, response)
		return
	}

	fmt.Println("39")
	respostas.JSON(w, response.StatusCode, nil)
}
