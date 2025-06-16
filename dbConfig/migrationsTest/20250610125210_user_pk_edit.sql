-- +goose Up
-- +goose StatementBegin
ALTER TABLE users RENAME COLUMN uid TO usid;

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE users RENAME COLUMN usid TO uid;
-- +goose StatementEnd
