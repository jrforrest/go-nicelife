package simthread

import . "lifegame/sim"
import . "lifegame/pos"
import "time"

type SimThread struct {
	sim      *Simulation
	HaltIn   chan bool
	MoveIn   chan Position
	StateOut chan []Position
}

func Run() *SimThread {
	sim := NewSimulation()
	thread := SimThread{
		sim:      &sim,
		HaltIn:   make(chan bool),
		MoveIn:   make(chan Position),
		StateOut: make(chan []Position),
	}
	go thread.start()

	return &thread
}

func (thread *SimThread) start() {

	for true {
		select {
		case move := <-thread.MoveIn:
			thread.sim.SpawnCell(move.X, move.Y)
		case <-time.After(1 * time.Second):
			thread.sim.Step()
		case <-thread.HaltIn:
			break
		}
		thread.StateOut <- thread.sim.LiveCellPositions()
	}
	close(thread.StateOut)
}
