package imageIO

import (
	"github.com/EduardGomezEscandell/GoMandelbrot/colors"
	"github.com/EduardGomezEscandell/GoMandelbrot/frames"
)

func IntToColor(frame *frames.IntFrame, cmap colors.Colormap) frames.Image {
	image := frames.NewImage(frame.Width(), frame.Height())
	frames.TransformAsync(frame.Begin(), frame.End(), image.Begin(), cmap.Eval)
	return image
}
