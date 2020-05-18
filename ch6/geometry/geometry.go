package geometry

import "math"

// Point is a structure which represents a point by its x and y coordinates.
type Point struct{ X, Y float64 }

// Add adds p's coordinates to q's
func (p Point) Add(q Point) Point { return Point{p.X + q.X, p.Y + q.Y} }

// Sub subtracts p's coordinates from q's
func (p Point) Sub(q Point) Point { return Point{p.X - q.X, p.Y - q.Y} }

// ScaleBy func receives pointer multiplies coord by factor
func (p *Point) ScaleBy(factor float64) {
	p.X *= factor
	p.Y *= factor
}

// Distance traditional function
func Distance(p, q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

// Distance is same thing , but as  a method of the Point type.
func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

// A Path is a journey connecting the points with straight lines.
type Path []Point

// Distance returns the distance traveled along the path.
func (path Path) Distance() float64 {
	sum := 0.0
	for i := range path {
		if i > 0 {
			sum += path[i-1].Distance(path[i])
		}
	}
	return sum
}

// TranslateBy translates all path' s coordinates to offset value
func (path Path) TranslateBy(offset Point, add bool) {
	var op func(p, q Point) Point
	if add {
		op = Point.Add
	} else {
		op = Point.Sub
	}
	for i := range path {
		path[i] = op(path[i], offset)
	}
}
