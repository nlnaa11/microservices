package checkout_v1

import (
	"context"

	"github.com/pkg/errors"
	desc "gitlab.ozon.dev/nlnaa/homework-1/checkout/pkg/checkout_v1"
	lomsServiceApi "gitlab.ozon.dev/nlnaa/homework-1/checkout/pkg/loms_v1"
	libErr "gitlab.ozon.dev/nlnaa/homework-1/libs/errors"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (i *Implemantation) AddToCart(ctx context.Context, req *desc.AddToCartRequest) (*emptypb.Empty, error) {
	// 1. проверить, есть ли товар с таким артикулoм на складах
	stocksReq := lomsServiceApi.StocksRequest{
		Sku: req.GetItem().GetSku(),
	}
	lomsStocksRes, err := i.lomsClient.Stocks(ctx, &stocksReq)
	if err != nil {
		return &emptypb.Empty{}, errors.WithMessage(err, "getting stocks")
	}

	// 2. проверить наличие достаточного количества
	sufficient := false
	cnt := int64(req.GetItem().Count)
	for _, stock := range lomsStocksRes.GetStocks() {
		cnt -= int64(stock.GetCount())
		if cnt <= 0 {
			sufficient = true
			break
		}
	}
	if !sufficient {
		return &emptypb.Empty{}, libErr.ErrInsufficientStocks
	}

	// 3. добавить в корзину
	err = i.checkoutService.AddToCart(ctx, req.GetUser(), req.GetItem().GetSku(), req.GetItem().GetCount())
	if err != nil {
		return &emptypb.Empty{}, err
	}

	return &emptypb.Empty{}, nil
}
