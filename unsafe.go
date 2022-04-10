package set

import (
	"fmt"
	"strings"
)

type unsafeSet[T comparable] map[T]struct{}

func (s *unsafeSet[T]) Add(values ...T) {
	for _, value := range values {
		(*s)[value] = struct{}{}
	}
}

func (s *unsafeSet[T]) Delete(values ...T) {
	for _, value := range values {
		delete(*s, value)
	}
}

func (s *unsafeSet[T]) Clear() {
	*s = make(unsafeSet[T])
}

func (s *unsafeSet[T]) Has(values ...T) bool {
	for _, value := range values {
		if _, ok := (*s)[value]; !ok {
			return false
		}
	}
	return true
}

func (s *unsafeSet[T]) Len() int {
	return len(*s)
}

func (s *unsafeSet[T]) Clone() Set[T] {
	set := new(unsafeSet[T])
	set.Add(s.ToSlice()...)
	return set
}

func (s *unsafeSet[T]) ToSlice() []T {
	slice := make([]T, 0, s.Len())
	for value := range *s {
		slice = append(slice, value)
	}
	return slice
}

func (s *unsafeSet[T]) ForEach(f func(T)) {
	for value := range *s {
		f(value)
	}
}

func (s *unsafeSet[T]) String() string {
	values := make([]string, 0, s.Len())
	for value := range *s {
		values = append(values, fmt.Sprint(value))
	}
	return fmt.Sprintf("{%s}", strings.Join(values, ", "))
}
