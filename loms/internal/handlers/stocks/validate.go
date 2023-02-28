package stocks

import (
	"gitlab.ozon.dev/nlnaa/homework-1/libs/errors"
)

func (r Request) Validate() error {
	if r.Sku < 1 {
		return errors.ErrInvalidSku
	}
	return nil
}
