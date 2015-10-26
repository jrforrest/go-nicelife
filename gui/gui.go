package gui

import . "github.com/jrforrest/go-nicelife/simthread"
import . "github.com/jrforrest/go-nicelife/pos"
import . "github.com/jrforrest/go-nicelife/gui/cmd"
import "github.com/veandco/go-sdl2/sdl"

type Gui struct {
	simThread  *SimThread
	sdlWindow  *sdl.Window
	sdlSurface *sdl.Surface
	cam        *camera
	fullscreen bool
	// Retains the last rendered cells so the display may be re-drawn between
	// receiving game state updates
	lastPositions []Position
	chCmd         chan Cmd //Commands from input handler to main loop
}

func NewGui(simThread *SimThread) *Gui {
	cam := newCamera(800, 600)

	return &Gui{
		simThread: simThread,
		cam:       cam,
		chCmd:     make(chan Cmd),
	}
}

func (gui *Gui) Start() {
	window, err := sdl.CreateWindow(
		"nice life",
		sdl.WINDOWPOS_UNDEFINED,
		sdl.WINDOWPOS_UNDEFINED,
		gui.cam.width,
		gui.cam.height,
		sdl.WINDOW_SHOWN|sdl.WINDOW_RESIZABLE)

	if err != nil {
		panic(err)
	}

	surface, err := window.GetSurface()
	if err != nil {
		panic(err)
	}

	gui.sdlWindow = window
	gui.sdlSurface = surface

	go gui.handleInput()

	gui.mainLoop()
	sdl.Quit()
}

func (gui *Gui) mainLoop() {
	for true {
		select {
		case state := <-gui.simThread.StateOut:
			gui.RenderSim(state)
		case cmd := <-gui.chCmd:
			gui.handleCmd(cmd)
		}
	}
}

func (gui *Gui) handleCmd(cmd Cmd) {
	switch cmd := cmd.(type) {
	case ToggleFullscreen:
		gui.toggleFullScreen()
	case MoveCamera:
		gui.cam.move(cmd.Direction, 10)
	case Zoom:
		gui.cam.zoom(cmd.Direction, 10)
	}
	gui.updateDisplay()
}

func (gui *Gui) updateDisplay() {
	gui.RenderSim(gui.lastPositions)
}

// Renders the current state of the simulation
func (gui *Gui) RenderSim(positions []Position) {
	gui.renderBackgroundGrid()
	gui.renderLiveCells(positions)
	gui.sdlWindow.UpdateSurface()
	gui.lastPositions = positions
}

func (gui *Gui) renderLiveCells(positions []Position) {
	for _, pos := range gui.cam.visiblePositions(positions) {
		gui.drawRelativeCellRect(pos.X, pos.Y, randomCellColor())
	}
}

func (gui *Gui) toggleFullScreen() {
	if !gui.fullscreen {
		gui.sdlWindow.SetFullscreen(sdl.WINDOW_FULLSCREEN)
		gui.fullscreen = true
	} else {
		gui.fullscreen = false
		gui.sdlWindow.SetFullscreen(sdl.WINDOW_RESIZABLE)
	}
}

func (gui *Gui) renderBackgroundGrid() {
	odd := true
	var color uint32

	for x := 0; x <= gui.cam.nCellsHoriz; x++ {
		for y := 0; y <= gui.cam.nCellsVert; y++ {
			if odd {
				color = 0xff111111
				odd = false
			} else {
				color = 0xff141414
				odd = true
			}

			gui.drawCellRect(x, y, color)
		}
	}
}

// Draws the cell rect relative to the camera position
func (gui *Gui) drawRelativeCellRect(x int, y int, color uint32) {
	width, height := gui.cam.cellSizes()
	gui.drawRect(
		int32(x*width-gui.cam.x),
		int32(y*height-gui.cam.y),
		int32(width),
		int32(height),
		color)
}

func (gui *Gui) drawCellRect(x int, y int, color uint32) {
	width, height := gui.cam.cellSizes()
	gui.drawRect(
		int32(x*width),
		int32(y*height),
		int32(width),
		int32(height),
		color)
}

func (gui *Gui) drawRect(x int32, y int32, width int32, height int32, color uint32) {
	rect := sdl.Rect{x, y, width, height}
	gui.sdlSurface.FillRect(&rect, color)
}
