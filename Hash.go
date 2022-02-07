package main

import (
	"fmt"
	"github.com/OneOfOne/xxhash"
	math "math"
	"sync"
)

const (
	// 扩容因子
	expandFactor = 0.75
)

// KeyPairs 键值对
type KeyPairs struct {
	key   string
	value interface{}
	next  *KeyPairs
}

// HashMap 哈希表
type HashMap struct {
	array        []*KeyPairs // 哈希表数组，每一个元素是一个键值对
	capacity     int         // 数组容量
	len          int         // 已添加键值对元素数量
	capacityMask int         // 掩码 capacity -1
	lock         sync.Mutex
}

func NewHashMap(capacity int) *HashMap {
	defaultCapacity := 1 << 4
	if capacity < defaultCapacity {
		capacity = defaultCapacity
	} else {
		capacity = 1 << int(math.Ceil(math.Log2(float64(capacity))))
	}
	m := new(HashMap)
	m.array = make([]*KeyPairs, capacity, capacity)
	m.capacity = capacity
	m.capacityMask = capacity - 1
	return m
}

func (m *HashMap) Len() int {
	return m.len
}

func (m *HashMap) Capacity() int {
	return m.capacity
}

var hashAlgorithm = func(key []byte) uint64 {
	h := xxhash.New64()
	h.Write(key)
	return h.Sum64()
}

// 对键进行哈希求值，并计算下标
func (m *HashMap) hashIndex(key string, mask int) int {
	hash := hashAlgorithm([]byte(key))
	index := hash & uint64(mask)
	return int(index)
}
func (m *HashMap) Put(key string, value interface{}) {
	m.lock.Lock()
	defer m.lock.Unlock()

	index := m.hashIndex(key, m.capacityMask)
	element := m.array[index]

	if element == nil {
		m.array[index] = &KeyPairs{
			key:   key,
			value: value,
		}
	} else {
		// 链表最后一个键值对
		var lastKeyPairs *KeyPairs
		// 遍历链表查看元素是否存在，存在则替换值，否则找到最后一个键值对
		for element != nil {
			// 键值对存在，那么更新值并返回
			if element.key == key {
				element.value = value
				return
			}
			lastKeyPairs = element
			element = element.next
		}
		// 找不到键值对，将新键值对添加到链表尾端
		lastKeyPairs.next = &KeyPairs{
			key:   key,
			value: value,
		}
	}

	newLen := m.len + 1
	if float64(newLen)/float64(m.capacity) >= expandFactor {
		newM := new(HashMap)
		newM.array = make([]*KeyPairs, 2*m.capacity, 2*m.capacity)
		newM.capacity = 2 * m.capacity
		newM.capacityMask = 2*m.capacity - 1
		for _, pairs := range m.array {
			for pairs != nil {
				newM.Put(pairs.key, pairs.value)
				pairs = pairs.next
			}
		}
		m.array = newM.array
		m.capacity = newM.capacity
		m.capacityMask = newM.capacityMask
	}
	m.len = newLen
}

func (m *HashMap) Get(key string) (value interface{}, ok bool) {
	m.lock.Lock()
	defer m.lock.Unlock()

	index := m.hashIndex(key, m.capacityMask)
	element := m.array[index]
	for element != nil {
		if element.key == key {
			return element.value, true
		}
		element = element.next
	}
	return
}

func (m *HashMap) Delete(key string) {
	m.lock.Lock()
	defer m.lock.Unlock()

	index := m.hashIndex(key, m.capacityMask)
	element := m.array[index]

	if element == nil {
		return
	}

	if element.key == key {
		m.array[index] = element.next
		m.len = m.len - 1
		return
	}
	nextElement := element.next
	for nextElement != nil {
		if nextElement.key == key {
			element.next = nextElement.next
			m.len = m.len - 1
			return
		}
		element = nextElement
		nextElement = nextElement.next
	}
}

func (m *HashMap) Range() {
	// 实现并发安全
	m.lock.Lock()
	defer m.lock.Unlock()
	for _, pairs := range m.array {
		for pairs != nil {
			fmt.Printf("'%v'='%v',", pairs.key, pairs.value)
			pairs = pairs.next
		}
	}
	fmt.Println()
}

func main() {
	// 新建一个哈希表
	hashMap := NewHashMap(16)
	// 放35个值
	for i := 0; i < 35; i++ {
		hashMap.Put(fmt.Sprintf("%d", i), fmt.Sprintf("v%d", i))
	}
	fmt.Println("cap:", hashMap.Capacity(), "len:", hashMap.Len())
	// 打印全部键值对
	hashMap.Range()
	key := "4"
	value, ok := hashMap.Get(key)
	if ok {
		fmt.Printf("get '%v'='%v'\n", key, value)
	} else {
		fmt.Printf("get %v not found\n", key)
	}
	// 删除键
	hashMap.Delete(key)
	fmt.Println("after delete cap:", hashMap.Capacity(), "len:", hashMap.Len())
	value, ok = hashMap.Get(key)
	if ok {
		fmt.Printf("get '%v'='%v'\n", key, value)
	} else {
		fmt.Printf("get %v not found\n", key)
	}
}
