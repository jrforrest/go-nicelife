package gui

import "testing"

func TestQuadWithin(t *testing.T) {
	testQuad := newQuad(0, 0, 10, 10)

	expectPointInQuad(t, true, newVec(5, 5), testQuad)
	expectPointInQuad(t, true, newVec(0, 9), testQuad)
	expectPointInQuad(t, false, newVec(-1, 5), testQuad)
	expectPointInQuad(t, false, newVec(1, 15), testQuad)
}

func expectPointInQuad(t *testing.T, in bool, point vec, q quad) {
	if in {
		if !q.pointWithin(point) {
			t.Errorf("Expected point: %v to be in quad: %v", point, q)
		}
	} else {
		if q.pointWithin(point) {
			t.Errorf("Expected point: %v to not be in quad: %v", point, q)
		}
	}
}
