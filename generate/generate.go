package generate

import (
	"github.com/EduardGomezEscandell/GoMandelbrot/image"
	"github.com/EduardGomezEscandell/GoMandelbrot/maths"
)

type GenerationData struct {
	Img     image.Image
	Cmap    image.Colormap
	Maxiter int
	Xspan   [2]float64
	Yspan   [2]float64
}

func (self *GenerationData) DefineComplexFrame(center maths.Complex, real_span float64, imag_span float64) {
	self.Xspan = [2]float64{center.Real - real_span/2, center.Real + real_span/2}
	self.Yspan = [2]float64{center.Imag - imag_span/2, center.Imag + imag_span/2}
}

func generate_row(row image.Range, gdata *GenerationData) {
	for it := row.Begin(); it != row.End(); it.Next() {
		c := maths.PixelToCoordinate(&it, gdata.Img.Width, gdata.Img.Height, gdata.Xspan, gdata.Yspan)
		niters := maths.MandelbrotDivergenceIter(c, gdata.Maxiter)
		it.Set(gdata.Cmap.Eval(niters))
	}
}

func (data *GenerationData) GenerateSequential() {
	for row := data.Img.RowsBegin(); row != data.Img.RowsEnd(); row.Next() {
		generate_row(row, data)
	}
}

func (data *GenerationData) GenerateConcurrent() {
	for row := data.Img.RowsBegin(); row != data.Img.RowsEnd(); row.Next() {
		go generate_row(row, data)
	}
}
