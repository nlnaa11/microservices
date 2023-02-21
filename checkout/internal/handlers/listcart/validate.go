package listcart

import (
	"strconv"

	"github.com/pkg/errors"
)

var (
	ErrEmptyCart      = errors.New("empty cart")
	ErrEmptyItemCount = errors.New("empty count")
	ErrEmptyItemName  = errors.New("empty name")
	ErrEmptyItemSku   = errors.New("empty sku")
	ErrEmptyPrice     = errors.New("empty price")
	ErrEmptyUser      = errors.New("empty user")
)

func (r Request) Validate() error {
	if r.User <= 0 {
		return ErrEmptyUser
	}
	return nil
}

func (r Response) Validate() error {
	if len(r.Items) < 1 {
		return ErrEmptyCart
	}
	items := r.Items

	for _, item := range items {
		suffix := string("item sku: ") + strconv.Itoa(int(item.Sku))
		if item.Sku == 0 {
			return errors.WithMessage(ErrEmptyItemSku, suffix)
		}
		if item.Count < 1 {
			return errors.WithMessage(ErrEmptyItemCount, suffix)
		}
		if len(item.Name) < 1 {
			return errors.WithMessage(ErrEmptyItemName, suffix)
		}
	}

	if r.TotalPrice < 0.09 {
		return ErrEmptyPrice
	}
	return nil
}
