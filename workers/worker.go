package workers

import "log"

type Worker struct {
	ID         int
	Queue      TaskQueue
	WorkerPool chan TaskQueue
	Quit       chan bool
}

func NewWorker(id int, workerPool chan TaskQueue) *Worker {
	return &Worker{
		ID:         id,
		Queue:      make(TaskQueue),
		WorkerPool: workerPool,
		Quit:       make(chan bool),
	}
}

func (w *Worker) Start() {
	go func() {
		for {
			w.WorkerPool <- w.Queue

			select {
			case task := <-w.Queue:
				task()
			case <-w.Quit:
				log.Printf("worker %d stopping", w.ID)
				return
			}
		}
	}()
}

func (w *Worker) Stop() {
	go func() {
		w.Quit <- true
	}()
}
