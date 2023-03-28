package workerpool

import (
	"context"
	"log"
	"sync"

	"gitlab.ozon.dev/nlnaa/homework-1/checkout/internal/clients/workerpool/task"
)

func worker(ctx context.Context, wg *sync.WaitGroup, tasks <-chan task.Task, results chan<- task.Result) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			log.Println("something was wrong")
			return

		case task, ok := <-tasks:
			if !ok {
				return
			}
			results <- task.Execute(ctx)
		}
	}
}
