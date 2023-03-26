package errors

import "errors"

var (
	ErrEmptyItems      = errors.New("empty items")
	ErrEmptyCart       = errors.New("empty cart")
	ErrEmptyCount      = errors.New("empty count")
	ErrEmptyListOfCart = errors.New("empty list of cart")
	ErrEmptyName       = errors.New("empty name")
	ErrEmptyStocks     = errors.New("empty stocks")

	ErrInsufficientStocks = errors.New("insufficient stocks")
	ErrOrderRegistation   = errors.New("order registration")

	ErrInvalidOrder   = errors.New("invalid order")
	ErrInvalidOrderId = errors.New("invalid order id")
	ErrInvalidPrice   = errors.New("invalid price")
	ErrInvalidSku     = errors.New("invalid sku")
	ErrInvalidUser    = errors.New("invalid user")

	ErrNoAddedRows            = errors.New("no added rows")
	ErrNoDeletedRows          = errors.New("no deleted rows")
	ErrNoUpdatedRows          = errors.New("no updated rows")
	ErrNoUpdatedOrAddedRows   = errors.New("no updated or added rows")
	ErrNoUpdatedOrDeletedRows = errors.New("no updated or deleted rows")

	ErrNoBusyCart    = errors.New("no busy cart")
	ErrItemNotFound  = errors.New("item not found")
	ErrOrderNotFound = errors.New("order not found")
	ErrSkuNotFound   = errors.New("sku not found")

	ErrUnknownOrderStatus = errors.New("unknown order status")
)

// for workerPool
var (
	ErrNoWorkers          = errors.New("no workers")
	ErrInvalidTasksBuffer = errors.New("invalid tasks buffer")

	ErrInvalidTaskFunction = errors.New("invalid task function")
)
