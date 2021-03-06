package frames

import "sync"

func ForEach[T any](begin Iterator[T], end Iterator[T], f func(*T)) {
	for ; begin.Index() != end.Index(); begin.Next() {
		f(begin.Ptr())
	}
}

func ForEachIndexed[T any](begin Iterator[T], end Iterator[T], f func(*T, uint)) {
	i := uint(0)
	for ; begin.Index() != end.Index(); begin.Next() {
		f(begin.Ptr(), i)
		i++
	}
}

func ForEachAsync[T any](begin Iterator[T], end Iterator[T], f func(*T)) {
	var wg sync.WaitGroup
	defer wg.Wait() // Barrier
	for ; begin.Index() != end.Index(); begin.Next() {
		wg.Add(1)
		go func(it Iterator[T]) { f(begin.Ptr()); defer wg.Done() }(begin)
	}
}

func Transform[I any, O any](begin Iterator[I], end Iterator[I], reciever Iterator[O], f func(I) O) Iterator[O] {
	for begin.Index() != end.Index() {
		*reciever.Ptr() = f(*begin.Ptr())
		begin.Next()
		reciever.Next()
	}
	return reciever
}

func TransformAsync[I any, O any](begin Iterator[I], end Iterator[I], reciever Iterator[O], f func(I) O) Iterator[O] {
	var wg sync.WaitGroup
	defer wg.Wait() // Barrier
	for begin.Index() != end.Index() {
		wg.Add(1)
		go func(recv *O, send *I) { *recv = f(*send); wg.Done() }(reciever.Ptr(), begin.Ptr())
		begin.Next()
		reciever.Next()
	}
	return reciever
}
