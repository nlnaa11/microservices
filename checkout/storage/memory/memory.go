package memory

import (
	"github.com/pkg/errors"
	"gitlab.ozon.dev/nlnaa/homework-1/checkout/storage"
)

var (
	ErrAccessToStor = errors.New("memory access problems")
	ErrEmpryCart    = errors.New("empty cart")
)

type MemoryWrapper struct {
	carts map[int64]storage.Cart
}

func Init() (*MemoryWrapper, error) {
	return &MemoryWrapper{
		carts: make(map[int64]storage.Cart),
	}, nil
}

// TODO: везде добавить контексы + реализовать норм

func (m *MemoryWrapper) AddTo(item storage.Item, user int64, successChan chan bool) {
	if m.carts == nil && successChan != nil {
		successChan <- false
	}

	cart, ok := m.carts[user]
	if ok {
		cart.Items = append(cart.Items, item)
	} else {
		items := make([]storage.Item, 0, 1)
		items = append(items, item)
		m.carts[user] = storage.Cart{
			Items: items,
		}
	}

	if successChan != nil {
		successChan <- true
		return
	}
}

func (m *MemoryWrapper) RemoveFrom(itemForRemove storage.Item, user int64, successChan chan bool) {
	if m.carts == nil && successChan != nil {
		successChan <- false
	}

	cart, ok := m.carts[user]
	if ok {
		for i, item := range cart.Items {
			if item.Sku == itemForRemove.Sku {
				item.Count -= itemForRemove.Count

				if item.Count <= 0 {
					//cart.items = append(cart.items[0:i], cart.items[i+1:]...)
					size := len(cart.Items)
					cart.Items[i] = cart.Items[size-1]
					cart.Items = cart.Items[:size-1]
				}
				break
			}

		}
	}

	if successChan != nil {
		successChan <- true
		return
	}
}

func (m *MemoryWrapper) Get(user int64) (*storage.Cart, error) {
	if m.carts == nil {
		return nil, ErrAccessToStor
	}

	cart, ok := m.carts[user]
	if ok {
		return &cart, nil
	}

	return nil, ErrEmpryCart
}
