package main

import "github.com/jrforrest/go-nicelife/simthread"
import . "github.com/jrforrest/go-nicelife/gui"

import "math/rand"

func initSim(sim *simthread.SimThread) {
	for i := 0; i < 1000; i++ {
		sim.SpawnCell(rand.Intn(100), rand.Intn(100))
	}
}

func main() {
	sim := simthread.Run()
	gui := NewGui(sim)

	go initSim(sim)

	gui.Start()
}
