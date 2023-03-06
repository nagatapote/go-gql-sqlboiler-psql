
-- +migrate Up
alter table users add column jobs_id bigint references jobs(id);

-- +migrate Down
alter table users drop column jobs_id;