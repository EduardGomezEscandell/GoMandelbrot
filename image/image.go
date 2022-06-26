package image

type Image struct {
	// private
	size   int
	pixels []Color

	// public
	Title  string
	Width  int
	Height int
}

func NewImage(filename string, width int, height int) Image {
	image := Image{}
	image.Width = width
	image.Height = height
	image.size = width * height
	image.pixels = make([]Color, image.size)

	return image
}

type Iterator struct {
	image *Image
	idx   int
}

type Range struct {
	begin Iterator
	end   Iterator
}

func (img *Image) Size() int {
	return img.size
}

func (self *Image) GetRow(row int) Range {
	return Range{
		begin: Iterator{image: self, idx: row * self.Width},
		end:   Iterator{image: self, idx: (row + 1) * self.Width},
	}
}

// Row-wise operations
func (self *Image) RowsBegin() Range {
	return self.GetRow(0)
}

func (self *Image) RowsEnd() Range {
	return self.GetRow(self.Height)
}

func (self *Range) Next() *Range {
	len := self.end.idx - self.begin.idx
	self.begin = self.end
	self.end = Iterator{image: self.begin.image, idx: self.end.idx + len}
	return self
}

// Pixel-wise operations
func (self *Range) Begin() Iterator {
	return self.begin
}

func (self *Range) End() Iterator {
	return self.end
}

func (self *Iterator) Next() *Iterator {
	self.idx = self.idx + 1
	return self
}

func (self *Iterator) Set(color Color) {
	self.image.pixels[self.idx] = color
}

func (self *Iterator) Get() Color {
	return self.image.pixels[self.idx]
}

func (self *Iterator) Position() (int, int) {
	return (self.idx / self.image.Width), (self.idx % self.image.Width)
}

func (self *Iterator) Index() int {
	return self.idx
}
