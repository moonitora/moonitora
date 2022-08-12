CREATE TABLE IF NOT EXISTS departamentos (
                                             id INTEGER PRIMARY KEY,
                                             title VARCHAR(64)
    );

CREATE TABLE IF NOT EXISTS usuarios (
    email VARCHAR(50) PRIMARY KEY,
    nome VARCHAR(100),
    ra VARCHAR(6),
    departamento INTEGER FOREIGN KEY,
    adm INTEGER,

    FOREIGN KEY (departamento) REFERENCES departamentos(id)
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
);

CREATE TABLE IF NOT EXISTS horarios (
    monitor VARCHAR(50),
    dia_da_semana INTEGER,
    hora INTEGER,
    minutos INTEGER,

    FOREIGN KEY (monitor) REFERENCES usuarios(email)
);


