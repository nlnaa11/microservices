package logisticsmem

const FirstWarehouseId int64 = 1

type Warehouse struct {
	// key: sku, value: count
	items map[uint32]uint64
}

type LogisticstMemoryWrapper struct {
	// key: warehouseId
	warehouses map[int64]*Warehouse
	// key: sku, value: count
	reserveItems    map[uint32]uint64
	nextWarehouseId int64
}

func Init() (*LogisticstMemoryWrapper, error) {
	return &LogisticstMemoryWrapper{
		warehouses:      make(map[int64]*Warehouse),
		reserveItems:    make(map[uint32]uint64),
		nextWarehouseId: FirstWarehouseId,
	}, nil
}
