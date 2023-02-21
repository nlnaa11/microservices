package purchase

import (
	"github.com/pkg/errors"
)

var (
	ErrEmptyUser = errors.New("empty user")
)

func (r Request) Validate() error {
	if r.User <= 0 {
		return ErrEmptyUser
	}
	return nil
}

func (r Response) Validate() error {
	if r.OrderId < 0 {
		return ErrInvalidOrder
	}

	return nil
}
