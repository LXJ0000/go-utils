package setx

import "sync"

// Set 泛型集合接口，定义了集合的基本操作。
type Set[T comparable] interface {
	Add(key T)
	Delete(key T)
	Exists(key T) bool
	Len() int
	Clear()
	Elements() []T
}

// MapSet 是 Set 接口的一个实现，使用 map 来存储集合元素，并使用 RWMutex 保护并发访问。
type MapSet[T comparable] struct {
	mu sync.RWMutex
	m  map[T]struct{}
}

// NewMapSet 创建一个新的 MapSet 实例。
func NewMapSet[T comparable](size int) *MapSet[T] {
	return &MapSet[T]{
		m: make(map[T]struct{}, size),
	}
}

func (s *MapSet[T]) WithSlice(slice []T) *MapSet[T] {
    s.mu.Lock()
    defer s.mu.Unlock()
    s.m = make(map[T]struct{}, len(slice))
    for _, v := range slice {
        s.m[v] = struct{}{}
    }
    return s
}

// Add 向集合中添加一个元素。
func (s *MapSet[T]) Add(key ...T) {
	s.mu.Lock()
	defer s.mu.Unlock()
	for _, k := range key {
		s.m[k] = struct{}{}
	}
}

// Delete 从集合中删除一个元素。
func (s *MapSet[T]) Delete(key T) {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.m, key)
}

// Exists 检查集合中是否存在某个元素。
func (s *MapSet[T]) Exists(key T) bool {
	s.mu.RLock()
	defer s.mu.RUnlock()
	_, ok := s.m[key]
	return ok
}

// Len 返回集合的大小。
func (s *MapSet[T]) Len() int {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return len(s.m)
}

// Clear 清空集合。
func (s *MapSet[T]) Clear() {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.m = make(map[T]struct{})
}

// Elements 返回集合中的所有元素。
func (s *MapSet[T]) Elements() []T {
	s.mu.RLock()
	defer s.mu.RUnlock()
	elements := make([]T, 0, len(s.m))
	for key := range s.m {
		elements = append(elements, key)
	}
	return elements
}
