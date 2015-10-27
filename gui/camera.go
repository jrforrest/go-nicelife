package gui

import . "github.com/jrforrest/go-nicelife/pos"
import . "github.com/jrforrest/go-nicelife/gui/cmd"

type camera struct {
	x           int
	y           int
	width       int
	height      int
	nCellsHoriz int
	nCellsVert  int
}

func newCamera(width int, height int) *camera {
	cam := camera{
		x:           0,
		y:           0,
		width:       width,
		height:      height,
		nCellsHoriz: 100,
		nCellsVert:  100,
	}

	return &cam
}

func (cam *camera) move(dir Direction, dist int) {
	cellWidth, cellHeight := cam.cellSizes()

	switch dir {
	case UP:
		cam.y -= dist * cellHeight
	case DOWN:
		cam.y += dist * cellHeight
	case LEFT:
		cam.x -= dist * cellWidth
	case RIGHT:
		cam.x += dist * cellWidth
	}
}

// dist is just the number of cells to add to the scene in both directions
func (cam *camera) zoom(dir Direction, dist int) {
	switch dir {
	case IN:
	case OUT:
		dist = -dist
	default:
		panic("Unknown zoom direction!")
	}

	newHoriz := cam.nCellsHoriz + dist
	newVert := cam.nCellsVert + dist

	if newHoriz > 0 && cam.width/newHoriz >= 1 {
		if newVert > 0 && cam.height/newVert >= 1 {
			cam.nCellsHoriz = newHoriz
			cam.nCellsVert = newVert
		}
	}
}

// A slice of the given positions that are within the camera bounds
func (cam *camera) visiblePositions(positions []Position) []Position {
	visible := make([]Position, 0, cam.nCellsHoriz*cam.nCellsVert)
	for _, pos := range positions {
		if (pos.X >= 0 && pos.X < cam.nCellsHoriz) && (pos.Y >= 0 && pos.Y < cam.nCellsVert) {
			visible = append(visible, pos)
		}
	}
	return visible
}

// Cell sizes in pixels for horiz and vert rendering
func (cam *camera) cellSizes() (int, int) {
	horiz := cam.width / cam.nCellsHoriz
	vert := cam.height / cam.nCellsVert
	return horiz, vert
}

func (cam *camera) relativeCoords(x int, y int) (int, int) {
	return x + cam.x, y + cam.y
}
