package workerpool

import (
	"context"
	"log"
	"sync"
)

func worker(ctx context.Context, wg *sync.WaitGroup, tasks <-chan Task, results chan<- Result) {
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
