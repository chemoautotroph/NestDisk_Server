package server

type WorkerPool struct {
	workerLen   int
	JobQueue    chan Job
	WorkerQueue chan chan Job
}

func NewWorkerPool(workerLen int) *WorkerPool {
	return &WorkerPool{
		workerLen:   workerLen,
		JobQueue:    make(chan Job),
		WorkerQueue: make(chan chan Job, workerLen),
	}
}

func (wp *WorkerPool) Run() {
	for i := 0; i < wp.workerLen; i++ {
		worker := NewWorker()
		worker.Run(wp.WorkerQueue)
	}
	go func() {
		for {
			select {
			case job := <-wp.JobQueue:
				worker := <-wp.WorkerQueue
				worker <- job
			}
		}
	}()
}
