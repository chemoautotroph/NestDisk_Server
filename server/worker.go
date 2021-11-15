package server

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
	go func() {
		for {
			wq <- w.JobQueue
			select {
			case job := <-w.JobQueue:
				 // log.Println("worker receive job: ", job)
				job.Do()
			case <-w.Quit:
				// log.Println("worker Quit")
				return
			}
		}
	}()
}
