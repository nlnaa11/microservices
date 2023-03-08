package purchase

import (
	"gitlab.ozon.dev/nlnaa/homework-1/libs/errors"
)

func (r Request) Validate() error {
	if r.User == 0 {
		return errors.ErrInvalidUser
	}
	return nil
}

func (r Response) Validate() error {
	if r.OrderId < 0 {
		return errors.ErrInvalidOrderId
	}

	return nil
}
