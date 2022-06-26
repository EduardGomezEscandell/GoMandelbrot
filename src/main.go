package main

import (
	"flag"
	"fmt"

	"github.com/EduardGomezEscandell/GoMandelbrot/generate"
	"github.com/EduardGomezEscandell/GoMandelbrot/image"
	"github.com/EduardGomezEscandell/GoMandelbrot/imageIO"
	"github.com/EduardGomezEscandell/GoMandelbrot/maths"
)

func Log(verbose bool, text string) {
	if verbose {
		println(text)
	}
}

func parseAndAssignDefaults() (generate.GenerationData, bool) {

	// Defaults
	max_iter := 1000

	image_width := 1920
	image_height := 1080
	colormap := "grayscale"
	colormap_lb := 0
	colormap_ub := -1
	output_filename := "mandelbrot.ppm"

	zcenter := "-0.8033+0.178i"
	zspan := 0.0017

	verbose := false

	// Parsing
	flag.IntVar(&max_iter, "maxiter", max_iter, "Maximum number of iterations")
	flag.IntVar(&image_width, "imw", image_width, "Image width")
	flag.IntVar(&image_height, "imh", image_height, "Image height")

	flag.StringVar(&colormap, "c", colormap, "Colormap name")
	flag.IntVar(&colormap_lb, "clb", colormap_lb, "Colormap lower bound")
	flag.IntVar(&colormap_ub, "cub", colormap_ub, "Colormap upper bound")

	flag.StringVar(&zcenter, "zc", zcenter, "Center of the complex plane")
	flag.Float64Var(&zspan, "zs", zspan, "Horizontal span of the complex plane")

	flag.BoolVar(&verbose, "v", verbose, "Verbose mode")
	flag.StringVar(&output_filename, "o", output_filename, "Output filename")

	flag.Parse()

	// Post-processing
	if colormap_lb < 0 {
		colormap_lb = 0
	}
	if colormap_ub < 0 {
		colormap_ub = max_iter
	}

	center := maths.ParseComplex(zcenter)
	aspect_ratio := float64(image_height) / float64(image_width)
	real_span := zspan
	imag_span := zspan * aspect_ratio

	// Printing info
	Log(verbose, fmt.Sprintf("Generating image with the following settings:"))
	Log(verbose, fmt.Sprintf("  Max iterations: %d", max_iter))
	Log(verbose, fmt.Sprintf("  Resolution: %d x %d", image_width, image_height))
	Log(verbose, fmt.Sprintf("  Colormap: %s [%d, %d]", colormap, colormap_lb, colormap_ub))
	Log(verbose, fmt.Sprintf("  Complex plane center: %f %+fi", center.Real, center.Imag))
	Log(verbose, fmt.Sprintf("  Complex plane span:   %f %+fi", real_span, imag_span))
	Log(verbose, fmt.Sprintf("  Output filename:   %s", output_filename))

	// Packing into GenerationData
	gdata := generate.GenerationData{
		Img:            image.NewImage(image_width, image_height),
		Maxiter:        max_iter,
		Cmap:           image.ColormapFactory(colormap, colormap_lb, colormap_ub),
		OutputFilename: output_filename,
	}
	gdata.DefineComplexFrame(center, real_span, imag_span)

	gdata.Img.Title = fmt.Sprintf(
		"Mandelbrot set, centered around %f%+fi, width %f, height %f",
		center.Real, center.Imag, real_span, imag_span)

	return gdata, verbose
}

func main() {
	gdata, verbose := parseAndAssignDefaults()

	IOformat, err := imageIO.GetFileFormat(gdata.OutputFilename)
	if err != nil {
		panic(err) // Failing early
	}

	gdata.GenerateConcurrent()
	Log(verbose, "Image generated. Storing...")

	imageIO.Save(&gdata.Img, gdata.OutputFilename, IOformat)
	Log(verbose, "Done")
}
