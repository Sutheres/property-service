-- migrate:up
create table real_estate_types
(
    id          bigserial                                not null
        constraint real_estate_types_pk
        primary key,
    real_estate_type text                               not null,
    created_at timestamptz default current_timestamp    not null,
	updated_at timestamptz default current_timestamp
);

insert into real_estate_types(real_estate_type) values ('Residential');
insert into real_estate_types(real_estate_type) values ('Commercial');
insert into real_estate_types(real_estate_type) values ('Industrial');
insert into real_estate_types(real_estate_type) values ('Land');

-- migrate:down

