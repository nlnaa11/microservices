-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS users_carts (
    user_id bigint NOT NULL,
    item_id bigint NOT NULL,
    count int NOT NULL,
    PRIMARY KEY (user_id, item_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS users_carts;
-- +goose StatementEnd
