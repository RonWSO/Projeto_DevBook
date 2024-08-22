package utils

import (
	"net/http"
	"text/template"
)

//Templates servem para usar o html e retornar as páginas solicitadas.

var templates *template.Template

// Carrega os templates e coloca dentro da variavel templates
func CarregarTemplates() {
	templates = template.Must(template.ParseGlob("views/*.html"))
}

// Executar template renderiza um template na tela
func ExecutarTemplate(w http.ResponseWriter, template string, dados interface{}) {
	templates.ExecuteTemplate(w, template, dados)
}
