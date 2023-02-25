package errors

import "errors"

var (
	ErrEmptyItems         = errors.New("empty items")
	ErrEmptyCart          = errors.New("empty cart")
	ErrEmptyCount         = errors.New("empty count")
	ErrEmptyListOfCart    = errors.New("empty list of cart")
	ErrEmptyName          = errors.New("empty name")
	ErrInsufficientStocks = errors.New("insufficient stocks")
	ErrInvalidOrder       = errors.New("invalid order")
	ErrInvalidOrderId     = errors.New("invalid order id")
	ErrInvalidPrice       = errors.New("invalid price")
	ErrInvalidSku         = errors.New("invalid sku")
	ErrInvalidUser        = errors.New("invalid user")
	ErrNoBusyCart         = errors.New("no busy cart")
	ErrOrderNotFound      = errors.New("order not found")
	ErrSkuNotFound        = errors.New("sku not found")
	ErrUnknownOrderStatus = errors.New("unknown order status")
)
