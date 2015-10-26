package sim

import . "github.com/jrforrest/go-nicelife/cell"
import . "github.com/jrforrest/go-nicelife/pos"

type Simulation struct {
	cells    map[Position]Cell // The cells in the current sim state
	newCells map[Position]Cell // The cells in the next iteration of the sim
}

func NewSimulation() Simulation {
	return Simulation{
		cells:    make(map[Position]Cell),
		newCells: make(map[Position]Cell),
	}
}

// Steps the simulation to the next frame
func (sim *Simulation) Step() {
	for _, cell := range sim.cells {
		if !sim.cellLives(cell) {
			sim.removeCell(cell.Position)
		}
	}

	for _, candidate := range sim.birthCandidates() {
		if sim.posFertile(candidate) {
			sim.addCell(candidate)
		}
	}

	sim.commit()
}

func (sim *Simulation) LiveCellPositions() []Position {
	positions := make([]Position, 0, 1024)
	for pos, _ := range sim.cells {
		positions = append(positions, pos)
	}
	return positions
}

func (sim *Simulation) SpawnCell(x int, y int) {
	pos := Position{X: x, Y: y}
	sim.addCell(pos)
	sim.commit()
}

// All of the cells that could potentially spawn a new cell
// (have a live cell as a neighbor)
func (sim *Simulation) birthCandidates() []Position {
	posSet := make(map[Position]bool)

	for pos, _ := range sim.cells {
		for _, neighbor := range pos.Neighbors() {
			if _, found := sim.cells[neighbor]; !found {
				posSet[neighbor] = true
			}
		}
	}

	candidates := make([]Position, 0, 1024)
	for pos, _ := range posSet {
		candidates = append(candidates, pos)
	}
	return candidates
}

// Should the given cell get kilt?
func (sim *Simulation) cellLives(cell Cell) bool {
	if n := len(sim.cellNeighbors(cell)); n < 2 {
		return false
	} else if n <= 3 {
		return true
	} else {
		return false
	}
}

// Should a new cell spawn in the given position
func (sim *Simulation) posFertile(pos Position) bool {
	testCell := Cell{Position: pos}
	return len(sim.cellNeighbors(testCell)) == 3
}

func (sim *Simulation) cellNeighbors(cell Cell) []Cell {
	neighbors := make([]Cell, 0, 8)

	for _, pos := range cell.Neighbors() {
		if cell, found := sim.cells[pos]; found {
			neighbors = append(neighbors, cell)
		}
	}

	return neighbors
}

func (sim Simulation) removeCell(pos Position) {
	delete(sim.newCells, pos)
}

// Is there a cell at the given position?
func (sim Simulation) cellPresent(pos Position) bool {
	_, present := sim.cells[pos]
	return present
}

func (sim Simulation) addCell(pos Position) Cell {
	cell := Cell{Position: pos}
	sim.newCells[pos] = cell
	return cell
}

func (sim *Simulation) commit() {
	sim.cells = sim.newCells
	sim.newCells = make(map[Position]Cell)
	for pos, cell := range sim.cells {
		sim.newCells[pos] = cell
	}
}
