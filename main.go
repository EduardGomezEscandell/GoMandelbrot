package main

import (
	"fmt"

	"github.com/EduardGomezEscandell/GoMandelbrot/generate"
	"github.com/EduardGomezEscandell/GoMandelbrot/image"
	"github.com/EduardGomezEscandell/GoMandelbrot/maths"
)

func main() {

	max_iter := 1000

	gdata := generate.GenerationData{
		Img:     image.NewImage("mandelbrot.ppm", 1080, 720),
		Maxiter: max_iter,
		Cmap:    image.ColormapFactory("grayscale", 0, 100),
	}
	gdata.Cmap.Invert()

	center := maths.Complex{Real: -0.9, Imag: 0.25}
	real_span := 0.1875
	imag_span := 0.125
	gdata.DefineComplexFrame(center, real_span, imag_span)

	gdata.Img.Title = fmt.Sprintf(
		"Mandelbrot set, centered around %f%+fi, width %f, height %f",
		center.Real, center.Imag, real_span, imag_span)

	gdata.GenerateConcurrent()

	println("Image generated. Storing...")
	image.ImagePPMOutput(&gdata.Img, "mandelbrot.ppm")

	println("Done")
}
