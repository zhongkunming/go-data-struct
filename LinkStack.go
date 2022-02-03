package main

// import (
// 	// "fmt"
// 	"sync"
// )

// type LinkNode4Stack struct {
// 	Next  *LinkNode4Stack
// 	Value interface{}
// }

// type LinkStack struct {
// 	root *LinkNode4Stack
// 	size int
// 	lock sync.Mutex
// }

// func (stack *LinkStack) Push(e interface{}) {
// 	stack.lock.Lock()
// 	defer stack.lock.Unlock()
// 	newNode := new(LinkNode4Stack)
// 	newNode.Value = e
// 	if stack.root != nil {
// 		currentNode := stack.root
// 		newNode.Next = currentNode
// 	}
// 	stack.root = newNode
// 	stack.size = stack.size + 1
// }

// // 出栈
// func (stack *LinkStack) Pop() interface{} {
// 	stack.lock.Lock()
// 	defer stack.lock.Unlock()
// 	// 栈中元素已空
// 	if stack.size == 0 {
// 		panic("empty")
// 	}
// 	// 顶部元素要出栈
// 	topNode := stack.root
// 	v := topNode.Value
// 	// 将顶部元素的后继链接链上
// 	stack.root = topNode.Next
// 	// 栈中元素数量-1
// 	stack.size = stack.size - 1
// 	return v
// }

// // 获取栈顶元素
// func (stack *LinkStack) Peek() interface{} {
// 	// 栈中元素已空
// 	if stack.size == 0 {
// 		panic("empty")
// 	}
	
// 	// 顶部元素值
// 	v := stack.root.Value
// 	return v
// } // 栈大小
// func (stack *LinkStack) Size() int {
// 	return stack.size
// }

// // 栈是否为空
// func (stack *LinkStack) IsEmpty() bool {
	
// 	return stack.size == 0
// }
// func main() {
// 	linkStack := new(LinkStack)
// 	linkStack.Push("cat")
// 	linkStack.Push("dog")
// 	linkStack.Push("hen")
// 	fmt.Println("size:", linkStack.Size())
// 	fmt.Println("pop:", linkStack.Pop())
// 	fmt.Println("pop:", linkStack.Pop())
// 	fmt.Println("size:", linkStack.Size())
// 	linkStack.Push("drag")
// 	fmt.Println("pop:", linkStack.Pop())
// }
