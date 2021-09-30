package server

import "log"

type Job interface {
	Do()
}

type Worker struct {
	JobQueue chan Job
	Quit     chan bool
}

func NewWorker() Worker {
	return Worker{
		JobQueue: make(chan Job),
		Quit:     make(chan bool),
	}
}

func (w Worker) Run(wq chan chan Job) {
	// 保证 每读到一个通道参数就 去做这件事，没读到就阻塞
	go func() {
		for {
			wq <- w.JobQueue
			select {
			case job := <-w.JobQueue:
				job.Do()
			case <-w.Quit:
				log.Printf("worker退出到pool\n")
				return
			}
		}
	}()
}

type WorkerPool struct {
	workerLen   int
	JobQueue    chan Job
	WorkerQueue chan chan Job
}

func NewWorkerPool(workerLen int) *WorkerPool {
	return &WorkerPool{
		workerLen:   workerLen,
		JobQueue:    make(chan Job),
		WorkerQueue: make(chan chan Job),
	}
}

func (wp *WorkerPool) Run() {
	for i := 0; i < wp.workerLen; i++ {
		worker := NewWorker()
		worker.Run(wp.WorkerQueue)
	}
	go func() {
		select {
		case job := <-wp.JobQueue: // 当读到任务
			log.Printf("读取到任务\n")
			worker := <-wp.WorkerQueue
			worker <- job // 把任务分给worker
		}
	}()
}
