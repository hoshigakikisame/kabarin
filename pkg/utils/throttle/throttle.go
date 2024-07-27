package throttle

import (
	"sync"
	"time"
)

type Throttle struct {
	execQueue chan func()
	ticker    time.Ticker
	wg        sync.WaitGroup
	delay     time.Duration
}

func New(rateLimit int, delay int) (*Throttle, error) {
	if rateLimit <= 0 {
		rateLimit = 1
	}
	requestPerSecond := time.Second / time.Duration(rateLimit)

	ticker := time.NewTicker(requestPerSecond)

	return &Throttle{
		execQueue: make(chan func()),
		ticker:    *ticker,
		delay:     time.Second * time.Duration(delay),
	}, nil
}

func (t *Throttle) AddJob(job func()) {
	t.wg.Add(1)
	t.execQueue <- job
}

func (t *Throttle) Run() {
	go func() {
		for job := range t.execQueue {
			<-t.ticker.C
			job()
			t.wg.Done()
			time.Sleep(time.Second * 5)
		}
	}()
}

func (t *Throttle) Wait() {
	t.wg.Wait()
}
