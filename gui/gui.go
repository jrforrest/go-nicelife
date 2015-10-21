package gui

import . "lifegame/sim"
import "github.com/veandco/go-sdl2/sdl"

import "fmt"

type Gui struct {
	sim         *Simulation
	sdlWindow   *sdl.Window
	sdlSurface  *sdl.Surface
	pxHeight    int
	pxWidth     int
	nCellsHoriz int
	nCellsVert  int
}

func NewGui(sim *Simulation) *Gui {
	return &Gui{
		sim:         sim,
		pxHeight:    600,
		pxWidth:     800,
		nCellsHoriz: 10,
		nCellsVert:  10,
	}
}

func (gui *Gui) Start() {
	window, err := sdl.CreateWindow(
		"lifegame",
		sdl.WINDOWPOS_UNDEFINED,
		sdl.WINDOWPOS_UNDEFINED,
		800,
		600,
		sdl.WINDOW_SHOWN)

	if err != nil {
		panic(err)
	}

	surface, err := window.GetSurface()
	if err != nil {
		panic(err)
	}

	gui.sdlWindow = window
	gui.sdlSurface = surface

	gui.mainLoop()
	sdl.Quit()
}

func (gui *Gui) mainLoop() {
	running := true
	for running {
		gui.sim.Step()
		sdl.Delay(1000)
		gui.RenderSim()
	}
}

// Renders the current state of the simulation
func (gui *Gui) RenderSim() {
	gui.renderBackgroundGrid()
	gui.renderLiveCells()
	gui.sdlWindow.UpdateSurface()
}

func (gui *Gui) renderLiveCells() {
	fmt.Printf("Drawing cells: %v\n", gui.sim.LiveCellPositions())
	for _, pos := range gui.sim.LiveCellPositions() {
		if (pos.X >= 0 && pos.X < gui.nCellsHoriz) && (pos.Y >= 0 && pos.Y < gui.nCellsVert) {
			gui.drawCellRect(pos.X, pos.Y, 0xff0000ff)
		}
	}
}

func (gui *Gui) renderBackgroundGrid() {
	for x := 0; x <= gui.nCellsHoriz; x++ {
		for y := 0; y <= gui.nCellsVert; y++ {
			gui.drawCellRect(x, y, 0xff00ff00)
		}
	}
}

func (gui *Gui) drawCellRect(x int, y int, color uint32) {
	width, height := gui.cellSizes()
	gui.drawRect(
		int32(x*(width+4)),
		int32(y*(height+4)),
		int32(width),
		int32(height),
		color)
}

func (gui *Gui) drawRect(x int32, y int32, width int32, height int32, color uint32) {
	rect := sdl.Rect{x, y, width, height}
	gui.sdlSurface.FillRect(&rect, color)
}

// Cell sizes in pixels for horiz and vert rendering
func (gui *Gui) cellSizes() (int, int) {
	horiz := (gui.pxWidth - 4*gui.nCellsHoriz) / gui.nCellsHoriz
	vert := (gui.pxHeight - 4*gui.nCellsVert) / gui.nCellsVert
	return horiz, vert
}
