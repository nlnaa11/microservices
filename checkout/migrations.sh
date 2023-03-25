export MIGRATION_DIR=./migrations
export DB_HOST=localhost
export DB_PORT=5432
export DB_NAME=checkout
export DB_USER=user
export DB_PASSWORD=password
export DB_SSL=disable

export PG_DSN="host=${DB_HOST} port=${DB_PORT} dbname=${DB_NAME} user=${DB_USER} password=${DB_PASSWORD} sslmode=${DB_SSL}"

goose -dir ${MIGRATION_DIR} postgres "${PG_DSN}" up
goose -dir ${MIGRATION_DIR} postgres "${PG_DSN}" status