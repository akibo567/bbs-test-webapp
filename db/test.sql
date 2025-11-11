create table test_kakikomi (
    id SERIAL NOT NULL,
    name varchar,
    message varchar NOT NULL,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    PRIMARY KEY (id)
);