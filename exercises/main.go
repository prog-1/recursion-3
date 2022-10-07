package main

import "fmt"

// 0 - wall    1 - way
const (
	N = 1 << iota
	E
	S
	W
)

func print(maze [][]byte) {
	// Creating char array, which will be displayed on screen:
	output := make([][]byte, len(maze)*3)
	for i := range output {
		output[i] = make([]byte, len(maze[0])*3)
	}

	// Filling output array:
	for i, k := 0, 1; i < len(maze); i, k = i+1, k+3 {
		for j, w := 0, 1; j < len(maze[i]); j, w = j+1, w+3 {
			output[k][w] = ' '     // Middle
			output[k-1][w-1] = '+' // NorthWest
			output[k-1][w+1] = '+' // NorthEast
			output[k+1][w+1] = '+' // SouthEash
			output[k+1][w-1] = '+' // SouthWest

			cell := maze[i][j]
			if cell&N == N {
				output[k-1][w] = ' '
			} else {
				output[k-1][w] = '-'
			}
			if cell&E == E {
				output[k][w+1] = ' '
			} else {
				output[k][w+1] = '|'
			}
			if cell&S == S {
				output[k+1][w] = ' '
			} else {
				output[k+1][w] = '-'
			}
			if cell&W == W {
				output[k][w-1] = ' '
			} else {
				output[k][w-1] = '|'
			}

		}
	}

	// Printing output array:
	for i := range output {
		for j := range output[i] {
			fmt.Print(string(output[i][j]))
		}
		fmt.Println()
	}
}

func main() {
	var maze [][]byte = [][]byte{
		{0b0010, 0b1100, 0b0000},
		{0b0010, 0b1011, 0b1000},
	}
	print(maze)
}
