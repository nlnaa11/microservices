package model

type Order struct {
	OrderId uint64 `json:"orderId"`
	Status  string `json:"status"`
}

type OrderInfo struct {
	Order Order  `json:"order"`
	User  int64  `json:"user"`
	Items []Item `json:"items"`
}
