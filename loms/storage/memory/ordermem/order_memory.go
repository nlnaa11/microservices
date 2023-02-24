// Это просто черновик, который помогает мне понять,
// как перемещаются данные.
package ordermem

import (
	"context"

	"github.com/pkg/errors"
	"gitlab.ozon.dev/nlnaa/homework-1/loms/storage"
)

var (
	ErrOrderNotFound      = errors.New("order not found")
	ErrUnknownOrderStatus = errors.New("unknown order status")
)

const (
	FirstOrderId uint64 = 1

	StatusNew = 0
)

type Item struct {
	sku   uint32
	count uint16
}

type OrderData struct {
	status uint16
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

func (l *OrderMemoryWrapper) CreateOrder(ctx context.Context, user int64, storItems []storage.Item) (uint64, error) {
	orderId := l.nextOrderId
	l.nextOrderId++

	items := make([]Item, 0, len(storItems))
	for _, storItem := range storItems {
		items = append(items, Item{
			sku:   storItem.Sku,
			count: uint16(storItem.Count),
		})
	}

	order := OrderData{
		status: StatusNew,
		user:   user,
		items:  items,
	}

	// если что, данные полностью перезапишутся. что ok
	l.orders[orderId] = &order

	return orderId, nil
}

func (l *OrderMemoryWrapper) SetOrderStatus(ctx context.Context, orderId uint64, status uint16) error {
	order, ok := l.orders[orderId]
	if !ok {
		return ErrOrderNotFound
	}

	order.status = status

	return nil
}

func (l *OrderMemoryWrapper) GetOrderData(ctx context.Context, orderId uint64) (*storage.OrderData, error) {
	order, ok := l.orders[orderId]
	if !ok {
		return nil, ErrOrderNotFound
	}

	storItems := make([]storage.Item, 0, len(order.items))
	for _, item := range order.items {
		storItems = append(storItems, storage.Item{
			Sku:   item.sku,
			Count: uint64(item.count),
		})
	}

	storOrder := storage.OrderData{
		Status: order.status,
		User:   order.user,
		Items:  storItems,
	}

	return &storOrder, nil
}
