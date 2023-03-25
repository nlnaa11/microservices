package table

const (
	OrdersItems = "orders_items"
	Orders      = "orders"
	OrdersUsers = "orders_users"

	ItemsStocks   = "items_stocks"
	ReservedItems = "reserved_items"
)

// orders_items: orderId + itemId + count
// orders: orderId + status
// orders_users: orderId + userId
// items_stocks: warehouseId + itemId + count
// reserved_items: itemId + count
