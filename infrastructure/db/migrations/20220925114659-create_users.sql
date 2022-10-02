-- +migrate Up
create table if not exists users(
	id bigserial not null,
	created_at timestamp with time zone not null default now(),
  updated_at timestamp with time zone not null default now(),
  deleted_at timestamp with time zone,
  name text not null,
  email text not null,
  unique (email),
  primary key (id)
);
-- +migrate Down
drop table if exists users;