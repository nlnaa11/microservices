package schema

type Stock struct {
	WarehouseId int64  `db:"warehouse_id"`
	Count       uint64 `db:"count"`
}
