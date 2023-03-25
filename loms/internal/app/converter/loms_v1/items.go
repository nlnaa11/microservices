package loms_v1

import (
	desc "gitlab.ozon.dev/nlnaa/homework-1/checkout/pkg/loms_v1"
	"gitlab.ozon.dev/nlnaa/homework-1/loms/internal/app/model"
)

func ToDescItems(items []model.Item) []*desc.Item {
	res := make([]*desc.Item, 0, len(items))
	for _, item := range items {
		res = append(res, ToDescItem(&item))
	}

	return res
}

func ToDescItem(item *model.Item) *desc.Item {
	return &desc.Item{
		Sku:   item.Sku,
		Count: item.Count,
	}
}

func ToItems(descItems []*desc.Item) []model.Item {
	res := make([]model.Item, 0, len(descItems))
	for _, descItem := range descItems {
		res = append(res, ToItem(descItem))
	}

	return res
}

func ToItem(descItem *desc.Item) model.Item {
	return model.Item{
		Sku:   descItem.GetSku(),
		Count: descItem.GetCount(),
	}
}
