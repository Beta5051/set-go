package set

type Set[T comparable] interface {
	Add(...T)
	Delete(...T)
	Clear()
	Has(...T) bool
	Len() int
	Clone() Set[T]
	ToSlice() []T
	ForEach(func(T))
	String() string
}

func NewSet[T comparable](values ...T) Set[T] {
	set := make(unsafeSet[T])
	set.Add(values...)
	return &safeSet[T]{unsafe: set}
}

func NewUnsafeSet[T comparable](values ...T) Set[T] {
	set := make(unsafeSet[T])
	set.Add(values...)
	return &set
}
