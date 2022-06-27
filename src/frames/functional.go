package frames

func ForEach[T any](begin Iterator[T], end Iterator[T], f func(*T)) {
	for ; begin != end; begin.Next() {
		f(begin.Ptr())
	}
}

func ForEachAsync[T any](begin Iterator[T], end Iterator[T], f func(*T)) {
	for ; begin != end; begin.Next() {
		go f(begin.Ptr())
	}
}

func Transform[I any, O any](begin Iterator[I], end Iterator[I], reciever Iterator[O], f func(I) O) Iterator[O] {
	for begin != end {
		*reciever.Ptr() = f(*begin.Ptr())
		begin.Next()
		reciever.Next()
	}
	return reciever
}

func TransformAsync[I any, O any](begin Iterator[I], end Iterator[I], reciever Iterator[O], f func(I) O) Iterator[O] {
	for begin.Index() != end.Index() {
		go func(recv *O, send *I) { *recv = f(*send) }(reciever.Ptr(), begin.Ptr())
		begin.Next()
		reciever.Next()
	}
	return reciever
}
