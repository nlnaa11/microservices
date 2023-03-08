package model

type Item struct {
	Sku   uint32 `json:"sku"`
	Count uint64 `json:"count"`
}

type ItemInfo struct {
	Sku   uint32  `json:"sku"`
	Count uint64  `json:"count"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}
