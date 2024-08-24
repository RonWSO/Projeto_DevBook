$('#login').on('submit', fazerLogin);

function fazerLogin(evento){
    evento.preventDefault();
    $.ajax({
        url:"/login",
        method:"POST",
        data:{
            email:$('#email').val(),
            senha:$('#senha').val(),
        }
    }).done(function(retorno){ // 200 201 204 São considerados status de sucesso
        window.location = '/home'
    }).fail(function(retorno){// 400 ou 500 São considerados status de falha
        console.log(retorno)
        alert("Não foi possível fazer o login")
    });
}