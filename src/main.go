package main

import (
	"fmt"

	"github.com/EduardGomezEscandell/GoMandelbrot/generate"
	"github.com/EduardGomezEscandell/GoMandelbrot/image"
	"github.com/EduardGomezEscandell/GoMandelbrot/imageIO"
	"github.com/EduardGomezEscandell/GoMandelbrot/maths"
)

func main() {

	max_iter := 1000

	gdata := generate.GenerationData{
		Img:     image.NewImage("mandelbrot.ppm", 1920, 1080),
		Maxiter: max_iter,
		Cmap:    image.ColormapFactory("grayscale", 0, 200),
	}
	gdata.Cmap.Invert()

	center := maths.Complex{Real: -0.8033, Imag: 0.178}
	real_span := 0.001 * 1920.0 / 1080.0
	imag_span := 0.001
	gdata.DefineComplexFrame(center, real_span, imag_span)

	gdata.Img.Title = fmt.Sprintf(
		"Mandelbrot set, centered around %f%+fi, width %f, height %f",
		center.Real, center.Imag, real_span, imag_span)

	gdata.GenerateConcurrent()

	println("Image generated. Storing...")
	imageIO.ImagePPMOutput(&gdata.Img, "mandelbrot.ppm", imageIO.BINARY)

	println("Done")
}
