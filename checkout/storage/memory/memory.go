package memory

type Item struct {
	sku   uint32
	count uint64
}

type Cart struct {
	items []Item
}

type MemoryWrapper struct {
	carts map[int64]*Cart
}

func Init() (*MemoryWrapper, error) {
	return &MemoryWrapper{
		carts: make(map[int64]*Cart),
	}, nil
}
