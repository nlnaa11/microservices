package orderpayed

import "errors"

var (
	ErrInvalidOrderId = errors.New("invalid order id")
)

func (r Request) Validate() error {
	if r.OrderId < 1 {
		return ErrInvalidOrderId
	}

	return nil
}
