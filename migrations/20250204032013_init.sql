-- +goose Up
-- +goose StatementBegin
create table accounts (
    user_id uuid primary key default gen_random_uuid(),
    balance int not null default 0
);

create table operations (
    id uuid primary key default gen_random_uuid(),
    source_user_id uuid not null,
    destination_user_id uuid,
    type integer,
    amount int not null,
    request_time timestamptz default now()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table operations;
drop table accounts;
-- +goose StatementEnd
