package main

import (
	"github.com/EduardGomezEscandell/mandelbrotgo/imageIO"
	"github.com/EduardGomezEscandell/mandelbrotgo/maths"
)

func main() {
	image := imageIO.NewImage("mandelbrot.ppm", 1080, 720)
	defer image.Close()

	center := maths.Complex{Real: -0.6, Imag: 0.0}
	width := 3.0
	height := 2.0

	xspan := [2]float64{center.Real - width/2, center.Real + width/2}
	yspan := [2]float64{center.Imag - height/2, center.Imag + height/2}
	maxiter := 100

	colormap, err := imageIO.ColormapFactory("grayscale", 0, maxiter)
	if err != nil {
		panic(err)
	}
	colormap.Invert()

	for image.Index() < image.Size() {
		c := maths.PixelToCoordinate(&image, xspan, yspan)
		it := maths.MandelbrotDivergenceIter(c, maxiter)
		image.NextPixel(colormap.Eval(it))
	}
}