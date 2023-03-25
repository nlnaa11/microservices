-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS items_stocks (
    warehouse_id bigint NOT NULL,
    item_id bigint NOT NULL,
    count bigint NOT NULL,
    PRIMARY KEY (warehouse_id, item_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS items_stocks;
-- +goose StatementEnd
