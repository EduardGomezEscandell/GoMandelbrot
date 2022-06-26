package imageIO

import (
	"fmt"
	"os"

	"github.com/EduardGomezEscandell/GoMandelbrot/image"
)

const (
	BINARY = iota
	ASCII
)

func ppm_header(filehandle *os.File, img *image.Image, encoding int) {

	if encoding == ASCII {
		fmt.Fprintf(filehandle, "P3\n")
	} else if encoding == BINARY {
		fmt.Fprintf(filehandle, "P6\n")
	} else {
		panic("Unknown encoding")
	}

	fmt.Fprintf(filehandle, "# %s\n", img.Title)
	fmt.Fprintf(filehandle, "%d %d\n", img.Width, img.Height)
	fmt.Fprintf(filehandle, "255\n")
}

func ppm_body(filehandle *os.File, img *image.Image, encoding int) {
	if encoding == ASCII {
		for _, px := range img.Pixels {
			fmt.Fprintf(filehandle, "%d %d %d ", px.R, px.G, px.B)
		}
	} else if encoding == BINARY {
		for _, px := range img.Pixels {
			filehandle.Write([]byte{px.R, px.G, px.B})
		}
	} else {
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
