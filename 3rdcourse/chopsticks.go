package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type ChopS struct{ sync.Mutex }

type Semaphore chan struct{}

func NewSemaphore(size int) Semaphore {
	return make(Semaphore, size)
}

func (s Semaphore) Lock() {
	s <- struct{}{}
}

func (s Semaphore) Unlock() {
	<-s
}

type Philo struct {
	number          int
	leftCS, rightCS *ChopS
}

func (p Philo) eat(wg *sync.WaitGroup, host *Semaphore) {
	for i := 0; i < 3; i++ {
		host.Lock()

		if rand.Intn(1) == 0 {
			p.leftCS.Lock()
			p.rightCS.Lock()
		} else {
			p.rightCS.Lock()
			p.leftCS.Lock()
		}

		fmt.Println("Starting to eat ", p.number)
		time.Sleep(time.Duration(rand.Intn(4)) * time.Second)
		fmt.Println("Finishing eating ", p.number)
		p.rightCS.Unlock()
		p.leftCS.Unlock()
		host.Unlock()
	}

	wg.Done()
}

func main() {
	var wg sync.WaitGroup

	var host = NewSemaphore(2)

	CSticks := make([]*ChopS, 5)
	for i := 0; i < 5; i++ {
		CSticks[i] = new(ChopS)
	}
	philos := make([]*Philo, 5)
	for i := 0; i < 5; i++ {
		philos[i] = &Philo{
			number:  i + 1,
			leftCS:  CSticks[i],
			rightCS: CSticks[(i+1)%5]}
	}

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go philos[i].eat(&wg, &host)
	}
	wg.Wait()

}
