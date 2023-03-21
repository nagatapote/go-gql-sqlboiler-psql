
-- +migrate Up
alter table users add column job_id bigint references jobs(id);

-- +migrate Down
alter table users drop column job_id;