package schema

type Order struct {
	Id     uint64 `db:"order_id"`
	Status Status `db:"status"`
}

type OrderData struct {
	Order  Order  `db:"order"`
	UserId int64  `db:"user_id"`
	Items  []Item `db:"items"`
}
