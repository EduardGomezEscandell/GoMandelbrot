package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/EduardGomezEscandell/GoMandelbrot/colors"
	"github.com/EduardGomezEscandell/GoMandelbrot/generate"
	"github.com/EduardGomezEscandell/GoMandelbrot/imageIO"
	"github.com/EduardGomezEscandell/GoMandelbrot/maths"
)

func Log(verbose bool, text string) {
	if verbose {
		fmt.Printf("%s\n", text)
	}
}

func TimedLog(verbose bool, start time.Time, text string) {
	if verbose {
		fmt.Printf("[%f\"] %s\n", time.Since(start).Seconds(), text)
	}
}

func parseAndAssignDefaults() generate.Config {

	// Defaults
	max_iter := 1000

	image_width := 1920
	image_height := 1080
	subsampling_factor := 1

	colormap := "grayscale"
	colormap_lb := 0
	colormap_ub := -1
	colormap_nl_inverted := 0.0 // Nonlinearity [0, +inf)

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
	flag.Float64Var(&colormap_nl_inverted, "cnl", colormap_nl_inverted, "Colormap nonlinearity (0 is none, up to +inf)")

	flag.StringVar(&zcenter, "zc", zcenter, "Center of the complex plane")
	flag.Float64Var(&zspan, "zs", zspan, "Horizontal span of the complex plane")

	flag.BoolVar(&verbose, "v", verbose, "Verbose mode")
	flag.StringVar(&output_filename, "o", output_filename, "Output filename")

	flag.IntVar(&subsampling_factor, "sf", subsampling_factor, "Subsampling factor (1 is no subsampling)")

	flag.Parse()

	// Post-processing
	if colormap_ub < 0 {
		colormap_ub = max_iter
	}
	color_nonlinearity := 1.0 / (colormap_nl_inverted + 1.0e-10)

	center := maths.ParseComplex(zcenter)
	aspect_ratio := float64(image_height) / float64(image_width)
	real_span := zspan
	imag_span := zspan * aspect_ratio

	// Printing info
	Log(verbose, fmt.Sprintf("Generating image with the following settings:"))
	Log(verbose, fmt.Sprintf("  Max iterations: %d", max_iter))
	Log(verbose, fmt.Sprintf("  Resolution: %d x %d", image_width, image_height))
	Log(verbose, fmt.Sprintf("  Subsampling factor: %d", subsampling_factor))
	Log(verbose, fmt.Sprintf("  Colormap: %s [%d, %d] nl=%g", colormap, colormap_lb, colormap_ub, colormap_nl_inverted))
	Log(verbose, fmt.Sprintf("  Complex plane center: %g%+gi", center.Real, center.Imag))
	Log(verbose, fmt.Sprintf("  Complex plane span:   %g%+gi", real_span, imag_span))
	Log(verbose, fmt.Sprintf("  Output filename:   %s", output_filename))

	// Sanity tests
	if image_width < 1 || image_height < 1 {
		fmt.Printf("Error: image resolution must be at least 1x1\n")
	}
	if subsampling_factor < 1 {
		fmt.Printf("Error: subsampling factor must be at least 1\n")
	}
	if colormap_nl_inverted < 0 {
		fmt.Printf("Error: colormap nonlinearity must be nonnegative\n")
	}

	// Packing into GenerationData
	gdata := generate.Config{
		Width:            uint(image_width),
		Height:           uint(image_height),
		Maxiter:          max_iter,
		Cmap:             colors.ColormapFactory(colormap, colormap_lb, colormap_ub, color_nonlinearity),
		OutputFilename:   output_filename,
		Verbosity:        verbose,
		SubscalingFactor: uint(subsampling_factor),
	}
	gdata.DefineComplexFrame(center, real_span, imag_span)
	gdata.MetaData = fmt.Sprintf(
		"Mandelbrot set, centered around %f%+fi, width %f, height %f",
		center.Real, center.Imag, real_span, imag_span)

	return gdata
}

func main() {
	start := time.Now()

	config := parseAndAssignDefaults()
	IOformat, err := imageIO.GetFileFormat(config.OutputFilename)
	if err != nil {
		panic(err) // Failing early
	}
	TimedLog(config.Verbosity, start, "Input read. Generating map...")

	frame := generate.GenerateConcurrent(&config)
	TimedLog(config.Verbosity, start, "Map generated. Coloring...")

	image := imageIO.IntToColor(&frame, config.Cmap)
	TimedLog(config.Verbosity, start, "Coloring done. Saving...")

	imageIO.Save(&image, config.OutputFilename, config.MetaData, IOformat)
	TimedLog(config.Verbosity, start, "Done")
}
