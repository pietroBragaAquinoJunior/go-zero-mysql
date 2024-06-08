USE go_zero_mysql_db;

CREATE TABLE IF NOT EXISTS bloco_notas (
    id CHAR(36) NOT NULL,
    titulo VARCHAR(255) NOT NULL,
    descricao text NOT NULL,
    PRIMARY KEY (id)
);

INSERT INTO bloco_notas (id, titulo, descricao) VALUES ('98dab354-a9b0-4b2b-9b3b-2a7c311882ab', 'minhas receitas', 'Chocolate com pimenta, leite com manga e miojo com frango.');