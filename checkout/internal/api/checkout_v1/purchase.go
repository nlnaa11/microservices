package checkout_v1

import (
	"context"

	"github.com/pkg/errors"
	desc "gitlab.ozon.dev/nlnaa/homework-1/checkout/pkg/checkout_v1"
	lomsServiceApi "gitlab.ozon.dev/nlnaa/homework-1/checkout/pkg/loms_v1"
)

func (i *Implemantation) Purchase(ctx context.Context, req *desc.PurchaseRequest) (*desc.PurchaseResponse, error) {
	// 1. Получить корзину покупателя
	cart, err := i.checkoutService.CartList(ctx, req.GetUser())
	if err != nil {
		return nil, errors.WithMessage(err, "getting user cart")
	}

	// 2. создать заказ в сервисе loms
	lomsItems := make([]*lomsServiceApi.Item, 0, len(cart.ItemsInfo))
	for _, item := range cart.ItemsInfo {
		lomsItems = append(lomsItems, &lomsServiceApi.Item{
			Sku:   item.Sku,
			Count: item.Count,
		})
	}
	createOrderReq := lomsServiceApi.CreateOrderRequest{
		User:  req.GetUser(),
		Items: lomsItems,
	}

	createOrderRes, err := i.lomsClient.CreateOrder(ctx, &createOrderReq)
	if err != nil {
		return nil, errors.WithMessage(err, "creating order")
	}

	// TODO: по идее, здесь еще должна быть часть с покупкой (провл, отмена, оплата прошла)
	// или не здесь. Ждемс

	orderInfo := desc.PurchaseResponse{
		OrderId: createOrderRes.GetOrderId(),
		Status:  createOrderRes.GetStatus(),
	}

	return &orderInfo, nil
}
