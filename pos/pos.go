package pos

type Position struct {
	X int
	Y int
}

// Returns all of the cells neighboring the given cell
func (pos Position) Neighbors() []Position {
	positions := make([]Position, 0, 8)
	for x := pos.X - 1; x <= pos.X+1; x++ {
		for y := pos.Y - 1; y <= pos.Y+1; y++ {
			// skip the given position itself
			if (x != pos.X) || (y != pos.Y) {
				positions = append(positions, Position{X: x, Y: y})
			}
		}
	}

	return positions
}
