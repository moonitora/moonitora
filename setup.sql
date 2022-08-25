CREATE TABLE IF NOT EXISTS departamentos (
                                             id VARCHAR(12) PRIMARY KEY,
                                             title VARCHAR(64)
    );

CREATE TABLE IF NOT EXISTS usuarios (
    email VARCHAR(50) PRIMARY KEY,
    nome VARCHAR(100),
    ra VARCHAR(6),
    departamento VARCHAR(12),
    adm INTEGER,

    FOREIGN KEY (departamento) REFERENCES departamentos(id)
);

CREATE TABLE IF NOT EXISTS monitorias (
    id VARCHAR(12) PRIMARY KEY,
    marcada_por VARCHAR(50),
    monitor VARCHAR(50),
    departamento VARCHAR(12),
    conteudo VARCHAR(128),
    disciplina VARCHAR(128),
    horario VARCHAR(24),
    aluno_nome VARCHAR(128),
    aluno_ra VARCHAR(6),
    data VARCHAR(16),
    status INTEGER,

    FOREIGN KEY (monitor) REFERENCES usuarios(email),
    FOREIGN KEY (departamento) REFERENCES departamentos(id),
    FOREIGN KEY (horario) REFERENCES horarios(id)
);

CREATE TABLE IF NOT EXISTS login (
    email VARCHAR(50) PRIMARY KEY,
    password VARCHAR(512),

    FOREIGN KEY (email) REFERENCES usuarios(email)
);

CREATE TABLE IF NOT EXISTS horarios (
    id VARCHAR(24) PRIMARY KEY,
    monitor VARCHAR(50),
    dia_da_semana INTEGER,
    inicio_horas INTEGER,
    inicio_minutos INTEGER,
    termino_horas INTEGER,
    termino_minutos INTEGER,

    FOREIGN KEY (monitor) REFERENCES usuarios(email)
);


