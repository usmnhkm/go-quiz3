-- +migrate Up
-- +migrate StatementBegin
create table Book (
    id serial not null primary key,
    title varchar(256) not null,
    description text not null,
    image_url varchar(256) not null,
    release_year int not null,
    price varchar(256) not null,
    total_page int not null,
    thickness varchar(256) not null,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    category_id int not null
)
-- +migrate StatementEnd