CREATE TABLE IF NOT EXISTS usuarios (
    email VARCHAR(50) PRIMARY KEY,
    nome VARCHAR(100),
    ra VARCHAR(6),
    curso INTEGER
);

CREATE TABLE IF NOT EXISTS monitorias (
    monitor VARCHAR(50),
    conteudo VARCHAR(512),
    dia INTEGER,
    mes INTEGER,
    ano INTEGER,
    ra_aluno VARCHAR(6),

    FOREIGN KEY (monitor) REFERENCES usuarios(email)
);

CREATE TABLE IF NOT EXISTS login (
    email VARCHAR(50) PRIMARY KEY,
    password VARCHAR(512),

    FOREIGN KEY (email) REFERENCES usuarios(email)
)


