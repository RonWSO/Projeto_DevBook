$('#formulario-cadastro').on('submit', criarUsuario);

function criarUsuario(evento){
    evento.preventDefault();

    if($('#senha').val() != $('#confirmar-senha').val()){
        swal("As senhas não coincidem", 
            {
                icon: "warning",
                title: "Ops..."
            }
        );
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
        swal("Usuário cadastrado com sucesso",{
            icon: "success",
            button: "Ok!",
        }).then(function(){ 
            $.ajax({
                url: "/login",
                method: "POST",
                data:{
                    email: $('#email').val(),
                    senha: $('#senha').val(),
                }
            }).done(function(){
                window.location = "/home";
            }).fail(function(){
                swal("Não foi possível fazer o login",{
                    icon: "error",
                    button: "Ok!",
                })
                .then(window.location = "/")
            });
        });
    }).fail(function(retorno){// 400 ou 500 São considerados status de falha
        console.log(retorno)
        swal("Erro ao cadastrar o usuário",{
            icon: "error",
            dangerMode: true,
        })
    });
}