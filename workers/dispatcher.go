package workers

type Dispatcher struct {
	WorkerPool chan TaskQueue
	MaxWorkers int
	TaskQueue  TaskQueue
}

func NewDispatcher(maxWorkers int, queueSize int) *Dispatcher {
	return &Dispatcher{
		WorkerPool: make(chan TaskQueue, maxWorkers),
		MaxWorkers: maxWorkers,
		TaskQueue:  make(TaskQueue, queueSize),
	}
}

func (d *Dispatcher) Run() {
	for i := 0; i < d.MaxWorkers; i++ {
		worker := NewWorker(i, d.WorkerPool)
		worker.Start()
	}

	go d.dispatch()
}

func (d *Dispatcher) dispatch() {
	for task := range d.TaskQueue {
		workerQueue := <-d.WorkerPool
		workerQueue <- task
	}
}
