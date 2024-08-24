$('#formulario-cadastro').on('submit', criarUsuario);

function criarUsuario(evento){
    evento.preventDefault();

    if($('#senha').val() != $('#confirmar-senha').val()){
        alert("As senhas não coincidem");
        return
    }
    $.ajax({
        url:"/usuarios",
        method:"POST",
        data:{
            nome:$('#nome').val(),
            email:$('#email').val(),
            nick:$('#nick').val(),
            senha:$('#senha').val(),
        }
    }).done(function(retorno){ // 200 201 204 São considerados status de sucesso
        alert("Usuário cadastrado com sucesso")
    }).fail(function(retorno){// 400 ou 500 São considerados status de falha
        console.log(retorno)
        alert("Erro ao cadastrar o usuário")
    });
}