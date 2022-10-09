package main

import (
	"fmt"
	"testing"
)

func TestGenMaze(t *testing.T) { // ? it always returns 'ok', even when it shouldn't
	for _, tc := range []struct {
		row, col int
	}{
		{0, 0},
		{1, 1},
		{5, 5},
		{8, 8},
		{8, 9},
		{8, 11},
		{11, 8},
		{11, 11},
	} {
		t.Run(fmt.Sprint(tc.row)+fmt.Sprint(tc.col), func(t *testing.T) {
			exits := make([][]Dir, tc.row)
			for i := range exits {
				exits[i] = make([]Dir, tc.col)
			}
			genMaze(exits, tc.row, tc.col)
			for i := range exits {
				for j := range exits[i] {
					if exits[i][j]&N == N {
						if i == 0 || exits[i][j]&exits[i-1][j] != N {
							t.Errorf("created maze is incorrect")
						}
					}
					if exits[i][j]&S == S {
						if i == len(exits)-1 || exits[i][j]&exits[i+1][j] != S {
							t.Errorf("created maze is incorrect")
						}
					}
					if exits[i][j]&W == W {
						if j == 0 || exits[i][j]&exits[i][j-1] != W {
							t.Errorf("created maze is incorrect")
						}
					}
					if exits[i][j]&E == E {
						if j == len(exits[0])-1 || exits[i][j]&exits[i][j+1] != E {
							t.Errorf("created maze is incorrect")
						}
					}
				}
			}
		})
	}
}
