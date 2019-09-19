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

ALTER TABLE properties
   ADD CONSTRAINT real_estate_types
   CHECK (real_estate_type = 'Residential' OR
   real_estate_type = 'Commercial' OR
   real_estate_type = 'Industrial' OR
   real_estate_type = 'Land');

ALTER TABLE properties
   ADD CONSTRAINT property_types
   CHECK (property_type = 'Single-Family' OR
   property_type = 'Multi-family' OR
   property_type = 'Condo' OR
   property_type = 'Co-op' OR
   property_type = 'Townhouse' OR
   property_type = 'Duplex' OR
   property_type = 'Triple decker' OR
   property_type = '4-plex' OR
   property_type = 'High value home' OR
   property_type = 'Generational' OR
   property_type = 'Vacation home' OR
   property_type = 'Shopping center' OR
   property_type = 'Strip mall' OR
   property_type = 'Medical building' OR
   property_type = 'Educational building' OR
   property_type = 'Hotel' OR
   property_type = 'Office building' OR
   property_type = 'Apartment building' OR
   property_type = 'Manufacturing building' OR
   property_type = 'Warehouse' OR
   property_type = 'Other' OR
   property_type = 'Vacant land' OR
   property_type = 'Working farm' OR
   property_type = 'Ranch');

-- migrate:down

