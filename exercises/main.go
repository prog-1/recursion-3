package main

import "fmt"

const (
	N = 1 << iota
	E
	S
	W
)

func fill(dirs [][]byte) [][]byte {
	n, m := len(dirs), len(dirs[0])

	// Creating a map, which will be displayed on terminal
	maze := make([][]byte, n*3)
	for i := range maze {
		maze[i] = make([]byte, m*3)
	}

	// Filling map with "rooms"
	for i, j := 0, 1; i < len(dirs); i, j = i+1, j+3 {
		for k, w := 0, 1; k < len(dirs[i]); k, w = k+1, w+3 {
			maze[j][w] = ' '   // Middle
			maze[j-1][w] = '-' // North
			maze[j][w+1] = '|' // East
			maze[j+1][w] = '-' // South
			maze[j][w-1] = '|' // East

			maze[j-1][w-1] = '+' // NorthWest
			maze[j-1][w+1] = '+' // NorthEast
			maze[j+1][w+1] = '+' // SouthEash
			maze[j+1][w-1] = '+' // SouthWest
		}
	}

	return maze
}

func print(maze [][]byte) {
	for i := range maze {
		for j := range maze[i] {
			fmt.Print(string(maze[i][j]))
		}
		fmt.Println()
	}
}

func main() {
	var n, m uint64 = 5, 20
	dirs := make([][]byte, n)
	for i := range dirs {
		dirs[i] = make([]byte, m)
	}

	maze := fill(dirs)
	print(maze)
}
