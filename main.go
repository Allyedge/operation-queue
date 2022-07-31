package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"

	"github.com/Allyedge/operation-queue/operation"
	"github.com/Allyedge/operation-queue/queue"
)

const MAXIMUM_WORKERS = 100
const MAXIMUM_CREATE_TIME = 100
const MAXIMUM_FINISH_TIME = 1000

func main() {
	operationQueue := queue.Init()

	notify := make(chan struct{}, 1)

	wg := &sync.WaitGroup{}

	for i := 0; i < MAXIMUM_WORKERS; i++ {
		wg.Add(1)
		go worker(wg, notify, &operationQueue)
	}

	for i := 0; i > -1; i++ {
		random := rand.Intn(MAXIMUM_CREATE_TIME)

		time.Sleep(time.Duration(random) * time.Millisecond)

		operationQueue.Push(operation.Operation{ID: i})

		fmt.Printf("+++ CREATE OPERATION %d +++\n", i)

		notify <- struct{}{}
	}
}

func worker(wg *sync.WaitGroup, notify chan struct{}, operationQueue *queue.Queue) {
	defer wg.Done()

	for range notify {
		task := operationQueue.Pop()

		doOperation(task)
	}
}

func doOperation(operation operation.Operation) {
	random := rand.Intn(MAXIMUM_FINISH_TIME)

	time.Sleep(time.Duration(random) * time.Millisecond)

	fmt.Printf("--- FINISH OPERATION %d ---\n", operation.ID)
}
