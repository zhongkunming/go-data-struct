package main

// import "sync"

// type LinkNode4Queue struct {
// 	Next *LinkNode4Queue
// 	Value interface{}
// }

// type LinkQueue struct {
// 	root *LinkNode4Queue
// 	size int 
// 	lock sync.Mutex
// }

// func (queue *LinkQueue) Add(e interface{})  {
// 	queue.lock.Lock()
// 	defer queue.lock.Unlock()

// 	newNode := new(LinkNode4Queue)
// 	newNode.Value = e
// 	if queue.root != nil {
// 		currentNode := queue.root
// 		for currentNode.Next != nil{
// 			currentNode = currentNode.Next
// 		}
// 		currentNode.Next  = newNode
// 	} else {
// 		queue.root = newNode
// 	}
	
// 	queue.size = queue.size + 1
// }

// // 出队
// func (queue *LinkQueue) Remove() interface{} {
//     queue.lock.Lock()
//     defer queue.lock.Unlock()
//     if queue.size == 0 {
//         panic("empty")
//     }
//     topNode := queue.root
//     v := topNode.Value
//     queue.root = topNode.Next
//     queue.size = queue.size - 1
//     return v
// }