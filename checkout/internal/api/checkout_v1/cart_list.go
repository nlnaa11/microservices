package checkout_v1

import (
	"context"

	desc "gitlab.ozon.dev/nlnaa/homework-1/checkout/pkg/checkout_v1"
	productServiceApi "gitlab.ozon.dev/nlnaa/homework-1/checkout/pkg/product"
)

// TODO: сходить в productService.GetProductInfo, чтобы узнать цены и названия товаров
func (i *Implemantation) CartList(ctx context.Context, req *desc.CartListRequest) (*desc.CartListResponse, error) {
	// 1. получить корзину пользователя
	cartList, err := i.checkoutService.CartList(ctx, req.GetUser())
	if err != nil {
		return nil, err
	}

	// 2. получить полную информацию о каждом товаре из корзины
	// + рассчитать общую стоимость корзины
	var totalPrice float64
	itemsInfo := make([]*desc.ItemInfo, 0, len(cartList.ItemsInfo))

	for _, item := range cartList.ItemsInfo {
		getProductReq := productServiceApi.GetProductRequest{
			Token: i.tokenGetter.GetToken(),
			Sku:   item.Sku,
		}

		getProductRes, _ := i.productService.GetProduct(ctx, &getProductReq)
		if getProductRes == nil {
			continue
		}

		itemsInfo = append(itemsInfo, &desc.ItemInfo{
			Item: &desc.Item{
				Sku:   item.Sku,
				Count: item.Count,
			},
			Name:  getProductRes.GetName(),
			Price: float64(getProductRes.GetPrice()),
		})
		totalPrice += item.Price
	}

	return &desc.CartListResponse{
		ItemsInfo:  itemsInfo,
		TotalPrice: totalPrice,
	}, nil
}
