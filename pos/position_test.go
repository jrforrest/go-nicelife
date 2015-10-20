package pos

import "testing"

var originPosition = Position{X: 0, Y: 0}

func TestNeighbors(t *testing.T) {
	neighbors := originPosition.Neighbors()

	if len(neighbors) != 8 {
		t.Errorf("Position neighbors should contain 8 elems")
	}
	if positionsInclude(neighbors, originPosition) {
		t.Errorf("Neighbors of a position should not include that position itself")
	}
}

func positionsInclude(positions []Position, pos Position) bool {
	for _, elem := range positions {
		if elem == pos {
			return true
		}
	}
	return false
}
