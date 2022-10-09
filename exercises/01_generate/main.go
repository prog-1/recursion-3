package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

type Dir rune

const (
	N Dir = 1 << iota
	S
	W
	E
)

func printExits(e [][]Dir, x, y int) {
	row := y*2 - 1
	res := strings.Repeat("_", row)
	fmt.Println(" " + res)
	for rows := 0; rows < x; rows++ {
		fmt.Print("|")
		for cols := 0; cols < y; cols++ {
			if e[rows][cols]&S != 0 {
				fmt.Print(" ")
			} else {
				fmt.Print("_")
			}
			if e[rows][cols]&E != 0 {
				if (e[rows][cols+1])|(e[rows][cols])&S != 0 {
					fmt.Print(" ")
				} else {
					fmt.Print("_")
				}

			} else {
				fmt.Print("|")
			}
		}
		fmt.Println()
	}
}

var (
	allDirs = []Dir{N, W, E, S}
)

func dir2(d Dir, nrow, ncol int) (int, int) {
	switch d {
	case N:
		return nrow - 1, ncol
	case W:
		return nrow, ncol - 1
	case E:
		return nrow, ncol + 1
	case S:
		return nrow + 1, ncol
	}
	return nrow, ncol
}

func diropos(d Dir) Dir {
	switch d {
	case N:
		return S
	case W:
		return E
	case E:
		return W
	case S:
		return N
	}
	return 0
}

func genExits(e [][]Dir, row, col int) {
	dirs := allDirs
	for i := range dirs {
		n := rand.Intn(len(dirs) - i)
		dirs[i], dirs[i+n] = dirs[i+n], dirs[i]
	}
	for _, d := range dirs {
		nrow, ncol := dir2(d, row, col)
		if ncol < 0 || ncol >= len(e[0]) || nrow < 0 || nrow >= len(e) {
			continue
		}
		if e[nrow][ncol] != 0 {
			continue
		}
		e[row][col] |= d
		e[nrow][ncol] |= diropos(d)
		genExits(e, nrow, ncol)
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	const rows, cols = 8, 14
	exits := make([][]Dir, rows)
	for i := range exits {
		exits[i] = make([]Dir, cols)
	}
	genExits(exits, 0, 0)
	printExits(exits, rows, cols)

}
