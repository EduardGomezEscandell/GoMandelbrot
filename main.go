package main

import (
	"fmt"

	"github.com/EduardGomezEscandell/GoMandelbrot/image"
	"github.com/EduardGomezEscandell/GoMandelbrot/maths"
)

func main() {

	center := maths.Complex{Real: -0.6, Imag: 0.0}
	width := 2.0 * 1.77777777778
	height := 2.0

	xspan := [2]float64{center.Real - width/2, center.Real + width/2}
	yspan := [2]float64{center.Imag - height/2, center.Imag + height/2}
	maxiter := 100

	colormap, err := image.ColormapFactory("grayscale", 0, maxiter)
	if err != nil {
		panic(err)
	}
	colormap.Invert()

	img := image.NewImage("mandelbrot.ppm", 1920, 1080)
	img.Title = fmt.Sprintf("Mandelbrot set, centered around %f%+fi, width %f, height %f", center.Real, center.Imag, width, height)

	for row := img.RowsBegin(); row != img.RowsEnd(); row.Next() {

		go func(row image.Range) {
			for it := row.Begin(); it != row.End(); it.Next() {
				c := maths.PixelToCoordinate(&it, img.Width, img.Height, xspan, yspan)
				niters := maths.MandelbrotDivergenceIter(c, maxiter)
				it.Set(colormap.Eval(niters))
			}
		}(row)

	}

	image.ImagePPMOutput(&img, "mandelbrot.ppm")
}
