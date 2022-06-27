package frames

import "github.com/EduardGomezEscandell/GoMandelbrot/colors"

type Frame[T any] struct {
	data   []T
	title  string
	width  uint
	height uint
}

func NewFrame[T any](width uint, height uint) Frame[T] {
	image := Frame[T]{}
	image.width = width
	image.height = height
	image.data = make([]T, width*height)

	return image
}

func (self *Frame[T]) Width() uint {
	return self.width
}

func (self *Frame[T]) Height() uint {
	return self.height
}

func (self *Frame[T]) GetRow(row uint) Row[T] {
	return Row[T]{
		begin: Iterator[T]{parent: self, idx: row * self.Width()},
		end:   Iterator[T]{parent: self, idx: (row + 1) * self.Width()},
	}
}

func (self *Frame[T]) Title() *string {
	return &self.title
}

func (self *Frame[T]) at(idx uint) *T {
	return &self.data[idx]
}

func (self *Frame[T]) Position(idx uint) (uint, uint) {
	return (idx / self.Width()), (idx % self.Width())
}

// Aliases
type Image = Frame[colors.Color]
type IntFrame = Frame[int]

var NewImage = NewFrame[colors.Color]
var NewIntFrame = NewFrame[int]
