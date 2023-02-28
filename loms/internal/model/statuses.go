package model

type Status uint8

const (
	StatusNew Status = iota
	StatusFailed
	StatusAwaitingPayment
	StatusPayed
	StatusCancelled
	StatusUnknown
)

const (
	awaitingPayment = "awaiting payment"
	cancelled       = "cancelled"
	failed          = "failed"
	newOrder        = "new"
	payed           = "payed"
	unknown         = "unknown"
)

func StatusFromString(status string) Status {
	switch status {
	case awaitingPayment:
		return StatusAwaitingPayment
	case cancelled:
		return StatusCancelled
	case failed:
		return StatusFailed
	case newOrder:
		return StatusNew
	case payed:
		return StatusPayed
	default:
		return StatusUnknown
	}
}

func (s Status) String() string {
	switch s {
	case StatusNew:
		return newOrder
	case StatusFailed:
		return failed
	case StatusAwaitingPayment:
		return awaitingPayment
	case StatusPayed:
		return payed
	case StatusCancelled:
		return cancelled
	default:
		return unknown
	}
}
