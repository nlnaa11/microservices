package deletefromcart

import "errors"

var (
	ErrEmptyCount = errors.New("empty count")
	ErrEmptyUser  = errors.New("empty user")
	ErrEmptySKU   = errors.New("empty sku")
)

func (r Request) Validate() error {
	if r.User <= 0 {
		return ErrEmptyUser
	}
	if r.Sku == 0 {
		return ErrEmptySKU
	}
	if r.Count < 1 {
		return ErrEmptyCount
	}
	return nil
}
