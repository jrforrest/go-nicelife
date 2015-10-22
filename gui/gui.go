package gui

import . "lifegame/simthread"
import . "lifegame/pos"
import "github.com/veandco/go-sdl2/sdl"

type Cmd int

const (
	TOGGLE_FULLSCREEN Cmd = iota
	EXIT
)

type Gui struct {
	simThread   *SimThread
	sdlWindow   *sdl.Window
	sdlSurface  *sdl.Surface
	pxHeight    int
	pxWidth     int
	nCellsHoriz int
	nCellsVert  int
	fullscreen  bool
	chCmd       chan Cmd //Commands from input handler to main loop
}

func NewGui(simThread *SimThread) *Gui {
	return &Gui{
		simThread:   simThread,
		pxHeight:    600,
		pxWidth:     800,
		nCellsHoriz: 100,
		nCellsVert:  100,
		chCmd:       make(chan Cmd),
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

	go gui.watchClick()

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
	switch cmd {
	case TOGGLE_FULLSCREEN:
		gui.toggleFullScreen()
	}
}

// Renders the current state of the simulation
func (gui *Gui) RenderSim(positions []Position) {
	gui.renderBackground()
	gui.renderBackgroundGrid()
	gui.renderLiveCells(positions)
	gui.sdlWindow.UpdateSurface()
}

func (gui *Gui) renderLiveCells(positions []Position) {
	for _, pos := range positions {
		if (pos.X >= 0 && pos.X < gui.nCellsHoriz) && (pos.Y >= 0 && pos.Y < gui.nCellsVert) {
			gui.drawCellRect(pos.X, pos.Y, randomCellColor())
		}
	}
}

func (gui *Gui) renderBackground() {
	gui.drawRect(0, 0, int32(gui.pxWidth), int32(gui.pxHeight), 0xff333333)
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

	for x := 0; x <= gui.nCellsHoriz; x++ {
		for y := 0; y <= gui.nCellsVert; y++ {
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

func (gui *Gui) drawCellRect(x int, y int, color uint32) {
	width, height := gui.cellSizes()
	gui.drawRect(
		int32(x*(width)),
		int32(y*(height)),
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
	horiz := gui.pxWidth / gui.nCellsHoriz
	vert := gui.pxHeight / gui.nCellsVert
	return horiz, vert
}

func (gui *Gui) watchClick() {
	for true {
		horiz, vert := gui.cellSizes()

		ev := sdl.WaitEvent()
		switch ev := ev.(type) {
		case *sdl.MouseButtonEvent:
			if ev.Type == sdl.MOUSEBUTTONDOWN {
				cellX := int(ev.X) / horiz
				cellY := int(ev.Y) / vert
				gui.simThread.SpawnCell(cellX, cellY)
			}
		case *sdl.KeyDownEvent:
			sym := ev.Keysym.Sym
			name := sdl.GetKeyName(sym)
			if ev.Type == sdl.KEYDOWN && name == "F" {
				gui.chCmd <- TOGGLE_FULLSCREEN
			}
		}
	}
}
