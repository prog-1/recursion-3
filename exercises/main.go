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

type Index struct {
	i, j int
}

// Returns randomly generated maze of given size(n - height, m - width)
// Note that maze is represented as 2D array of bytes
func generateWays(n, m uint) [][]byte {
	// Creating 2D byte array, which represents a maze:
	maze := make([][]byte, n)
	for i := range maze {
		maze[i] = make([]byte, m)
	}

	// Creating array to store whether some cell was visited or not:
	visited := make([][]bool, n)
	for i := range visited {
		visited[i] = make([]bool, m)
	}

	gen(maze, visited, Index{0, 0}, N)

	return maze
}

// Auxiliary function for generateMaze()
func gen(maze [][]byte, visited [][]bool, cur Index, entry byte) {
	cell := &maze[cur.i][cur.j]

	visited[cur.i][cur.j] = true // Marking current cell as visited

	ways := getAvailableWays(*cell)

	// Removing entry way:
	var k int
	for k = 0; ways[k] != entry; k++ { // Getting index of entry way in slice of available ways
	}
	*cell |= ways[k]          // Marking way as visited
	ways = removeWay(ways, k) // Removing entry way out of slice of available ways

	for len(ways) != 0 {
		// Getting random available direction:
		r := rand.Intn(len(ways)) // Getting index of random available direction
		way := ways[r]
		ways = removeWay(ways, r)

		// Try to proceed with next cell at direction:
		n := getIndexByWay(way, cur)
		// Checking if we haven't went out of maze's bounds or if we've visited it already:
		if n.i >= 0 && n.i < len(maze) && n.j >= 0 && n.j < len(maze[0]) && !visited[n.i][n.j] {
			*cell |= way // Marking way as visited
			gen(maze, visited, n, getEntry(way))
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
func getIndexByWay(way byte, i Index) Index {
	switch way {
	case N:
		return Index{i.i - 1, i.j}
	case E:
		return Index{i.i, i.j + 1}
	case S:
		return Index{i.i + 1, i.j}
	case W:
		return Index{i.i, i.j - 1}
	default: // Should never happen
		return Index{0, 0}
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
	// Printing output array:
	for i := range maze {
		for j := range maze[i] {
			fmt.Print(string(maze[i][j]))
		}
		fmt.Println()
	}
}

// Returns 2D slice of bytes(chars) as graphical represenation of maze
func getMaze(ways [][]byte) [][]byte {
	// Creating char array, which will be displayed on screen:
	output := make([][]byte, len(ways)*3)
	for i := range output {
		output[i] = make([]byte, len(ways[0])*3)
	}

	// Filling output array:
	for i, k := 0, 1; i < len(ways); i, k = i+1, k+3 {
		for j, w := 0, 1; j < len(ways[i]); j, w = j+1, w+3 {
			output[k][w] = ' '     // Middle
			output[k-1][w-1] = '+' // NorthWest
			output[k-1][w+1] = '+' // NorthEast
			output[k+1][w+1] = '+' // SouthEash
			output[k+1][w-1] = '+' // SouthWest

			cell := ways[i][j]
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

	return output
}

// Writes solution path into given graphical represenation of maze
func solveMaze(maze [][]byte, exit Index) {
	// Creating array to store whether some char was visited or not:
	visited := make([][]bool, len(maze))
	for i := range visited {
		visited[i] = make([]bool, len(maze[i]))
	}

	maze[exit.i][exit.j] = 'E' // Marking exit

	start := Index{1, 1}
	var solved bool
	solve(start, maze, visited, &solved)
}

// Auxiliary function for solveMaze()
func solve(i Index, maze [][]byte, visited [][]bool, solved *bool) {
	// Processing current char:
	char := &maze[i.i][i.j]
	if *char == 'E' { // Exit
		*solved = true
		return
	}
	if (*char != ' ') { // Wall
		return
	} else { // Empty
		visited[i.i][i.j] = true
		*char = '*' // Marking as way
	}

	// Proceeding with the neighbors:
	ways := []byte{N, E, S, W}
	for len(ways) != 0 && !*solved {// NOTE: If maze is solved, then we have no need to go anywhere anymore
		way := ways[0]
		ways = removeWay(ways, 0)
		n := getIndexByWay(way, i)
		if n.i > 0 && n.i < len(maze)-1 && n.j > 0 && n.j < len(maze[0])-1 && !visited[n.i][n.j] {
			solve(n, maze, visited, solved)
		} // NOTE: We can skip border character processing cause we have nothing to do there
	}

	// On backtracking:
	if !*solved {
		*char = '@' // Marking unsuccessful way
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	var n, m uint
	fmt.Scanln(&n, &m)
	ways := generateWays(n, m)
	maze := getMaze(ways)
	exit := Index{len(maze) - 2, len(maze[0]) - 2}
	solveMaze(maze, exit)
	printMaze(maze)
}
