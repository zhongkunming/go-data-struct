package main

import "sync"

type ArrayQueue struct {
	array []interface{}
	size  int
	lock  sync.Mutex
}

func (queue *ArrayQueue) Add(e interface{}) {
	queue.lock.Lock()
	defer queue.lock.Unlock()

	queue.array = append(queue.array, e)
	queue.size = queue.size + 1
}

func (queue *ArrayQueue) Remove() interface{} {
	queue.lock.Lock()
	defer queue.lock.Unlock()
	if queue.size == 0 {
		panic("emoty")
	}
	v := queue.array[0]
	newArray := make([]interface{}, queue.size-1, queue.size-1)
	for i := 1; i < queue.size; i++ {
		newArray[i-1] = queue.array[i]
	}
	queue.array = newArray
	queue.size = queue.size - 1
	return v
}
