package workers

type WorkerPool struct {
    dispatcher *Dispatcher
}

func NewWorkerPool(maxWorkers int) *WorkerPool {
    dispatcher := NewDispatcher(maxWorkers, 1000)
    dispatcher.Run()
    return &WorkerPool{dispatcher: dispatcher}
}

func (wp *WorkerPool) Submit(task Task) {
    wp.dispatcher.TaskQueue <- task
}
