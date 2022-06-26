package image

import (
	"fmt"
	"os"
)

func ImagePPMOutput(img *Image, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	// Header
	fmt.Fprintf(file, "P3\n")
	fmt.Fprintf(file, "%d %d\n", img.Width, img.Height)
	fmt.Fprintf(file, "255\n")

	for _, px := range img.pixels {
		fmt.Fprintf(file, "%d %d %d\n", px.R, px.G, px.B)
	}

	return nil
}
