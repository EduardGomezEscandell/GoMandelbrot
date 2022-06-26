package image

type Image struct {
	// private
	size   int
	Pixels []Color

	// public
	Title  string
	Width  int
	Height int
}

func NewImage(width int, height int) Image {
	image := Image{}
	image.Width = width
	image.Height = height
	image.size = width * height
	image.Pixels = make([]Color, image.size)

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

func (self *Range) RowIndex() int {
	row, _ := self.begin.Position()
	return row
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
	self.image.Pixels[self.idx] = color
}

func (self *Iterator) Get() Color {
	return self.image.Pixels[self.idx]
}

func (self *Iterator) Position() (int, int) {
	return (self.idx / self.image.Width), (self.idx % self.image.Width)
}

func (self *Iterator) Index() int {
	return self.idx
}
