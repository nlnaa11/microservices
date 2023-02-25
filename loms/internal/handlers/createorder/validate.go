package createorder

import (
	"strconv"

	"github.com/pkg/errors"
	internalErr "gitlab.ozon.dev/nlnaa/homework-1/libs/errors"
)

func (r Request) Validate() error {
	if len(r.Items) < 1 {
		return internalErr.ErrEmptyItems
	}

	for _, item := range r.Items {
		suffix := string(" item sku: ") + strconv.Itoa(int(item.Sku))
		if item.Sku == 0 {
			return errors.WithMessage(internalErr.ErrInvalidSku, suffix)
		}
		if item.Count == 0 {
			return errors.WithMessage(internalErr.ErrEmptyCount, suffix)
		}
	}

	if r.User < 1 {
		return internalErr.ErrInvalidUser
	}

	return nil
}
