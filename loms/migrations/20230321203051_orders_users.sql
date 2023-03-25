-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS orders_users (
    order_id bigint PRIMARY KEY,
    users_id bigint NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS orders_users;
-- +goose StatementEnd
