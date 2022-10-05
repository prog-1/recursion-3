// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
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

/*
┌───┬───┐
│   │   │
├───┼───┤
│   │   │
└───┴───┘
*/
type (
	// Coord represents an (X, Y) coordinate on the map.
	Coord struct{ X, Y int }
)

// AnyPath find the first path found between start and end coordinates. The
// functions returns coordinates representing the path. If the path cannot be
// found, the function returns nil.
var dirUpd = map[Dir]func(col int, row int) (ncol, nrow int){
	N: func(c int, r int) (int, int) { return c, r - 1 },
	S: func(c int, r int) (int, int) { return c, r + 1 },
	W: func(c int, r int) (int, int) { return c - 1, r },
	E: func(c int, r int) (int, int) { return c + 1, r },
}

func cycle(e [][]Dir, end Coord, way []Coord) []Coord {
	var dirs [4]Dir
	perm := rand.Perm(4)
	// fmt.Println(way)
	for i := range perm {
		dirs[i] = allDirs[perm[i]]
	}
	for _, d := range dirs {
		// fmt.Println(d, "d")
		// fmt.Println(e[way[len(way)-1].X][way[len(way)-1].Y])
		if e[way[len(way)-1].Y][way[len(way)-1].X]&d == 0 {
			continue
		}
		ncol, nrow := dirUpd[d](way[len(way)-1].X, way[len(way)-1].Y)
		if len(way) > 1 {
			if ncol == way[len(way)-2].X && nrow == way[len(way)-2].Y {
				continue
			}
		}

		way = append(way, Coord{ncol, nrow})
		if way[len(way)-1].X == end.X && way[len(way)-1].Y == end.Y {
			return way
		}
		return cycle(e, end, way)

	}

	way = way[:len(way)-1]
	return cycle(e, end, way)
}
func AnyPath(e [][]Dir, start, end Coord) []Coord {
	var way []Coord
	way = append(way, start)
	if start.X == end.X && start.Y == end.Y {
		return nil
	}
	return cycle(e, end, way)
}

func HasDir(e [][]Dir, row, col int, d Dir) bool {
	if len(e) == 0 {
		return false
	}
	if row < 0 || row >= len(e) || col < 0 || col >= len(e[0]) {
		return false
	}
	return e[row][col]&d == d
}

var (
	allDirs = []Dir{N, W, E, S}
	// dirUpd  = map[Dir]func(col int, row int) (ncol, nrow int){
	// 	N: func(c int, r int) (int, int) { return c, r - 1 },
	// 	S: func(c int, r int) (int, int) { return c, r + 1 },
	// 	W: func(c int, r int) (int, int) { return c - 1, r },
	// 	E: func(c int, r int) (int, int) { return c + 1, r },
	// }
	dirOpos = map[Dir]Dir{N: S, S: N, W: E, E: W}
)

func genExits(e [][]Dir, col, row int) {
	var dirs [4]Dir
	perm := rand.Perm(4)
	fmt.Println(perm)
	for i := range perm {
		dirs[i] = allDirs[perm[i]]
	}

	for _, d := range dirs {
		ncol, nrow := dirUpd[d](col, row)
		if ncol < 0 || ncol >= len(e[0]) || nrow < 0 || nrow >= len(e) {
			continue
		}
		if e[nrow][ncol] != 0 {
			continue
		}
		e[row][col] |= d
		e[nrow][ncol] |= dirOpos[d]

		genExits(e, ncol, nrow)

	}

}

func output(e [][]Dir) {
	for i := range e {
		for j := range e[i] {
			if 
		}

}
}
func main() {
	rand.Seed(time.Now().UnixNano())

	const rows, cols = 5, 5
	exits := make([][]Dir, rows)
	for i := range exits {
		exits[i] = make([]Dir, cols)
	}
	genExits(exits, 0, 0)
	output(exits)
	for i := range exits {
		for j := range exits[i] {
			fmt.Printf("%3d", exits[i][j])
		}
		fmt.Println()
	}
	fmt.Println(AnyPath(exits, Coord{0, 0}, Coord{len(exits) - 1, len(exits[0]) - 1}))
}
