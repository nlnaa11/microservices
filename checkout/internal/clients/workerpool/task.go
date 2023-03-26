package workerpool

import (
	"context"

	libErr "gitlab.ozon.dev/nlnaa/homework-1/libs/errors"
)

type Task interface {
	Execute(ctx context.Context) Result
}

type Result struct {
	taskId uint64
	Value  interface{}
	Err    error
}

type TaskFn func(ctx context.Context, args ...interface{}) (interface{}, error)

var _ Task = (*task)(nil)

type task struct {
	id   uint64
	fn   TaskFn
	args []interface{}
}

func NewTask(id uint64, fn TaskFn, inputArgs ...interface{}) (Task, error) {
	if fn == nil {
		return nil, libErr.ErrInvalidTaskFunction
	}

	args := make([]interface{}, 0, len(inputArgs))
	args = append(args, inputArgs...)

	return &task{
		id:   id,
		fn:   fn,
		args: args,
	}, nil
}

func (t *task) Execute(ctx context.Context) Result {
	res, err := t.fn(ctx, t.args)
	if err != nil {
		return Result{
			taskId: t.id,
			Err:    err,
		}
	}

	return Result{
		taskId: t.id,
		Value:  res,
	}
}
