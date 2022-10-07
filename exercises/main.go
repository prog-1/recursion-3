package main

import (
	"fmt"
	"math/rand"
	"time"
)

// NOTE: 0 - wall    1 - way
const (
	N = 1 << iota
	E
	S
	W
)

// Returns randomly generated maze of given size(n - height, m - width)
// Note that maze is represented as 2D array of bytes
func generateMaze(n, m uint) [][]byte {
	// Creating 2D byte array, which represents a maze:
	maze := make([][]byte, n)
	for i := range maze {
		maze[i] = make([]byte, m)
	}

	// Create array to store whether some cell was visited or not:
	visited := make([][]bool, n)
	for i := range visited {
		visited[i] = make([]bool, m)
	}

	gen(maze, visited, 0, 0, N)

	return maze
}

// Auxiliary function for generateMaze()
func gen(maze [][]byte, visited [][]bool, i, j int, entry byte) {

	visited[i][j] = true // Marking current cell as visited

	ways := getAvailableWays(maze[i][j])

	// Removing entry way:
	var k int
	for k = 0; ways[k] != entry; k++ { // Getting index of entry way in slice of available ways
	}
	maze[i][j] |= ways[k]     // Marking way as visited
	ways = removeWay(ways, k) // Removing entry way out of slice of available ways

	for len(ways) != 0 {
		// Getting random available direction:
		r := rand.Intn(len(ways)) // Getting index of random available direction
		way := ways[r]
		ways = removeWay(ways, r)

		// Try to proceed with next cell at direction:
		in, jn := getIndexByWay(way, i, j)
		// Checking if we haven't went out of maze's bounds or if we've visited it already:
		if in >= 0 && in < len(maze) && jn >= 0 && jn < len(maze[0]) && !visited[in][jn] {
			maze[i][j] |= way // Marking way as visited
			gen(maze, visited, in, jn, getEntry(way))
		}

	}
}

// Returns slice of all available ways for the cell
func getAvailableWays(cell byte) []byte {
	var ways []byte
	if cell&N != N {
		ways = append(ways, N)
	}
	if cell&E != E {
		ways = append(ways, E)
	}
	if cell&S != S {
		ways = append(ways, S)
	}
	if cell&W != W {
		ways = append(ways, W)
	}
	return ways
}

// Removes way under the specified index
func removeWay(ways []byte, i int) []byte {
	if len(ways) == 1 { // Single element
		return nil
	} else if i == len(ways)-1 { // Last element
		return ways[:i]
	} else if i == 0 { // First element
		return ways[i+1:]
	} else { // Not last, not first
		return append(ways[:i], ways[i+1])
	}
}

// Returns index of the next neighbor at the specified way(direction)
func getIndexByWay(way byte, i, j int) (int, int) {
	switch way {
	case N:
		return i - 1, j
	case E:
		return i, j + 1
	case S:
		return i + 1, j
	case W:
		return i, j - 1
	default: // Should never happen
		return 0, 0
	}
}

// Returns entry way for the next cell according to the exit way of the current cell
func getEntry(exit byte) byte {
	switch exit {
	case N:
		return S
	case E:
		return W
	case S:
		return N
	case W:
		return E
	default: // Should never happen
		return 0
	}
}

// Prints maze in terminal window
func printMaze(maze [][]byte) {
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
	rand.Seed(time.Now().UnixNano())
	var n, m uint = 5, 10
	maze := generateMaze(n, m)
	printMaze(maze)
}
