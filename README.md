# Nice Life

A simple Conway's Game of Life implementation

## About

I needed some visible Golang code for contract work so I wrote this.  Now
it's just a fun toy project for testing AI operations on matrixes.

## Installation

There's no binary distributions of this package, so you'll have to build 
from source.  If you don't have a Golang dev env set up this probably isn't
worth the effort.

0. Set up a working Go environment: [official guide](https://golang.org/doc/code.html)
0. Install libsdl2-dev (or equivalent) via your system's package manager
0. `go get github.com/jrforrest/lifegame`
0. `$GOPATH/

## Usage

### System Deps

- libsdl2

### Running

- `go get github.com/jrforrest/lifegame`
- `$GOPATH/bin/lifegame`

### Controls

- f: Fullscreen
- wasd: scrolling
- jk: Zooming
- click: Spawn cell

## Known Bugs

- Zooming in then out too far breaks the simulation

## License

WTFPL - See LICENSE
