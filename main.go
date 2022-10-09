package main

import (
	"math/rand"
	"time"
)

type Dir byte

const (
	N Dir = 1 << iota
	S
	W
	E
)

var (
	dirUpdX = map[Dir]int{N: 0, S: 0, W: -1, E: 1}
	dirUpdY = map[Dir]int{N: -1, S: 1, W: 0, E: 0}
	dirOpos = map[Dir]Dir{N: S, S: N, W: E, E: W}
)

func genExits(e [][]Dir, x, y int) {
	dirs := []Dir{N, S, W, E}
	for i := range dirs {
		n := rand.Intn(len(dirs) - i)
		dirs[i], dirs[i+n] = dirs[i+n], dirs[i]
	}
	for _, d := range dirs {
		nx, ny := x+dirUpdX[d], y+dirUpdY[d]
		if ny >= 0 && ny < len(e) && nx >= 0 && nx < len(e[0]) && e[ny][nx] == 0 {
			e[y][x] |= d
			e[ny][nx] |= dirOpos[d]
			genExits(e, nx, ny)
		}
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	const rows, cols = 10, 10
	exits := make([][]Dir, rows)
	for i := range exits {
		exits[i] = make([]Dir, cols)
	}
	genExits(exits, 0, 0)
}
