package storage

type CartStorage interface {
	AddTo(item Item, user int64, successChan chan bool)
	RemoveFrom(item Item, user int64, successChan chan bool)
	Get(user int64) (*Cart, error)
}

type WrapStorage struct {
	cartStor CartStorage
}

func New(stor CartStorage) *WrapStorage {
	return &WrapStorage{
		cartStor: stor,
	}
}
