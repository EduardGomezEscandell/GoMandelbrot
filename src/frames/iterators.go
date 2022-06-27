package frames

// Frame iterator accessors
func (self *Frame[T]) Begin() Iterator[T] {
	return Iterator[T]{parent: self, idx: 0}
}

func (self *Frame[T]) End() Iterator[T] {
	return Iterator[T]{parent: self, idx: self.Width() * self.Height()}
}

// Rows
type Row[T any] struct {
	begin Iterator[T]
	end   Iterator[T]
	idx   uint
}

func (self *Row[T]) Begin() Iterator[T] {
	return self.begin
}

func (self *Row[T]) End() Iterator[T] {
	return self.end
}

func (self *Row[T]) Index() uint {
	return self.idx
}

// Cells
type Iterator[T any] struct {
	parent *Frame[T]
	idx    uint
}

func (self *Iterator[T]) Next() {
	self.idx = self.idx + 1
}

func (self *Iterator[T]) Ptr() *T {
	return (*self.parent).at(self.idx)
}

func (self *Iterator[T]) Index() uint {
	return self.idx
}
