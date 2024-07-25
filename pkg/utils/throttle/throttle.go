package throttle

import (
	"sync"
	"time"
)

type Throttle struct {
	execQueue chan func()
	ticker    time.Ticker
	wg        sync.WaitGroup
}

func New(rateLimit time.Duration) (*Throttle, error) {
	if rateLimit <= 0 {
		rateLimit = 1
	}
	rateLimit = time.Second / rateLimit
	ticker := time.NewTicker(rateLimit)

	return &Throttle{
		execQueue: make(chan func()),
		ticker:    *ticker,
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
		}
	}()
}

func (t *Throttle) Wait() {
	t.wg.Wait()
}
