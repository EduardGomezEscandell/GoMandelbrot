package imageIO

import (
	"github.com/EduardGomezEscandell/GoMandelbrot/colors"
	"github.com/EduardGomezEscandell/GoMandelbrot/frames"
)

// Sequential version, for debugging
func int_to_color_sequential(frame *frames.IntFrame, cmap colors.Colormap) frames.Image {
	image := frames.NewImage(frame.Width(), frame.Height())
	frames.Transform(frame.Begin(), frame.End(), image.Begin(), cmap.Eval)
	return image
}

// Concurrent version, for production
func int_to_color_concurrent(frame *frames.IntFrame, cmap colors.Colormap) frames.Image {
	image := frames.NewImage(frame.Width(), frame.Height())
	frames.TransformAsync(frame.Begin(), frame.End(), image.Begin(), func(iters int) colors.Color {
		return cmap.Eval(iters)
	})
	return image
}

var IntToColor = int_to_color_concurrent
