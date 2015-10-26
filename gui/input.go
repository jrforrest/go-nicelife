package gui

import . "github.com/jrforrest/go-nicelife/gui/cmd"
import "github.com/veandco/go-sdl2/sdl"

func (gui *Gui) handleInput() {
	for true {
		ev := sdl.WaitEvent()
		switch ev := ev.(type) {
		case *sdl.MouseButtonEvent:
			if ev.Type == sdl.MOUSEBUTTONDOWN {
				gui.handleClick(ev)
			}
		case *sdl.KeyDownEvent:
			gui.handleKeyDown(ev)
		}
	}
}

func (gui *Gui) handleClick(ev *sdl.MouseButtonEvent) {
	horiz, vert := gui.cam.cellSizes()
	relX, relY := gui.cam.relativeCoords(int(ev.X), int(ev.Y))

	cellX := relX / horiz
	cellY := relY / vert
	gui.simThread.SpawnCell(cellX, cellY)
}

func (gui *Gui) handleKeyDown(ev *sdl.KeyDownEvent) {
	sym := ev.Keysym.Sym

	switch sdl.GetKeyName(sym) {
	case "F":
		gui.chCmd <- ToggleFullscreen{}
	case "W":
		gui.chCmd <- MoveCamera{Direction: UP}
	case "A":
		gui.chCmd <- MoveCamera{Direction: LEFT}
	case "S":
		gui.chCmd <- MoveCamera{Direction: DOWN}
	case "D":
		gui.chCmd <- MoveCamera{Direction: RIGHT}
	case "J":
		gui.chCmd <- Zoom{Direction: IN}
	case "K":
		gui.chCmd <- Zoom{Direction: OUT}
	}
}
