package main

import (
	// "fmt"
	"sync"
)

type ListNode struct {
	next, pre *ListNode
	value     interface{}
}

type DoubleList struct {
	head, tail *ListNode
	len        int
	lock       sync.Mutex
}

// GetValue 获取节点值
func (node *ListNode) GetValue() interface{} {
	return node.value
}

// GetPre 获取节点前驱节点
func (node *ListNode) GetPre() *ListNode {
	return node.pre
}

// GetNext 获取节点后驱节点
func (node *ListNode) GetNext() *ListNode {
	return node.next
}

// HashNext 是否存在后驱节点
func (node *ListNode) HashNext() bool {
	return node.pre != nil
}

// HashPre 是否存在前驱节点
func (node *ListNode) HashPre() bool {
	return node.next != nil
}

// IsNil 是否为空节点
func (node *ListNode) IsNil() bool {
	return node == nil
}

// Len 返回列表长度
func (list *DoubleList) Len() int {
	return list.len
}

// AddNodeFromHead 从头部开始 某个位置前插入节点
func (list *DoubleList) AddNodeFromHead(n int, v interface{}) {
	list.lock.Lock()
	defer list.lock.Unlock()

	if n != 0 && n >= list.len {
		panic("index out")
	}
	// 先取出head
	node := list.head
	// 往后遍历
	for i := 1; i <= n; i++ {
		node = node.next
	}
	newNode := new(ListNode)
	newNode.value = v
	if node.IsNil() {
		list.head = newNode
		list.tail = newNode
	} else {
		// 取插入节点的前一节点
		pre := node.pre
		// 如果该节点的前一节点是空
		if pre.IsNil() {
			newNode.next = node
			node.pre = newNode
			list.head = newNode
		} else {
			// 连接前面的节点
			pre.next = newNode
			newNode.pre = pre
			// 连接后面的节点
			node.pre = newNode
			newNode.next = node
		}
	}
	list.len = list.len + 1
}

// AddNodeFromTail 从尾部开始 某个位置后插入节点
func (list *DoubleList) AddNodeFromTail(n int, v interface{}) {
	list.lock.Lock()
	defer list.lock.Unlock()
	if n != 0 && n >= list.len {
		panic("index out")
	}
	// 取尾部
	node := list.tail
	for i := 1; i <= n; i++ {
		node = node.pre
	}
	newNode := new(ListNode)
	newNode.value = v
	if node.IsNil() {
		list.head = newNode
		list.tail = newNode
	} else {
		// 取节点的后驱节点
		next := node.next
		// 后驱节点为空，说明节点是最后一个节点
		if next.IsNil() {
			// 把新节点连接在尾部
			node.next = newNode
			newNode.pre = node
			list.tail = newNode
		} else {
			node.next = newNode
			newNode.pre = node

			newNode.next = next
			next.pre = newNode
		}
	}
	list.len = list.len + 1
}

// First 返回列表链表头结点
func (list *DoubleList) First() *ListNode {
	return list.head
}

// Last 返回列表链表尾结点
func (list *DoubleList) Last() *ListNode {
	return list.tail
}
func (list *DoubleList) IndexFromHead(n int) *ListNode {
	if n >= list.len {
		return nil
	}
	node := list.head
	for i := 1; i <= n; i++ {
		node = node.next
	}
	return node
}

func (list *DoubleList) IndexFromTail(n int) *ListNode {
	if n >= list.len {
		return nil
	}
	node := list.tail
	for i := 1; i <= n; i++ {
		node = node.pre
	}
	return node
}

func (list *DoubleList) PopFromHead(n int) *ListNode {
	list.lock.Lock()
	defer list.lock.Unlock()
	if n >= list.len {
		return nil
	}
	node := list.head
	for i := 1; i <= n; i++ {
		node = node.next
	}
	// 移除的节点 前驱节点和后驱节点
	next := node.next
	pre := node.pre
	// 前驱后驱都是空
	if next.IsNil() && pre.IsNil() {
		list.tail = nil
		list.head = nil
	} else if pre.IsNil() {
		list.head = next
		next.pre = nil
	} else if next.IsNil() {
		pre.next = nil
		list.tail = pre
	} else {
		pre.next = next
		next.pre = pre
	}
	list.len = list.len - 1
	return node
}

func (list *DoubleList) PopFromTail(n int) *ListNode {
	list.lock.Lock()
	defer list.lock.Unlock()
	if n >= list.len {
		return nil
	}
	node := list.tail
	for i := 1; i <= n; i++ {
		node = node.pre
	}
	// 移除的节点的前驱和后驱
	pre := node.pre
	next := node.next
	// 如果前驱和后驱都为nil，那么移除的节点为链表唯一节点
	if pre.IsNil() && next.IsNil() {
		list.head = nil
		list.tail = nil
	} else if pre.IsNil() {
		// 表示移除的是头部节点，那么下一个节点成为头节点
		list.head = next
		next.pre = nil
	} else if next.IsNil() {
		// 表示移除的是尾部节点，那么上一个节点成为尾节点
		list.tail = pre
		pre.next = nil
	} else {
		// 移除的是中间节点
		pre.next = next
		next.pre = pre
	}
	// 节点减一
	list.len = list.len - 1
	return node
}

// func main() {
// 	list := new(DoubleList)
// 	// 在列表头部插入新元素
// 	list.AddNodeFromHead(0, "I")
// 	list.AddNodeFromHead(0, "love")
// 	list.AddNodeFromHead(0, "you")
// 	// 在列表尾部插入新元素
// 	list.AddNodeFromTail(0, "may")
// 	list.AddNodeFromTail(0, "happy")
// 	list.AddNodeFromTail(list.Len()-1, "begin second")
// 	list.AddNodeFromHead(list.Len()-1, "end second")
// 	// 正常遍历，比较慢，因为内部会遍历拿到值返回
// 	for i := 0; i < list.Len(); i++ {
// 		// 从头部开始索引
// 		node := list.IndexFromHead(i)
// 		// 节点为空不可能，因为list.Len()使得索引不会越界
// 		if !node.IsNil() {
// 			fmt.Println(node.GetValue())
// 		}
// 	}
// 	fmt.Println("----------")
// 	// 正常遍历，特别快，因为直接拿到的链表节点
// 	// 先取出第一个元素
// 	first := list.First()
// 	for !first.IsNil() {
// 		// 如果非空就一直遍历
// 		fmt.Println(first.GetValue())
// 		// 接着下一个节点
// 		first = first.GetNext()
// 	}
// 	fmt.Println("----------")
// 	// 元素一个个 POP 出来
// 	for {
// 		node := list.PopFromHead(0)
// 		if node.IsNil() {
// 			// 没有元素了，直接返回
// 			break
// 		}
// 		fmt.Println(node.GetValue())
// 	}
// 	fmt.Println("----------")
// 	fmt.Println("len", list.Len())
// }
