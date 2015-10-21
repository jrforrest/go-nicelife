package main

import "lifegame/simthread"
import . "lifegame/gui"

func initSim(sim *simthread.SimThread) {
	sim.SpawnCell(1, 1)
	sim.SpawnCell(0, 1)
	sim.SpawnCell(2, 1)
	sim.SpawnCell(2, 0)
	sim.SpawnCell(1, 0)
}

func main() {
	sim := simthread.Run()
	gui := NewGui(sim)

	go initSim(sim)

	gui.Start()
}
