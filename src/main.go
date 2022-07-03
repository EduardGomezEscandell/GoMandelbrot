package main

import (
	"flag"
	"fmt"
	"strconv"
	"strings"
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

func parseAndAssignDefaults() (generate.Config, generate.IterationCounter) {

	// Defaults
	max_iter := 1000

	image_width := 1920
	image_height := 1080
	subsampling_level := 1

	colormap := "grayscale"
	colormap_lb := 0
	colormap_ub := -1
	colormap_nl_inverted := 0.0 // Nonlinearity [0, +inf)

	output_filename := "mandelbrot.ppm"

	zcenter := "-0.6"
	zspan := 4.0

	julia_param := "false"

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

	flag.IntVar(&subsampling_level, "sf", subsampling_level, "Subsampling factor (1 is no subsampling)")

	flag.StringVar(&julia_param, "julia", julia_param, "Complex number: use Julia set. None: use Mandelbrot set")

	flag.Parse()

	// Post-processing
	if colormap_ub < 0 {
		colormap_ub = max_iter
	}
	color_nonlinearity := 1.0 / (colormap_nl_inverted + 1.0e-10)

	center, err := strconv.ParseComplex(zcenter, 64)
	if err != nil {
		panic(fmt.Sprintf("Failed to parse complex number: %s", zcenter))
	}

	isjulia, generator := func() (bool, generate.IterationCounter) {
		if strings.ToLower(julia_param) == "false" {
			return false, maths.MandelbrotDivergencePeriod
		} else {
			c, err := strconv.ParseComplex(julia_param, 64)
			if err != nil {
				panic(fmt.Sprintf("Failed to parse complex number: %s", julia_param))
			}
			return true, func(z complex128, maxiter uint) uint { return maths.JuliaDivergencePeriod(z, c, maxiter) }
		}
	}()

	aspect_ratio := float64(image_height) / float64(image_width)
	real_span := zspan
	imag_span := zspan * aspect_ratio

	// Printing info
	Log(verbose, fmt.Sprintf("Generating image with the following settings:"))
	Log(verbose, fmt.Sprintf("  Max iterations: %d", max_iter))
	Log(verbose, fmt.Sprintf("  Resolution: %d x %d", image_width, image_height))
	Log(verbose, fmt.Sprintf("  Subsampling level: %d", subsampling_level))
	Log(verbose, fmt.Sprintf("  Colormap: %s [%d, %d] nl=%g", colormap, colormap_lb, colormap_ub, colormap_nl_inverted))
	Log(verbose, fmt.Sprintf("  Complex plane center: %G", center))
	Log(verbose, fmt.Sprintf("  Complex plane span:   %G", complex(real_span, imag_span)))
	if isjulia {
		Log(verbose, fmt.Sprintf("  Mathematical object:  Julia set with param=%s", julia_param))
	} else {
		Log(verbose, fmt.Sprintf("  Mathematical object:  Mandelbrot set"))
	}
	Log(verbose, fmt.Sprintf("  Output filename:   %s", output_filename))

	// Sanity tests
	if image_width < 1 || image_height < 1 {
		fmt.Printf("Error: image resolution must be at least 1x1\n")
	}
	if subsampling_level < 1 {
		fmt.Printf("Error: subsampling factor must be at least 1\n")
	}
	if colormap_nl_inverted < 0 {
		fmt.Printf("Error: colormap nonlinearity must be nonnegative\n")
	}
	if max_iter < 0 {
		panic("Error: max iterations must be positive\n")
	}

	// Packing into GenerationData
	gdata := generate.Config{
		Width:           uint(image_width),
		Height:          uint(image_height),
		Maxiter:         uint(max_iter),
		Cmap:            colors.ColormapFactory(colormap, colormap_lb, colormap_ub, color_nonlinearity),
		OutputFilename:  output_filename,
		Verbosity:       verbose,
		SubscalingLevel: uint(subsampling_level),
	}
	gdata.DefineComplexFrame(center, real_span, imag_span)
	gdata.MetaData = fmt.Sprintf(
		"Mandelbrot set, centered around %g, width %f, height %f",
		center, real_span, imag_span)

	return gdata, generator
}

func main() {
	start := time.Now()

	config, generator := parseAndAssignDefaults()
	IOformat, err := imageIO.GetFileFormat(config.OutputFilename)
	if err != nil {
		panic(err) // Failing early
	}
	TimedLog(config.Verbosity, start, "Input read. Generating map...")

	frame := generate.Generate(&config, generator)
	TimedLog(config.Verbosity, start, "Map generated. Coloring...")

	image := imageIO.IntToColor(&frame, config.Cmap)
	TimedLog(config.Verbosity, start, "Coloring done. Saving...")

	imageIO.Save(&image, config.OutputFilename, config.MetaData, IOformat)
	TimedLog(config.Verbosity, start, "Done")
}
