package deletefromcart

import (
	"gitlab.ozon.dev/nlnaa/homework-1/libs/errors"
)

func (r Request) Validate() error {
	if r.User == 0 {
		return errors.ErrInvalidUser
	}
	if r.Sku == 0 {
		return errors.ErrInvalidSku
	}
	if r.Count < 1 {
		return errors.ErrEmptyCount
	}
	return nil
}
