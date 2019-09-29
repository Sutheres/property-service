-- migrate:up
create table images(
    id bigserial not null
        constraint images_pk
        primary key,
    title text not null,
    url text not null,
    created_at timestamptz default current_timestamp not null,
	updated_at timestamptz default current_timestamp
)


-- migrate:down

