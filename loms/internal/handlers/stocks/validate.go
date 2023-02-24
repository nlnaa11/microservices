package stocks

import "errors"

var (
	ErrInvalidSku = errors.New("invalid sku")
)

func (r Request) Validate() error {
	if r.Sku < 1 {
		return ErrInvalidSku
	}
	return nil
}
