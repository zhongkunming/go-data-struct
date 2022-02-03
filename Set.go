package main

// import (
// 	"fmt"
// 	"sync"
// )

// type Set struct {
// 	m   map[interface{}]struct{}
// 	len int
// 	sync.RWMutex
// }

// func NewSet(cap int) *Set {
// 	temp := make(map[interface{}]struct{}, cap)
// 	return &Set{
// 		m: temp,
// 	}
// }

// func (s *Set) Add(e interface{}) {
// 	s.Lock()
// 	defer s.Unlock()
// 	s.m[e] = struct{}{}
// 	s.len = len(s.m)
// }

// func (s *Set) Remove(e interface{}) {
// 	s.Lock()
// 	defer s.Unlock()
// 	if s.len == 0 {
// 		return
// 	}
// 	delete(s.m, e)
// 	s.len = len(s.m)
// }

// func (s *Set) Has(e interface{}) bool {
// 	s.Lock()
// 	defer s.Unlock()
// 	_, ok := s.m[e]
// 	return ok
// }

// func (s *Set) Len() int {
// 	return s.len
// }

// func (s *Set) IsEmpty() bool {
// 	return s.len == 0
// }

// func (s *Set) Clear() {
// 	s.Lock()
// 	defer s.Unlock()
// 	s.m = map[interface{}]struct{}{}
// 	s.len = 0
// }

// func (s *Set) List() []interface{} {
// 	s.RLock()
// 	defer s.RUnlock()
// 	list := make([]interface{}, 0, s.len)
// 	for k := range s.m {
// 		list = append(list, k)
// 	}
// 	return list

// }
// func main() {
// 	//other()
// 	// 初始化一个容量为5的不可重复集合
// 	s := NewSet(5)
// 	s.Add(1)
// 	s.Add(1)
// 	s.Add(2)
// 	fmt.Println("list of all items", s.List())
// 	s.Clear()
// 	if s.IsEmpty() {
// 		fmt.Println("empty")
// 	}
// 	s.Add(1)
// 	s.Add(2)
// 	s.Add(3)
// 	if s.Has(2) {
// 		fmt.Println("2 does exist")
// 	}
// 	s.Remove(2)
// 	s.Remove(3)
// 	fmt.Println("list of all items", s.List())
// }
