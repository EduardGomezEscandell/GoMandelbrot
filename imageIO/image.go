package imageIO

import (
	"fmt"
	"os"
)

type SequentialImage struct {
	// private
	file          *os.File
	size          int
	curr_position int

	// public
	Width  int
	Height int
}

func NewImage(filename string, width int, height int) SequentialImage {
	image := SequentialImage{}
	image.Width = width
	image.Height = height
	image.size = width * height

	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	image.file = file

	// Header
	fmt.Fprintf(file, "P3\n")
	fmt.Fprintf(file, "%d %d\n", image.Width, image.Height)
	fmt.Fprintf(file, "255\n")

	return image
}

func (img *SequentialImage) Comment(comment string) {
	fmt.Fprintf(img.file, "# %s\n", comment)
}

func (img *SequentialImage) Close() {
	for img.curr_position < img.size {
		img.NextPixel(Color{})
	}
	img.file.Close()
}

func (img *SequentialImage) CloseWithColor(color Color) {
	for img.curr_position < img.size {
		img.NextPixel(color)
	}
	img.file.Close()
}

func (img *SequentialImage) NextPixel(color Color) error {
	_, err := fmt.Fprintf(img.file, "%d %d %d\n", color.R, color.G, color.B)
	img.curr_position++
	return err
}

// Returs (row, column)
func (img *SequentialImage) Position() (int, int) {
	return (img.curr_position / img.Width), (img.curr_position % img.Width)
}

func (img *SequentialImage) Index() int {
	return img.curr_position
}

func (img *SequentialImage) Size() int {
	return img.size
}
