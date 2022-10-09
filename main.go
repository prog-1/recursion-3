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

func cycle(e [][]Dir, end Coord, way []Coord, visited map[Coord]bool) []Coord {
	var dirs [4]Dir
	perm := rand.Perm(4)
	for i := range perm {
		dirs[i] = allDirs[perm[i]]
	}
	for _, d := range dirs {
		if e[way[len(way)-1].Y][way[len(way)-1].X]&d == 0 {
			continue
		}
		ncol, nrow := dirUpd[d](way[len(way)-1].X, way[len(way)-1].Y)
		if ncol < 0 || ncol >= len(e[0]) || nrow < 0 || nrow >= len(e) {
			continue
		}
		if v := visited[Coord{ncol, nrow}]; v {
			continue
		} else {
			visited[Coord{ncol, nrow}] = true
		}

		way = append(way, Coord{ncol, nrow})
		if way[len(way)-1].X == end.X && way[len(way)-1].Y == end.Y {
			return way
		}
		return cycle(e, end, way, visited)

	}

	way = way[:len(way)-1]
	return cycle(e, end, way, visited)
}
func AnyPath(e [][]Dir, start, end Coord) []Coord {
	var way []Coord
	way = append(way, start)
	if start.X == end.X && start.Y == end.Y {
		return nil
	}
	visited := make(map[Coord]bool)
	visited[Coord{0, 0}] = true
	return cycle(e, end, way, visited)
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
	dirOpos = map[Dir]Dir{N: S, S: N, W: E, E: W}
)

func genExits(e [][]Dir, col, row int) {
	var dirs [4]Dir
	perm := rand.Perm(4)
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

func output(e [][]Dir, m map[Coord]tag) {
	var str string
	for i := 0; i <= len(e)*6+2; i++ {
		str += "#"
	}
	fmt.Println(str)
	var tmph, tmpv string
	for i, v := range e {
		tmph += "###"
		tmpv += "###"
		for j, v1 := range v {

			if _, ok := m[Coord{j, i}]; ok {
				tmph += " . "
			} else {
				tmph += "   "
			}

			if v1&2 != 0 {
				tmpv += "   "
			} else {
				tmpv += "###"
			}

			if v1&8 != 0 {
				if _, ok := m[Coord{j, i}]; ok {
					tmph += " . "
				} else {
					tmph += "   "
				}

			} else {
				tmph += "###"
			}
			tmpv += "###"

		}
		fmt.Println(tmph)
		fmt.Println(tmpv)
		tmph = ""
		tmpv = ""
	}

}

type tag struct{}

func main() {
	rand.Seed(time.Now().UnixNano())
	const rows, cols = 20, 20
	exits := make([][]Dir, rows)
	for i := range exits {
		exits[i] = make([]Dir, cols)
	}
	genExits(exits, 0, 0)
	s := AnyPath(exits, Coord{0, 0}, Coord{len(exits) - 1, len(exits[0]) - 1})

	m := make(map[Coord]tag)
	for _, v := range s {
		m[v] = tag{}
	}
	output(exits, m)
}
