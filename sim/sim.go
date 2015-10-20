package sim

import "fmt"
import . "lifegame/cell"
import . "lifegame/pos"

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

func (sim Simulation) step() {
}

func (sim Simulation) cellNeighbors(cell Cell) []Cell {
	neighbors := make([]Cell, 0, 8)

	for _, pos := range cell.Neighbors() {
		if cell, found := sim.cells[pos]; found {
			neighbors = append(neighbors, cell)
		}
	}

	return neighbors
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
