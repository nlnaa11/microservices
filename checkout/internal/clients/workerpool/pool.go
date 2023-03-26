package workerpool

import (
	"context"
	"sync"

	libErr "gitlab.ozon.dev/nlnaa/homework-1/libs/errors"
)

type Pool interface {
	Start(ctx context.Context)
	Results() <-chan Result
}

var _ Pool = (*pool)(nil)

type pool struct {
	workersCount int

	tasks []Task

	// tasksBuffer to execute
	tasksBuffer chan Task
	// results after completing tasks
	results chan Result
}

func New(tasks []Task, workersCount int, tasksBufferSize int) (Pool, error) {
	if workersCount < 1 {
		return nil, libErr.ErrNoWorkers
	}
	if tasksBufferSize < 1 {
		return nil, libErr.ErrInvalidTasksBuffer
	}

	tasksBuffer := make(chan Task, tasksBufferSize)
	results := make(chan Result, workersCount)

	return &pool{
		workersCount: workersCount,
		tasks:        tasks,
		tasksBuffer:  tasksBuffer,
		results:      results,
	}, nil
}

func (p *pool) Start(ctx context.Context) {
	var wg sync.WaitGroup

	defer func() {
		// will close when all tasks flow into the tasks buffer
		// [from tasks to tasksBuffer]
		close(p.tasksBuffer)

		// wait for completion of all tasks from the buffer
		wg.Wait()

		// will close when all results are received
		close(p.results)
	}()

	// run workers
	for i := 0; i < p.workersCount; i++ {
		wg.Add(1)

		go worker(ctx, &wg, p.tasksBuffer, p.results)
	}

	// fill tasks to execute
	for _, task := range p.tasks {
		p.tasksBuffer <- task
	}
}

func (p *pool) Results() <-chan Result {
	return p.results
}
