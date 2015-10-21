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

func TestBirthCandidates(t *testing.T) {
	sim := NewSimulation()

	sim.addCell(Position{X: 0, Y: 0})
	candidatePos := Position{X: 1, Y: 0}
	sim.commit()

	candidates := sim.birthCandidates()
	if !positionsInclude(candidates, candidatePos) {
		t.Errorf("Candidate cells did not include: %v\n", candidatePos)
	}
}

func TestPosFertile(t *testing.T) {
	sim := newFertileSim()
	if !sim.posFertile(Position{X: 1, Y: 1}) {
		t.Errorf("Expected 1,1 to be fertile!  Cells: %v\n", sim.cells)
	}
}

func TestStep(t *testing.T) {
	sim := newFertileSim()
	sim.Step()

	if !sim.cellPresent(Position{X: 1, Y: 1}) {
		t.Errorf("Expected there to be a cell at 1,1!  Cells: %v\n", sim.cells)
	}
}

func positionsInclude(positions []Position, pos Position) bool {
	for _, element := range positions {
		if element == pos {
			return true
		}
	}

	return false
}

// TODO: name shadowing problem with cell
func cellsInclude(cells []Cell, cell Cell) bool {
	for _, elem := range cells {
		if elem == cell {
			return true
		}
	}
	return false
}

// Returns a sim with a fertile space
func newFertileSim() *Simulation {
	sim := NewSimulation()
	sim.addCell(Position{X: 0, Y: 0})
	sim.addCell(Position{X: 0, Y: 1})
	sim.addCell(Position{X: 2, Y: 1})
	sim.commit()
	return &sim
}
