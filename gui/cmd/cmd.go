// Commands for async gui operations
package cmd

type Cmd interface{}

type ToggleFullscreen struct{}
type MoveCamera struct {
	Direction Direction
}

type Direction int

const (
	UP Direction = iota
	DOWN
	LEFT
	RIGHT
)
