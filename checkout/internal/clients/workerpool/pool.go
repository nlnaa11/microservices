// worker pool локальный. Создается по месту необходимости.
// Сразу получает все задачи на выполнение (срез tasks),
// которые ставит "в очередь" (tasksBuffer) для выполнения
// ограниченным числом рабочих (workersCount вызовов go worker(...))

// т.к. воркер пул локальный, после выполнения им всех задач
// переиспользовать его нельзя. Нужно создавать новый(

// для чтения результатов возвращает канал на чтение с результатами

package workerpool

import (
	"context"
	"sync"

	"gitlab.ozon.dev/nlnaa/homework-1/checkout/internal/clients/workerpool/task"
	libErr "gitlab.ozon.dev/nlnaa/homework-1/libs/errors"
)

type WorkerPool interface {
	Start(ctx context.Context)
	Results() <-chan task.Result
}

var _ WorkerPool = (*pool)(nil)

type pool struct {
	workersCount int

	tasks []task.Task

	// tasksBuffer to execute
	tasksBuffer chan task.Task
	// results after completing tasks
	results chan task.Result
}

func New(tasks []task.Task, workersCount int, tasksBufferSize int) (WorkerPool, error) {
	if workersCount < 1 {
		return nil, libErr.ErrNoWorkers
	}
	if tasksBufferSize < 1 {
		return nil, libErr.ErrInvalidTasksBuffer
	}

	tasksBuffer := make(chan task.Task, tasksBufferSize)
	results := make(chan task.Result, workersCount)

	return &pool{
		workersCount: workersCount,
		tasks:        tasks,
		tasksBuffer:  tasksBuffer,
		results:      results,
	}, nil
}

func (p *pool) Start(ctx context.Context) {
	var wg sync.WaitGroup

	// run workers
	for i := 0; i < p.workersCount; i++ {
		wg.Add(1)

		go worker(ctx, &wg, p.tasksBuffer, p.results)
	}

	// fill tasks to execute
	for _, task := range p.tasks {
		p.tasksBuffer <- task
	}

	// will close when all tasks flow into the tasks buffer
	close(p.tasksBuffer)

	// wait for completion of all tasks from the buffer
	wg.Wait()

	// will close when all results are received
	close(p.results)
}

func (p *pool) Results() <-chan task.Result {
	return p.results
}
