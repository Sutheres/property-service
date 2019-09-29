-- migrate:up
alter table images add column property_id bigserial not null;
alter table images add constraint properties_fkey foreign key (property_id) references properties (id);

-- migrate:down

