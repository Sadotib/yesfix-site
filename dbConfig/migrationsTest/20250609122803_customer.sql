-- +goose Up
-- +goose StatementBegin
CREATE TABLE customers (
    cid         VARCHAR(16) NOT NULL PRIMARY KEY,           -- e.g., 'DBG000123'
    firstname   TEXT NOT NULL,
    midname     TEXT,
    lastname    TEXT NOT NULL,
    email       TEXT NOT NULL UNIQUE,
    phone       VARCHAR(15) NOT NULL,
    city        TEXT NOT NULL,
    locality    TEXT NOT NULL,
    pincode     VARCHAR(10) NOT NULL,
    created_at  TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at  TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at  TIMESTAMPTZ
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS customers;
-- +goose StatementEnd