$('#nova-publicacao').on('submit',criarPublicacao);
$('#editar-publicacao').on('submit',editarPublicacao);
$(document).on('click','.curtir-publicacao',curtirPublicacao);
$(document).on('click','.descurtir-publicacao',descurtirPublicacao);
$(document).on('click','.deletar-publicacao',deletarPublicacao);
function criarPublicacao(evento){
    evento.preventDefault();
    $.ajax({
        url: "/publicacoes",
        method: "POST",
        data:{
            titulo: $('#titulo').val(),
            conteudo: $('#conteudo').val(),
        }
    }).done(function(){
       window.location = "/home"
    }).fail(function(){
        swal("Erro ao criar a publicação",{
            icon:"error",
        })
    })
}
function editarPublicacao(evento){
    evento.preventDefault();
    const button = $('#btn-salvar')
    $(this).prop('disabled',true)
    const idPublicacao = button.data('publicacao-id')
    $.ajax({
        url: `/publicacoes/${idPublicacao}/editar`,
        method: "POST",
        data:{
            titulo: $('#titulo').val(),
            conteudo: $('#conteudo').val(),
        }
    }).done(function(){
        swal("Feito!","Publicação editada com sucesso!","success")
        .then(
            window.location = "/home"
        )
    }).fail(function(){
        swal("Falha","Publicação não pôde ser editada!","fail")
    }).always(function(){
        button.prop('disabled',false)
    })
}
function curtirPublicacao(evento){
    evento.preventDefault();
    const elementoClicado = $(evento.target);
    const publicacaoID = elementoClicado.closest('div').data('publicacao-id');
    elementoClicado.prop('disabled',true)

    elementoClicado.addClass("descurtir-publicacao")
    elementoClicado.removeClass("curtir-publicacao")
    elementoClicado.css('color','#e00a97')
    $.ajax({
        url: `/publicacoes/${publicacaoID}/curtir`,
        method: "POST",
    }).done(function(){
       const contadorDeCurtidas = elementoClicado.next('span');
       const quantidadeDeCurtidas = parseInt(contadorDeCurtidas.text());
       
       contadorDeCurtidas.text(quantidadeDeCurtidas + 1);
    }).fail(function(){
        elementoClicado.addClass("curtir-publicacao")
        elementoClicado.removeClass("descurtir-publicacao")
        elementoClicado.css('color','#A9A9A9')
    }).always(function(){
        elementoClicado.prop('disabled',false)
    });
}


function descurtirPublicacao(evento){
    evento.preventDefault();
    const elementoClicado = $(evento.target);
    const publicacaoID = elementoClicado.closest('div').data('publicacao-id');
    elementoClicado.prop('disabled',true)

    elementoClicado.addClass("curtir-publicacao")
    elementoClicado.removeClass("descurtir-publicacao")
    elementoClicado.css('color','#A9A9A9')
    $.ajax({
        url: `/publicacoes/${publicacaoID}/descurtir`,
        method: "POST",
    }).done(function(){
       const contadorDeCurtidas = elementoClicado.next('span');
       const quantidadeDeCurtidas = parseInt(contadorDeCurtidas.text());
       
       contadorDeCurtidas.text(quantidadeDeCurtidas - 1);
       
    }).fail(function(){
        elementoClicado.addClass("descurtir-publicacao")
        elementoClicado.removeClass("curtir-publicacao")
        elementoClicado.css('color','#A9A9A9')
    }).always(function(){
        elementoClicado.prop('disabled',false)
    });
}
async function deletarPublicacao(evento){
    evento.preventDefault();

    swal("Tem certeza que deseja excluir essa publicação? Essa ação é irreversível.",{
        icon: "warning",
        dangerMode: true,
        title: "Atenção!",
        buttons: {
            cancel: "Cancelar",
            confirm:{ 
                text: "Confirmar",
                value: true,
            },
        },
    }).then((value) =>{
            if (value){
                const elementoClicado = $(evento.target);
                const publicacao = elementoClicado.closest('div');
                const publicacaoID = publicacao.data('publicacao-id');
                elementoClicado.prop('disabled',true)
                $.ajax({
                    url: `/publicacoes/${publicacaoID}`,
                    method: "DELETE",
                }).done(function(){
                    publicacao.fadeOut("slow", function(){
                        $(this).remove();
                    });
                }).fail(function(){
                    console.log("publicação não pode ser deletada")
                }).always(function(){
                    elementoClicado.prop('disabled',false)
                });
            }
        });

    
}