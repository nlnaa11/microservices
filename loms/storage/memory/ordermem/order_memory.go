package ordermem

const (
	FirstOrderId uint64 = 1
)

type Item struct {
	sku   uint32
	count uint64
}

type OrderData struct {
	status Status
	user   int64
	items  []Item
}

type OrderMemoryWrapper struct {
	// key: orderId
	orders      map[uint64]*OrderData
	nextOrderId uint64
}

func Init() (*OrderMemoryWrapper, error) {
	return &OrderMemoryWrapper{
		orders:      make(map[uint64]*OrderData),
		nextOrderId: FirstOrderId,
	}, nil
}
