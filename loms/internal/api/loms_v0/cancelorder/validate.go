package cancelorder

import (
	"gitlab.ozon.dev/nlnaa/homework-1/libs/errors"
)

func (r Request) Validate() error {
	if r.OrderId < 1 {
		return errors.ErrInvalidOrderId
	}

	return nil
}
