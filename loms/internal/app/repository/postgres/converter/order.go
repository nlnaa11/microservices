package converter

import (
	"gitlab.ozon.dev/nlnaa/homework-1/loms/internal/app/model"
	"gitlab.ozon.dev/nlnaa/homework-1/loms/internal/app/repository/postgres/schema"
)

func ToModelOrderInfo(items []schema.Item, order *schema.Order, user int64) *model.OrderInfo {
	var result model.OrderInfo
	result.Items = ToModelItems(items)
	result.Order = *ToModelOrder(order)
	result.User = user

	return &result
}

func ToModelOrder(order *schema.Order) *model.Order {
	return &model.Order{
		OrderId: order.Id,
		Status:  order.Status.String(),
	}
}
