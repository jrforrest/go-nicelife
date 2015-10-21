package simthread

import "testing"

import . "lifegame/pos"

func TestSimThreadRun(t *testing.T) {
	thread := Run()

	thread.MoveIn <- Position{X: 1, Y: 1}
	state := <-thread.StateOut

	if len(state) < 1 {
		t.Errorf("Expected a move to be made!")
	}

	thread.HaltIn <- true
}
