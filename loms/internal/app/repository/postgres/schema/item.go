package schema

type Item struct {
	Id    uint32 `db:"item_id"`
	Count uint64 `db:"count"`
}
