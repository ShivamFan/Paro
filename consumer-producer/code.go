// 1. Program a multi-threaded producer and consumer with locks on concurrent
// access to a memory location using C/C++/Golang.
// Explanation:
// Have a global list / array where two or more producers append randomly
// generated numbers concurrently. A consumer thread must dequeue from the global
// list and print it!

package consumerproducer

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	numProducers = 2
	numConsumers = 1
	bufferSize   = 10
)

var (
	globalList []int
	mutex      sync.Mutex
	notEmpty   = sync.NewCond(&mutex)
	notFull    = sync.NewCond(&mutex)
)

func producer(id int, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		num := rand.Intn(100)
		fmt.Printf("Producer %d produced: %d\n", id, num)

		mutex.Lock()
		for len(globalList) >= bufferSize {
			notFull.Wait()
		}
		globalList = append(globalList, num)
		mutex.Unlock()

		notEmpty.Signal()

		time.Sleep(time.Millisecond * time.Duration(rand.Intn(500)))
	}
}

func consumer(id int, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		mutex.Lock()
		for len(globalList) == 0 {
			notEmpty.Wait()
		}
		num := globalList[0]
		globalList = globalList[1:]
		fmt.Printf("Consumer %d consumed: %d\n", id, num)
		mutex.Unlock()

		notFull.Signal()

		time.Sleep(time.Millisecond * time.Duration(rand.Intn(500)))
	}
}

func Ques1() {
	rand.Seed(time.Now().UnixNano())

	var wg sync.WaitGroup

	for i := 0; i < numProducers; i++ {
		wg.Add(1)
		go producer(i, &wg)
	}

	for i := 0; i < numConsumers; i++ {
		wg.Add(1)
		go consumer(i, &wg)
	}

	wg.Wait()
}
