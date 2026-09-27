-- +goose Up
-- +goose StatementBegin
CREATE SCHEMA IF NOT EXISTS data;

CREATE TABLE data.examples (
    id            UUID PRIMARY KEY,
    cdate         TIMESTAMPTZ NOT NULL,
    udate         TIMESTAMPTZ NOT NULL,
    name          TEXT NOT NULL,
    email         TEXT NOT NULL,
    status        TEXT NOT NULL,
    retry_count   INTEGER NOT NULL DEFAULT 0,
    next_retry_at TIMESTAMPTZ NOT NULL
);

CREATE UNIQUE INDEX ux_examples_email ON data.examples (email);
CREATE INDEX ix_examples_status_next_retry ON data.examples (status, next_retry_at);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS data.examples;
-- +goose StatementEnd
