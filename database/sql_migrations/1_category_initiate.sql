-- +migrate Up
-- +migrate StatementBegin

create table category (
    id serial not null primary key,
    name varchar(256),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
)

-- +migrate StatementEnd