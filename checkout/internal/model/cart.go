package model

type Cart struct {
	Items []Item `json:"items"`
}

type CartList struct {
	ItemsInfo  []ItemInfo `json:"itemsInfo"`
	TotalPrice float64    `json:"totalPrice"`
}
