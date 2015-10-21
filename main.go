package main

import . "lifegame/sim"
import . "lifegame/gui"

func main() {
	sim := NewSimulation()
	sim.SpawnCell(1, 1)
	sim.SpawnCell(1, 0)
	sim.SpawnCell(1, 2)
	sim.SpawnCell(2, 2)
	sim.SpawnCell(4, 2)
	sim.SpawnCell(5, 1)
	sim.SpawnCell(6, 3)
	sim.SpawnCell(7, 3)
	sim.SpawnCell(5, 2)
	gui := NewGui(&sim)

	gui.Start()
}
