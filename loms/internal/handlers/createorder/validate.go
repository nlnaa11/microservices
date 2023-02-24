package createorder

import (
	"strconv"

	"github.com/pkg/errors"
)

var (
	ErrEmptyItems     = errors.New("empty items")
	ErrEmptyItemCount = errors.New("empty count")
	ErrEmptyItemSku   = errors.New("empty sku")
	ErrInvalidUser    = errors.New("invalid user")
)

func (r Request) Validate() error {
	if len(r.Items) < 1 {
		return ErrEmptyItems
	}

	for _, item := range r.Items {
		suffix := string("item sku: ") + strconv.Itoa(int(item.Sku))
		if item.Sku == 0 {
			return errors.WithMessage(ErrEmptyItemSku, suffix)
		}
		if item.Count < 1 {
			return errors.WithMessage(ErrEmptyItemCount, suffix)
		}
	}

	if r.User < 1 {
		return ErrInvalidUser
	}

	return nil
}
