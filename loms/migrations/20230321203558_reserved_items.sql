-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS reserved_items (
    item_id bigint PRIMARY KEY,
    count bigint NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS reserved_items;
-- +goose StatementEnd
