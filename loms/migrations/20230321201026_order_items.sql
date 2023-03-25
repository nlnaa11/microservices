-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS orders_items (
    order_id bigint NOT NULL,
    item_id bigint NOT NULL,
    count int NOT NULL,
    PRIMARY KEY (order_id, item_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS orders_items;
-- +goose StatementEnd
