package main

import (
	"fmt"
	"sync"
)

type Array struct {
	array []interface{}
	len   int // 长度
	cap   int // 容量
	lock  sync.Mutex
}

func Make(len, cap int) *Array {
	arr := new(Array)
	if len > cap {
		panic("len > cap")
	}
	array := make([]interface{}, cap, cap)
	arr.array = array
	arr.len = 0
	arr.cap = cap
	return arr
}

func (a *Array) Append(e interface{}) {
	a.lock.Lock()
	defer a.lock.Unlock()

	if a.len == a.cap {
		// 扩容
		newCap := 2 * a.len
		if a.cap == 0 {
			newCap = 1
		}
		newArray := make([]interface{}, newCap, newCap)
		for k, v := range a.array {
			newArray[k] = v
		}
		a.array = newArray
		a.cap = newCap
	}
	a.array[a.len] = e
	a.len = a.len + 1
}

// AppendMany 增加多个元素
func (a *Array) AppendMany(e ...interface{}) {
	for _, v := range e {
		a.Append(v)
	}
}

// Get 获取某个下标的元素
func (a *Array) Get(index int) interface{} {
	// 越界了
	if a.len == 0 || index >= a.len {
		panic("index over len")
	}
	return a.array[index]
}

// Len 返回真实长度
func (a *Array) Len() int {
	return a.len
}

// Cap 返回容量
func (a *Array) Cap() int {
	return a.cap
}

// Print 辅助打印
func Print(array *Array) (result string) {
	result = "["
	for i := 0; i < array.Len(); i++ {
		// 第一个元素
		if i == 0 {
			result = fmt.Sprintf("%s%d", result, array.Get(i))
			continue
		}
		result = fmt.Sprintf("%s %d", result, array.Get(i))
	}
	result = result + "]"
	return
}
// func main() {
// 	// 创建一个容量为3的动态数组
// 	a := Make(0, 3)
// 	fmt.Println("cap", a.Cap(), "len", a.Len(), "array:", Print(a))
// 	// 增加一个元素
// 	a.Append(10)
// 	fmt.Println("cap", a.Cap(), "len", a.Len(), "array:", Print(a))
// 	// 增加一个元素
// 	a.Append(9)
// 	fmt.Println("cap", a.Cap(), "len", a.Len(), "array:", Print(a))
// 	// 增加多个元素
// 	a.AppendMany(8, 7)
// 	fmt.Println("cap", a.Cap(), "len", a.Len(), "array:", Print(a))
// }
