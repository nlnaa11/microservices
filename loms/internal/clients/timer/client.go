package timer

import (
	"context"
)

type Order any

var (
	ToComplete chan Order
	ToQueue    chan Order
)

type Timer interface {
	Start(ctx context.Context)
	Stop()
}

type TimeManager interface {
	Close() error
	StartTimer(ctx context.Context) error
}

var _ TimeManager = (*timeManager)(nil)

type timeManager struct {
	timer Timer

	closeFunc context.CancelFunc
}

func New(ctx context.Context) *timeManager {
	_, cancel := context.WithCancel(ctx)

	return &timeManager{
		timer:     &timer{},
		closeFunc: cancel,
	}
}

func (tm *timeManager) StartTimer(ctx context.Context) error {
	tm.timer.Start(ctx)

	return nil
}

func (tm *timeManager) Close() error {
	if tm != nil {
		if tm.timer != nil {
			tm.timer.Stop()
		}

		if ToComplete != nil {
			close(ToComplete)
		}
		if ToQueue != nil {
			close(ToQueue)
		}

		if tm.closeFunc != nil {
			tm.closeFunc()
		}
	}

	return nil
}
