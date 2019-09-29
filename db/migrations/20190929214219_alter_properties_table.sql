-- migrate:up
alter table properties add column square_feet float(1) not null;
alter table properties add column description text not null;
alter table properties add column status text not null;
alter table properties add column note text;


-- migrate:down

