package imageIO

import (
	"fmt"
	"os"
	"strings"

	"github.com/EduardGomezEscandell/GoMandelbrot/image"
)

const (
	ppm_NONE = iota
	ppm_BINARY
	ppm_ASCII
)

func ppm_header(filehandle *os.File, img *image.Image, encoding int) {

	switch encoding {
	case ppm_ASCII:
		fmt.Fprintf(filehandle, "P3\n")
	case ppm_BINARY:
		fmt.Fprintf(filehandle, "P6\n")
	default:
		panic("Unknown encoding")
	}

	fmt.Fprintf(filehandle, "# %s\n", img.Title)
	fmt.Fprintf(filehandle, "%d %d\n", img.Width, img.Height)
	fmt.Fprintf(filehandle, "255\n")
}

func ppm_body(filehandle *os.File, img *image.Image, encoding int) {
	switch encoding {
	case ppm_ASCII:
		for _, px := range img.Pixels {
			fmt.Fprintf(filehandle, "%d %d %d ", px.R, px.G, px.B)
		}
	case ppm_BINARY:
		for _, px := range img.Pixels {
			filehandle.Write([]byte{px.R, px.G, px.B})
		}
	default:
		panic("Unknown encoding")
	}
}

func ImagePPMOutput(img *image.Image, filename string, encoding int) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	ppm_header(file, img, encoding)
	ppm_body(file, img, encoding)

	return nil
}

func PPMcheck(filename string) int {
	data := strings.Split(filename, ".")

	if len(data) == 2 && data[1] == "ppm" {
		return ppm_BINARY
	}

	if len(data) == 3 && data[1] == "ascii" && data[2] == "ppm" {
		return ppm_ASCII
	}

	if len(data) == 3 && data[1] == "bin" && data[2] == "ppm" {
		return ppm_BINARY
	}

	return -1
}
