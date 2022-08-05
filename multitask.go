package multitask

import (
	"sync"
)

// Task represent a single task.
// resultCh receive result from task, you also can ingore it if you have nothing to send.
type Task func(resultCh chan<- interface{})

// ResultFunc receive all result from all task.
type ResultFunc func(result interface{})

type MultiTask struct {
	queue     []Task
	msgCh     chan interface{}
	done      chan struct{}
	queueSize int
	isSetRecv bool
}

type TaskOpt func(t *MultiTask)

func WithQueueSize(size int) TaskOpt {
	return func(t *MultiTask) {
		t.queueSize = size
	}
}

// NewMultiTask create MultiTask instance
func NewMultiTask(opts ...TaskOpt) *MultiTask {
	t := &MultiTask{
		queueSize: 400,
	}

	for _, opt := range opts {
		opt(t)
	}
	t.msgCh = make(chan interface{}, t.queueSize)
	t.done = make(chan struct{}, 1)
	return t
}

// Do accept a task, push it to task queue, execute until MultiTask.Wait() is called
func (t *MultiTask) Do(task Task) {
	t.queue = append(t.queue, task)
}

// Fetch set a result function to receive result from task function.
func (t *MultiTask) Fetch(fn ResultFunc) {
	t.isSetRecv = true
	go func() {
		for msg := range t.msgCh {
			if fn == nil {
				return
			}
			fn(msg)
		}
		t.done <- struct{}{}
	}()
}

// Wait() block until all tasks done
func (t *MultiTask) Wait() {
	if !t.isSetRecv {
		t.Fetch(doNothing)
	}
	var wg sync.WaitGroup
	wg.Add(len(t.queue))
	for _, task := range t.queue {
		task := task
		go func() {
			defer wg.Done()
			task(t.msgCh)
		}()
	}
	wg.Wait()
	close(t.msgCh)
	<-t.done
}

func doNothing(msg interface{}) {
}
