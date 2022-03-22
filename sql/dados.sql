insert into users(name, nick, email, password)
values
("Usuário 1", "usuario_1", "usuario1@gmail.com", "$2a$10$1ZqlVOpd7hPwu/xCuEEMr.Yr8XqkioX5emtUiG2xS1KIuTINGz4eG"), -- usuario1
("Usuário 2", "usuario_2", "usuario2@gmail.com", "$2a$10$1ZqlVOpd7hPwu/xCuEEMr.Yr8XqkioX5emtUiG2xS1KIuTINGz4eG"), -- usuario2
("Usuário 3", "usuario_3", "usuario3@gmail.com", "$2a$10$1ZqlVOpd7hPwu/xCuEEMr.Yr8XqkioX5emtUiG2xS1KIuTINGz4eG"); -- usuario3

insert into followers(user_id, follower_id)
values
(1, 2),
(3, 1),
(1, 3);