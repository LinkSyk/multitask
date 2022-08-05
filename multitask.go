package multitask

import (
	"sync"
)

type Source func(ch chan<- interface{})

type Sink func(msg interface{})

type MultiTask struct {
	queue     []Source
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

func (t *MultiTask) Do(w Source) {
	t.queue = append(t.queue, w)
}

func (t *MultiTask) Recv(s Sink) {
	t.isSetRecv = true
	go func() {
		for msg := range t.msgCh {
			if s == nil {
				return
			}
			s(msg)
		}
		t.done <- struct{}{}
	}()
}

func (t *MultiTask) Wait() {
	if !t.isSetRecv {
		panic("not set Recv function!")
	}
	var wg sync.WaitGroup
	wg.Add(len(t.queue))
	for _, s := range t.queue {
		go func(s Source) {
			defer wg.Done()
			s(t.msgCh)
		}(s)
	}
	wg.Wait()
	close(t.msgCh)
	<-t.done
}
