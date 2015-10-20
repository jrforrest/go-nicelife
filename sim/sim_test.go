package sim

import "testing"
import . "lifegame/pos"
import . "lifegame/cell"

func TestNeighbors(t *testing.T) {
	sim := NewSimulation()

	neighborCell := sim.addCell(Position{X: 1, Y: 0})
	originCell := sim.addCell(Position{X: 0, Y: 0})
	sim.commit()

	neighbors := sim.cellNeighbors(originCell)

	if !cellsInclude(neighbors, neighborCell) {
		t.Errorf("Neighbor cells for %v should include %v\n"+
			"got %v",
			originCell, neighborCell, neighbors)
	}
}

func cellsInclude(cells []Cell, cell Cell) bool {
	for _, cell := range cells {
		if cell == cell {
			return true
		}
	}
	return false
}
