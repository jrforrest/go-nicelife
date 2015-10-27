package gui

type quad struct {
	origin vec
	dims   vec
}

type vec struct {
	x int
	y int
}

func newQuad(x int, y int, width int, height int) quad {
	return quad{
		origin: newVec(x, y),
		dims:   newVec(width, height),
	}
}

func newVec(x int, y int) vec {
	return vec{x: x, y: y}
}

func (quad quad) pointWithin(point vec) bool {
	return (point.x >= quad.origin.x) &&
		(point.x < quad.origin.x+quad.dims.x) &&
		(point.y >= quad.origin.y) &&
		(point.y < quad.origin.y+quad.dims.y)
}
