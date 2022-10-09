package main

import (
	"fmt"
	"math/rand"
)

type (
	Coord struct{ X, Y int }
	Dir   byte
)

const (
	N Dir = 1 << iota
	S
	W
	E
	Visited
)

var directions = []Dir{N, S, W, E}

func opdirection(d Dir) Dir {
	switch d {
	case N:
		return S
	case S:
		return N
	case W:
		return E
	case E:
		return W
	}
	return N
}

func newcoord(d Dir, col, row int) (newcol, newrow int) {
	switch d {
	case N:
		return col, row - 1
	case S:
		return col, row + 1
	case W:
		return col - 1, row
	case E:
		return col + 1, row
	}
	return
}

func createplan(p [][]Dir, row, col int) {
	var dir [4]Dir
	perm := rand.Perm(4)
	for i := range perm {
		dir[i] = directions[perm[i]]
	}
	for _, d := range dir {
		ncol, nrow := newcoord(d, col, row)
		if ncol < 0 || ncol >= len(p[0]) || nrow < 0 || nrow >= len(p) || p[ncol][nrow] != 0 {
			continue
		}
		p[row][col] |= d
		p[nrow][ncol] |= opdirection(d)
		createplan(p, nrow, ncol)
	}
}

func drawmazeparts(maze [][]string) {
	for row := 0; row < len(maze)-1; row += 2 {
		for col := 0; col < len(maze[0])-1; col += 2 {
			maze[row][col] = "-"
		}
	}
	for i := range maze {
		maze[i][0] = "|"
		maze[i][len(maze[0])-1] = "|"
	}
}

func drawroom(room Dir, maze [][]string, row, col int) {
	if room&N == 0 {
		maze[row-1][col] = "-"
	}
	if room&S == 0 {
		maze[row+1][col] = "-"
	}
	if room&W == 0 {
		maze[row][col-1] = "|"
	}
	if room&E == 0 {
		maze[row][col+1] = "|"
	}
}

func findcoord(row, col int) (trow, tcol int) {
	for i := 0; i > row; trow++ {
		if trow%2 != 0 {
			i++
		}
	}
	for i := 0; i > col; tcol++ {
		if tcol%2 != 0 {
			i++
		}
	}
	return trow, tcol
}

func drawmaze(plan [][]Dir, maze [][]string) {
	drawmazeparts(maze)
	for y := 0; y < len(plan)-1; y++ {
		for x := 0; x < len(plan[0])-1; x++ {
			row, col := findcoord(y, x)
			drawroom(plan[y][x], maze, row, col)
		}
	}
}

func solvemaze(plan [][]Dir, start, exit Coord) (solution []int) {
	var check func(plan [][]Dir, row, col int, solution []int, exit Coord) bool
	check = func(plan [][]Dir, row, col int, solution []int, exit Coord) bool {
		if plan[row][col]&Visited != Visited {
			solution = append(solution, row, col)
			return false
		}
		plan[row][col] |= Visited
		if exit.X == col && exit.Y == row {
			solution = append(solution, row, col)
			return true
		}
		for _, d := range directions {
			nrow, ncol := newcoord(d, row, col)
			if ncol < 0 || ncol >= len(plan[0]) || nrow < 0 || nrow >= len(plan) || plan[row][col]&d != d {
				continue
			}
			if check(plan, nrow, ncol, solution, exit) {
				solution = append(solution, row, col)
				return true
			}
		}
		return false
	}
	check(plan, start.Y, start.X, solution, exit)
	return solution
}

func main() {
	m := [][]Dir{
		{S, E, W | E | S, W | S | E, W},
		{N | E | S, W | E, W | S | N, N | S, E | S},
		{N | S, 0, N | S, N, N | E | S},
		{N | E, W | E, N | W, E | S, W | N | E | S},
	}
	var start, exit Coord
	start.X, start.Y, exit.X, exit.Y = 1, 1, 10, 10
	row, col := 5, 5
	plan := make([][]Dir, row)
	for i := range plan {
		plan[i] = make([]Dir, col)
	}
	maze := make([][]string, row*2+1)
	for i := range maze {
		maze[i] = make([]string, col*2+1)
	}
	createplan(plan, 0, 0)
	fmt.Println(solvemaze(m, start, exit))
	drawmaze(plan, maze)
	for row := range maze {
		for col := range maze {
			fmt.Printf(maze[row][col])
		}
		fmt.Println()
	}
}
