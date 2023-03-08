package model

type Stock struct {
	WarehouseId int64  `json:"warehouseId"`
	Count       uint64 `json:"count"`
}
