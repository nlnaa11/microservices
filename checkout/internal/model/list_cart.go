package model

import (
	"context"
	"log"

	"github.com/pkg/errors"
	internalErr "gitlab.ozon.dev/nlnaa/homework-1/libs/errors"
)

type Item struct {
	Sku   uint32 `json:"sku"`
	Count uint64 `json:"count"`
}

type ItemInfo struct {
	Sku   uint32  `json:"sku"`
	Count uint64  `json:"count"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

type Cart struct {
	Items []Item `json:"items"`
}

type ListOfCart struct {
	ItemsInfo  []ItemInfo `json:"itemsInfo"`
	TotalPrice float64    `json:"totalPrice"`
}

func (m *Model) ListCart(ctx context.Context, user int64) (*ListOfCart, error) {
	cart, err := m.stor.GetCart(ctx, user)
	if err != nil {
		return nil, errors.WithMessage(err, "getting cart")
	}
	if len(cart.Items) < 1 {
		return nil, internalErr.ErrEmptyCart
	}

	token := m.tokenGetter.GetToken()

	itemsInfo := make([]ItemInfo, 0, len(cart.Items))
	totalPrice := 0.0

	for _, item := range cart.Items {
		info, err := m.productInformator.GetProductInfo(ctx, token, item.Sku)
		if err != nil {
			log.Println(err)
			_ = m.stor.RemoveFromCart(ctx, user, item)
			continue
		}

		itemsInfo = append(itemsInfo, ItemInfo{
			Sku:   item.Sku,
			Count: item.Count,
			Name:  info.Name,
			Price: info.Price,
		})

		totalPrice += float64(item.Count) * info.Price
	}

	if len(itemsInfo) < 1 {
		return nil, internalErr.ErrEmptyCart
	}

	return &ListOfCart{
		ItemsInfo:  itemsInfo,
		TotalPrice: totalPrice,
	}, nil
}
