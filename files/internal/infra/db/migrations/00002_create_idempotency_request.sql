-- +goose Up
-- +goose StatementBegin
-- Таблица идемпотентности HTTP/consumer'ов. Колонки должны совпадать с моделью
-- gocore/idempotency.IdempotencyRequest (id/cdate/udate + ключ/тело/срок).
CREATE TABLE data.idempotency_request (
    id              UUID PRIMARY KEY,
    cdate           TIMESTAMPTZ NOT NULL,
    udate           TIMESTAMPTZ NOT NULL,
    idempotency_key TEXT NOT NULL,
    response_data   JSONB,
    expires_at      TIMESTAMPTZ NOT NULL
);

CREATE UNIQUE INDEX ux_idempotency_request_key ON data.idempotency_request (idempotency_key);
CREATE INDEX ix_idempotency_request_expires_at ON data.idempotency_request (expires_at);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS data.idempotency_request;
-- +goose StatementEnd
