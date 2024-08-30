INSERT INTO usuarios (nome, nick, email, senha)
VALUES ("Usuário_1", "User 1", "EmailUsuario1@exemplo.com", "$2a$10$MpDPzzgFDIR79P2jkrujceP.uAkhMuozxjNU8AjdoIX6VTm2W4hAa");
INSERT INTO usuarios (nome, nick, email, senha)
VALUES ("Usuário_2", "User 2", "EmailUsuario2@exemplo.com", "$2a$10$MpDPzzgFDIR79P2jkrujceP.uAkhMuozxjNU8AjdoIX6VTm2W4hAa");
INSERT INTO usuarios (nome, nick, email, senha)
VALUES ("Usuário_3", "User 3", "EmailUsuario3@exemplo.com", "$2a$10$MpDPzzgFDIR79P2jkrujceP.uAkhMuozxjNU8AjdoIX6VTm2W4hAa");
INSERT INTO usuarios (nome, nick, email, senha)
VALUES ("Usuário_4", "User 4", "EmailUsuario4@exemplo.com", "$2a$10$MpDPzzgFDIR79P2jkrujceP.uAkhMuozxjNU8AjdoIX6VTm2W4hAa");

INSERT INTO seguidores (usuario_id, seguidor_id) 
VALUES (1,2);
INSERT INTO seguidores (usuario_id, seguidor_id) 
VALUES (1,3);
INSERT INTO seguidores (usuario_id, seguidor_id)
VALUES (3,1);
INSERT INTO seguidores (usuario_id, seguidor_id)
VALUES (2,3);
INSERT INTO seguidores (usuario_id, seguidor_id)
VALUES (3,2);

INSERT INTO publicacoes(titulo, conteudo, autor_id) VALUES ("Publicação do usuário1", "Conteúdo", 1);
INSERT INTO publicacoes(titulo, conteudo, autor_id) VALUES ("Publicação do usuário2", "Conteúdo", 2);
INSERT INTO publicacoes(titulo, conteudo, autor_id) VALUES ("Publicação do usuário3", "Conteúdo", 3);
INSERT INTO publicacoes(titulo, conteudo, autor_id) VALUES ("Publicação do usuário4", "Conteúdo", 4);