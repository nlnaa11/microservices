package model

type Item struct {
	Sku   uint32 `json:"sku"`
	Count uint64 `json:"count"`
}

type ItemInfo struct {
	Item  Item    `json:"item"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}
