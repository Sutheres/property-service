-- migrate:up
create table property_types
(
    id bigserial not null constraint property_types_pk primary key,
    property_type text not null,
    created_at timestamptz default current_timestamp not null,
	updated_at timestamptz default current_timestamp
);

insert into property_types(property_type) values ('Single-family');
insert into property_types(property_type) values ('Multi-family');
insert into property_types(property_type) values ('Co-op');
insert into property_types(property_type) values ('Condo');
insert into property_types(property_type) values ('Townhouse');
insert into property_types(property_type) values ('Duplex');
insert into property_types(property_type) values ('Triple decker');
insert into property_types(property_type) values ('4-plex');
insert into property_types(property_type) values ('High value home');
insert into property_types(property_type) values ('Generational');
insert into property_types(property_type) values ('Vacation home');
insert into property_types(property_type) values ('Shopping center');
insert into property_types(property_type) values ('Strip mall');
insert into property_types(property_type) values ('Medical building');
insert into property_types(property_type) values ('Educational building');
insert into property_types(property_type) values ('Hotel');
insert into property_types(property_type) values ('Office building');
insert into property_types(property_type) values ('Apartment building');
insert into property_types(property_type) values ('Manufacturing building');
insert into property_types(property_type) values ('Warehouse');
insert into property_types(property_type) values ('Other');
insert into property_types(property_type) values ('Vacant land');
insert into property_types(property_type) values ('Working farm');
insert into property_types(property_type) values ('Ranch');

-- migrate:down