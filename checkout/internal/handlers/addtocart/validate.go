package addtocart

import "errors"

var (
	ErrEmptyCount = errors.New("empty count")
	ErrEmptyUser  = errors.New("empty user")
	ErrInvalidSku = errors.New("invalid sku")
)

func (r Request) Validate() error {
	if r.User <= 0 {
		return ErrEmptyUser
	}
	if r.Sku < 1 {
		return ErrInvalidSku
	}
	if r.Count < 1 {
		return ErrEmptyCount
	}
	return nil
}
