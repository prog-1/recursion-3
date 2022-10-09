package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Dir represents a bit mask direction for north (1), south (2), west (4) and east (8).
type Dir byte

const (
	N Dir = 1 << 0 // 1
	S Dir = 1 << 1 // 2
	W Dir = 1 << 2 // 4
	E Dir = 1 << 3 // 8
)

func returnNewDirection(row, col int, d Dir) (nrow, ncol int) {
	if d == N {
		return row - 1, col
	}
	if d == S {
		return row + 1, col
	}
	if d == W {
		return row, col - 1
	}
	if d == E {
		return row, col + 1
	}
	panic("incorrect direction")
}

func returnOppositeDirection(d Dir) Dir {
	if d == N {
		return S
	}
	if d == S {
		return N
	}
	if d == W {
		return E
	}
	if d == E {
		return W
	}
	panic("incorrect direction")
}

func genMaze(e [][]Dir, row, col int) {
	dirs := []Dir{N, S, W, E} // shuffle
	rand.Seed(time.Now().UnixNano())
	for i := range dirs {
		n := rand.Intn(len(dirs) - i)
		dirs[i], dirs[i+n] = dirs[i+n], dirs[i]
	}

	for _, d := range dirs {
		nrow, ncol := returnNewDirection(row, col, d)
		if nrow < 0 || nrow >= len(e) || ncol < 0 || ncol >= len(e[0]) {
			continue
		}
		if e[nrow][ncol] != 0 {
			continue
		}
		e[row][col] |= d
		e[nrow][ncol] |= returnOppositeDirection(d)
		genMaze(e, nrow, ncol)
	}
}

type (
	// Coord represents an (X, Y) coordinate on the map.
	Coord struct{ X, Y int }
)

// solveMaze finds the first path found between start and end coordinates. The
// function returns coordinates representing the path. If the path cannot be
// found, the function returns nil.
func solveMaze(e [][]Dir, visited [][]int, path []Coord, start, end Coord) { // Will finish later
	if start == end {
		return
	}
	dirs := []Dir{N, S, W, E}
	for _, d := range dirs {
		nrow, ncol := returnNewDirection(start.Y, start.X, d)
		if nrow < 0 || nrow >= len(e) || ncol < 0 || ncol >= len(e[0]) {
			continue
		}
		if e[nrow][ncol] == 0 || e[nrow][ncol]&d != d || visited[nrow][ncol] == 1 {
			continue
		}
		visited[nrow][ncol] = 1
		start.X, start.Y = ncol, nrow
		path = append(path, start)
		solveMaze(e, visited, path, start, end)
	}
	path = append(path, Coord{0, 0})
	for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1 {
		path[i], path[j] = path[j], path[i]
	}
}

func printOneLine(square []string, e [][]Dir, row, col int) {
	for i := range square {
		for j := 0; j < col; j++ {
			if e[row][j]&N == N && i == 0 {
				fmt.Print("+  ")
				continue
			}
			if e[row][j]&W == W && i == 1 {
				fmt.Print("   ")
				continue
			}
			fmt.Print(square[i])
		}
		if i == 0 {
			fmt.Print("+")
		} else {
			fmt.Print("|")
		}
		fmt.Println()
	}
}

func main() {
	exits := make([][]Dir, 8)
	for i := range exits {
		exits[i] = make([]Dir, 11)
	}
	genMaze(exits, 0, 0)

	square := []string{"+--", "|  "}
	for row, col := 0, 11; row < 8; row++ {
		printOneLine(square, exits, row, col)
		if row == 7 {
			for cnt := 0; cnt < col; cnt++ { // cnt < 11
				for _, i := range square[0] {
					fmt.Print(string(i))
				}
			}
			fmt.Print("+\n")
		}
	}

	visited := make([][]int, 8)
	for i := range visited {
		visited[i] = make([]int, 11)
	}
	var path []Coord
	solveMaze(exits, visited, path, Coord{0, 0}, Coord{0, 2})
	fmt.Println(path)
}
