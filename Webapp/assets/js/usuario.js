$('#parar-de-seguir').on('click',pararDeSeguir);
$('#seguir').on('click',seguir);
$('#editar-usuario').on('submit',editar);
$('#editar-senha').on('submit',editarSenha);
$('#deletar-usuario').on('click',deletarConta);

function pararDeSeguir(){
    const usuarioID = $(this).data('usuario-id')
    $(this).prop('disabled', true);

    $.ajax({
        url: `/usuario/${usuarioID}/desseguir`,
        method: "POST"
    }).done(function(){
        window.location = `/usuario/${usuarioID}`
    }).fail(function(){
        swal("Ops, algo errado aconteceu",{
            icon: "error",
            text: "tente novamente mais tarde"
        })
        $(this).prop('disabled',false);
    })
}
function seguir(){
    const usuarioID = $(this).data('usuario-id')
    $(this).prop('disabled', true);

    $.ajax({
        url: `/usuario/${usuarioID}/seguir`,
        method: "POST"
    }).done(function(){
        window.location = `/usuario/${usuarioID}`
    }).fail(function(){
        swal("Ops, algo errado aconteceu",{
            icon: "error",
            text: "tente novamente mais tarde"
        })
        $(this).prop('disabled',false);
    })

}

function editar(evento){
        evento.preventDefault();
        const elementoClicado = $(evento.target);
        elementoClicado.prop('disabled',true);
    
        $.ajax({
            url:"/editar-usuario",
            method: "PUT",
            data:{
                nome:$('#nome').val(),
                email:$('#email').val(),
                nick:$('#nick').val(),
            }
        }).done(function(){
            swal("Usuário atualizado com sucesso",{
                icon:"success",
            }).then(function(){
                window.location = "/perfil";
            })
        }).fail(function(){
            swal("Problema ao atualizar o usuário",{
                icon:"error",
                text: "tente novamente mais tarde"
            })
        });
}

function editarSenha(evento){
    evento.preventDefault();
    if($('#senhaNova').val() != $('#senhaConfirmada').val()){
        swal("Ops, as senhas não coincidem",{
            icon:"error",
            text:"Cheque suas senhas e tente novamente."
        })
        return;
    }
    $.ajax({
        url:"/atualizar-senha",
        method: "POST",
        data: {
            "atual":$('#senhaAtual').val(),
            "nova":$('#senhaNova').val(),
        }
    }).done(function(){
        swal("Senha atualizada com sucesso!",{
            icon:"success"
        }).then(function(){
            window.location = "/perfil"
        });
    }).fail(function(){
        swal("Erro ao atualizar a senha!",{
            icon:"error"
        }).then(function(){
            window.location = "/atualizar-senha"
        });
    })
}

function deletarConta(evento){
    evento.preventDefault();
    console.log("Entrou")
    swal("Tem certeza que deseja excluir sua conta?",{
        title:"Tem certeza que deseja excluir sua conta?",
        dangerMode:true,
        icon:"warning",
        text:"Está ação é irreversível, tem certeza que deseja deletar PERMANENTEMENTE sua conta?",
        buttons:{
            cancel: "Cancelar",
            confirm:{ 
                text: "Confirmar",
                value: true,
            },
        }
    }).then((value) =>{
        if (value){
            const elementoClicado = $(evento.target);
            elementoClicado.prop('disabled',true)
            $.ajax({
                url: `/deletar-usuario`,
                method: "DELETE",
            }).done(function(){
                swal("Seu usuário foi excluído com sucesso, sentiremos saudade. =(",{
                    icon:"success"
                }).then(function(){
                    window.location = "/logout"
                });
            }).fail(function(){
                swal("Problema ao tentar deletar esta conta",{
                    icon:"error"
                })
            }).always(function(){
                elementoClicado.prop('disabled',false)
            });
        }
    });
    
}