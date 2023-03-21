
-- +migrate Up
create table if not exists jobs(
  id bigserial not null,
  name text not null,
  created_at timestamp with time zone not null default now(),
  updated_at timestamp with time zone not null default now(),
  deleted_at timestamp with time zone,
  unique (name),
  primary key (id)
);

-- +migrate Down
drop table if exists jobs;