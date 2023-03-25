package schema

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
	AwaitingPayment = "awaiting payment"
	Cancelled       = "cancelled"
	Failed          = "failed"
	NewOrder        = "new"
	Payed           = "payed"
	Unknown         = "unknown"
)

func StatusFromString(status string) Status {
	switch status {
	case AwaitingPayment:
		return StatusAwaitingPayment
	case Cancelled:
		return StatusCancelled
	case Failed:
		return StatusFailed
	case NewOrder:
		return StatusNew
	case Payed:
		return StatusPayed
	default:
		return StatusUnknown
	}
}

func (s Status) String() string {
	switch s {
	case StatusNew:
		return NewOrder
	case StatusFailed:
		return Failed
	case StatusAwaitingPayment:
		return AwaitingPayment
	case StatusPayed:
		return Payed
	case StatusCancelled:
		return Cancelled
	default:
		return Unknown
	}
}
