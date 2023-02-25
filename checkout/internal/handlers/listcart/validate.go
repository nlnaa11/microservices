package listcart

import (
	"strconv"

	"github.com/pkg/errors"
	internalErr "gitlab.ozon.dev/nlnaa/homework-1/libs/errors"
)

const minimumCost = 0.0

func (r Request) Validate() error {
	if r.User == 0 {
		return internalErr.ErrInvalidUser
	}
	return nil
}

func (r Response) Validate() error {
	if len(r.Items) < 1 {
		return internalErr.ErrEmptyCart
	}

	for _, item := range r.Items {
		suffix := string(" item sku: ") + strconv.Itoa(int(item.Sku))
		if item.Sku == 0 {
			return errors.WithMessage(internalErr.ErrInvalidSku, suffix)
		}
		if item.Count < 1 {
			return errors.WithMessage(internalErr.ErrEmptyCount, suffix)
		}
		if len(item.Name) < 1 {
			return errors.WithMessage(internalErr.ErrEmptyName, suffix)
		}
	}

	if r.TotalPrice <= minimumCost {
		return internalErr.ErrInvalidPrice
	}
	return nil
}
