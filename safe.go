package set

import "sync"

type safeSet[T comparable] struct {
	sync.RWMutex
	unsafe unsafeSet[T]
}

func (s *safeSet[T]) Add(values ...T) {
	s.Lock()
	s.unsafe.Add(values...)
	s.Unlock()
}

func (s *safeSet[T]) Delete(values ...T) {
	s.Lock()
	s.unsafe.Delete(values...)
	s.Unlock()
}

func (s *safeSet[T]) Clear() {
	s.Lock()
	s.unsafe.Clear()
	s.Unlock()
}

func (s *safeSet[T]) Has(values ...T) bool {
	s.RLock()
	result := s.unsafe.Has(values...)
	s.RUnlock()
	return result
}

func (s *safeSet[T]) Len() int {
	s.RLock()
	result := s.unsafe.Len()
	s.RUnlock()
	return result
}

func (s *safeSet[T]) Clone() Set[T] {
	s.RLock()
	result := s.unsafe.Clone()
	s.RUnlock()
	return result
}

func (s *safeSet[T]) ToSlice() []T {
	s.RLock()
	result := s.unsafe.ToSlice()
	s.RUnlock()
	return result
}

func (s *safeSet[T]) ForEach(f func(T)) {
	s.RLock()
	s.unsafe.ForEach(f)
	s.RUnlock()
}

func (s *safeSet[T]) String() string {
	s.RLock()
	result := s.unsafe.String()
	s.RUnlock()
	return result
}
