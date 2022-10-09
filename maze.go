package main

import (
	"fmt"
	"io"
	"math/rand"
	"os"
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
	dirUpd = map[Dir]func(col int, row int) (ncol, nrow int){
		N: func(c int, r int) (int, int) { return c, r - 1 },
		S: func(c int, r int) (int, int) { return c, r + 1 },
		W: func(c int, r int) (int, int) { return c - 1, r },
		E: func(c int, r int) (int, int) { return c + 1, r },
	}
	dirOpos = map[Dir]Dir{N: S, S: N, W: E, E: W}
)

func HasDir(e [][]Dir, row, col int, d Dir) bool {
	if len(e) == 0 {
		return false
	}
	if row < 0 || row >= len(e) || col < 0 || col >= len(e[0]) {
		return false
	}
	return e[row][col]&d == d
}

func updateTop(last []rune, row []Dir) {
	hasDir := func(col int, d Dir) bool {
		if col < 0 || col >= len(row) {
			return false
		}
		return row[col]&d == d
	}
	noDir := func(col int, d Dir) bool { return !hasDir(col, d) }
	forCornerVert := func(corner rune, d Dir, hasVert bool) rune {
		switch corner {
		case '└':
			return '├'
		case '┴':
			return '┼'
		case '┘':
			return '┤'
		case '┌':
			if d == E {
				return '┬'
			}
		case '┐':
			if d == W && !hasVert {
				return '┬'
			}
		case '─':
			if d == W && !hasVert {
				return '┌'
			} else if d == E && !hasVert {
				return '┐' // ┌
			}
		case ' ':
			if hasVert {
				return '│'
			}
		}
		return corner
	}

	forCornerHorz := func(corner rune, d Dir, hasHor bool) rune {
		switch corner {
		case ' ':
			if hasHor {
				return '─'
			} else if d == W {
				return '┌'
			} else if d == E {
				return '┐'
			}
		}
		return corner
	}
	for col := range row {
		if noDir(col, W) {
			last[col*4] = forCornerVert(last[col*4], W, hasDir(col, N))
		}
		if noDir(col, E) {
			last[col*4+4] = forCornerVert(last[col*4+4], E, hasDir(col, N))
		}
		if noDir(col, N) {
			last[col*4] = forCornerHorz(last[col*4], W, hasDir(col, W))
			for i := 1; i < 4; i++ {
				last[col*4+i] = '─'
			}
			last[col*4+4] = forCornerHorz(last[col*4+4], E, hasDir(col, E))
		}
	}
}

func updateMid(last []rune, row []Dir) {
	hasDir := func(col int, d Dir) bool {
		if col < 0 || col >= len(row) {
			return false
		}
		return row[col]&d == d
	}
	noDir := func(col int, d Dir) bool { return !hasDir(col, d) }
	for col := range row {
		if noDir(col, W) {
			last[col*4] = '│'
		}
		if noDir(col, E) {
			last[col*4+4] = '│'
		}
	}
}

func updateBottom(last []rune, row []Dir) {
	hasDir := func(col int, d Dir) bool {
		if col < 0 || col >= len(row) {
			return false
		}
		return row[col]&d == d
	}
	noDir := func(col int, d Dir) bool { return !hasDir(col, d) }
	for col := range row {
		if noDir(col, W) {
			if noDir(col, S) {
				if last[col*4] == '┘' {
					last[col*4] = '┴'
				} else {
					last[col*4] = '└'
				}
				for i := 1; i < 4; i++ {
					last[col*4+i] = '─'
				}
			} else {
				if col > 0 && last[col*4-1] == '─' {
					last[col*4] = '┘'
				} else {
					last[col*4] = '│'
				}
			}
		} else if noDir(col, S) {
			for i := 0; i < 5; i++ {
				last[col*4+i] = '─'
			}
		}
		if noDir(col, E) {
			if noDir(col, S) {
				last[col*4+4] = '┘'
			} else {
				last[col*4+4] = '│'
			}
		}
	}
}

func printExits(w io.Writer, e [][]Dir) {
	lastRow := make([]rune, 4*len(e[0])+1)
	resetLast := func() {
		for i := range lastRow {
			lastRow[i] = ' '
		}
	}
	outputLast := func() { fmt.Fprintln(w, string(lastRow)) }
	resetLast()
	for row := range e {
		updateTop(lastRow, e[row])
		outputLast()
		resetLast()
		updateMid(lastRow, e[row])
		outputLast()
		resetLast()
		updateBottom(lastRow, e[row])
	}
	outputLast()
}

func genExits(x [][]Dir, col, row int) {
	dirs := []Dir{N, S, E, W}
	for i := range dirs {
		n := rand.Intn(len(dirs) - i)
		dirs[i], dirs[i+n] = dirs[i+n], dirs[i]
	}

	for _, d := range dirs {
		ncol, nrow := dirUpd[d](col, row)
		if nrow < 0 || nrow >= len(x) || ncol < 0 || ncol >= len(x[0]) {
			continue
		}
		if x[nrow][ncol] != 0 {
			continue
		}
		x[nrow][ncol] |= dirOpos[d]
		x[row][col] |= d

		genExits(x, ncol, nrow)

	}

}

func main() {
	rand.Seed(time.Now().UnixNano())
	const rows, cols = 12, 12
	exits := make([][]Dir, rows)
	for i := range exits {
		exits[i] = make([]Dir, cols)
	}
	genExits(exits, 0, 0)
	printExits(os.Stdout, exits)
}
