package loms_v1

import (
	desc "gitlab.ozon.dev/nlnaa/homework-1/checkout/pkg/loms_v1"
	"gitlab.ozon.dev/nlnaa/homework-1/loms/internal/model"
)

func ToDescOrderInfo(orderInfo *model.OrderInfo) *desc.OrderInfo {
	res := desc.OrderInfo{
		Order: ToDescOrderItem(&orderInfo.Order),
		User:  orderInfo.User,
		Items: ToDescItems(orderInfo.Items),
	}

	return &res
}

func ToDescOrderItem(order *model.Order) *desc.OrderItem {
	return &desc.OrderItem{
		OrderId: order.OrderId,
		Status:  order.Status,
	}
}
