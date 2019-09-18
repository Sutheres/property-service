-- migrate:up
create table properties
(
    id bigserial not null
        constraint properties_pk
        primary key,
    property_id text not null,
    street_number text not null,
    street text not null,
    street_suffix text not null,
    city text not null,
    state text not null,
    zip text not null,
    bedrooms float(1),
    bathrooms float(1),
    price text not null,
    real_estate_type text not null,
    property_type text not null,
    created_at timestamptz default current_timestamp not null,
	updated_at timestamptz default current_timestamp

);
create unique index properties_real_estate_type_uindex
    on properties (real_estate_type);

create unique index properties_property_type_uindex
    on properties (property_type);

-- migrate:down

