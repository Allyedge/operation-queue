package queue

import (
	"sync"

	"github.com/Allyedge/operation-queue/operation"
)

type Queue struct {
	lock   *sync.Mutex
	Values []operation.Operation
}

func Init() Queue {
	return Queue{&sync.Mutex{}, make([]operation.Operation, 0)}
}

func (queue *Queue) Push(value operation.Operation) {
	queue.lock.Lock()
	queue.Values = append(queue.Values, value)
	queue.lock.Unlock()
}

func (queue *Queue) Pop() operation.Operation {
	if len(queue.Values) > 0 {
		queue.lock.Lock()

		pop := queue.Values[0]

		queue.Values = queue.Values[1:]

		queue.lock.Unlock()

		return pop
	}

	return operation.Operation{}
}
