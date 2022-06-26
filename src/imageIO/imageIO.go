package imageIO

import (
	"errors"

	"github.com/EduardGomezEscandell/GoMandelbrot/image"
)

const (
	NONE = iota
	PPM
)

type ImageIOFormat struct {
	Format    int
	Subformat int
}

func Save(img *image.Image, filename string, format ImageIOFormat) error {
	switch format.Format {
	case PPM:
		return ImagePPMOutput(img, filename, format.Subformat)
	}
	return errors.New("Unknown file format")
}

func GetFileFormat(filename string) (ImageIOFormat, error) {
	if subformat := PPMcheck(filename); subformat > 0 {
		return ImageIOFormat{Format: PPM, Subformat: subformat}, nil
	}

	return ImageIOFormat{Format: NONE}, errors.New("Unknown file format")

}
