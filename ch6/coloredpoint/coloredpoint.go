package coloredpoint

import (
	"RichKernigan/ch6/geometry"
	"image/color"
)

// ColoredPoint represents point and its color
type ColoredPoint struct {
	*geometry.Point
	Color color.RGBA
}

// Distance is the wrapper method for the Point distance method
func (p ColoredPoint) Distance(q geometry.Point) float64 {
	return p.Point.Distance(q)
}

// ScaleBy wrapps the Point's type ScaleBy method
func (p *ColoredPoint) ScaleBy(factor float64) {
	p.Point.ScaleBy(factor)
}
