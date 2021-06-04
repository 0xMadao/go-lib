package workpool

import (
	"sync"
)

// WaitGroup pool struct
type WaitGroup struct {
	workChan chan int
	wg       sync.WaitGroup
}

// NewPool create a new pool with concurrency limit
func NewPool(coreNum int) *WaitGroup {
	ch := make(chan int, coreNum)
	return &WaitGroup{
		workChan: ch,
		wg:       sync.WaitGroup{},
	}
}

// Add add
func (ap *WaitGroup) Add(num int) {
	for i := 0; i < num; i++ {
		ap.workChan <- i
		ap.wg.Add(1)
	}
}

// Done finish
func (ap *WaitGroup) Done() {
	<-ap.workChan
	ap.wg.Done()
}

// Wait wait
func (ap *WaitGroup) Wait() {
	ap.wg.Wait()
}
