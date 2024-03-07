-- +migrate Up
-- +migrate StatementBegin

CREATE TYPE gender AS ENUM ('laki-laki', 'perempuan');

CREATE SEQUENCE IF NOT EXISTS user_pkey_seq;
CREATE TABLE "user"
(
    id         BIGINT       NOT NULL DEFAULT nextval('user_pkey_seq'::regclass),
    username   VARCHAR(20)  NOT NULL,
    password   VARCHAR(255) NOT NULL,
    created_at TIMESTAMP WITHOUT TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITHOUT TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP WITHOUT TIME ZONE DEFAULT NULL,
    deleted    BOOLEAN               DEFAULT FALSE,
    CONSTRAINT pk_user_id PRIMARY KEY (id),
    CONSTRAINT uq_user_name UNIQUE (username)
);

INSERT INTO "user"(username, password)
VALUES ('System', 'system');

CREATE SEQUENCE IF NOT EXISTS employee_pkey_seq;
CREATE TABLE "employee"
(
    id           BIGINT       NOT NULL DEFAULT nextval('employee_pkey_seq'::regclass),
    name         VARCHAR(255) NOT NULL,
    nip          VARCHAR(255) NOT NULL,
    birthplace   VARCHAR(255),
    birthdate    TIMESTAMP WITHOUT TIME ZONE,
    age          INTEGER,
    address      VARCHAR(255),
    religion     VARCHAR(20),
    gender       gender       NOT NULL,
    phone_number VARCHAR(20),
    email        VARCHAR(255) NOT NULL,
    created_by   BIGINT,
    created_at   TIMESTAMP WITHOUT TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at   TIMESTAMP WITHOUT TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_by   BIGINT,
    deleted_at   TIMESTAMP WITHOUT TIME ZONE DEFAULT NULL,
    deleted      BOOLEAN               DEFAULT FALSE,
    CONSTRAINT pk_employee_id PRIMARY KEY (id)
);

-- +migrate StatementEnd