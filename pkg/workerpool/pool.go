package workerpool

import "sync"

const PoolCapacity = 2

type Pool struct {
	Tasks     []*Task
	capacity  int
	collector chan *Task
	wg        *sync.WaitGroup
}

func NewPool(tasks []*Task, capacity int) Pool {
	return Pool{
		Tasks:     tasks,
		capacity:  capacity,
		collector: make(chan *Task, PoolCapacity),
		wg:        &sync.WaitGroup{},
	}
}

//Run создает и запускает на выполнение пулл задач
func (p *Pool) Run() {
	for i := 0; i < p.capacity; i++ {
		worker := NewWorker(p.collector)
		worker.Start(p.wg)
	}
	for i := range p.Tasks {
		p.collector <- p.Tasks[i]
	}
	close(p.collector)
	p.wg.Wait()
}
