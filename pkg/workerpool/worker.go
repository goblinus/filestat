package workerpool

import "sync"

const SingleWorker = 1

type Worker struct {
	TaskCh chan *Task
}

func NewWorker(ch chan *Task) Worker {
	return Worker{
		TaskCh: ch,
	}
}

func (w Worker) Start(wg *sync.WaitGroup) {
	wg.Add(SingleWorker)
	go func() {
		defer wg.Done()
		for task := range w.TaskCh {
			proccess(task)
		}
	}()
}
