package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
PLAN: ⬇⬆
Generation:

#1 Initialize cell struct ✔
	Cell struct contains 5 booleans. First 4 bool represend directions (N,E,S,W). True - emptyness, False - wall.
	Last bool respond on is visited or not.

#2 Initialize 2d cell struct slice. ✔
	All cells are covered with walls by default.

#3 Make getRandDir function ✔
	Select random direction from given direction slice

#4 Generate Maze function ✔
	Call walk function and start recursion ✔

#5 Walk function
	- Check if there are neighbours that you can visit in the cell. ✔

#6 Initialize "Stack" slize (for backtracking)
	To save the cells we went through and delete when backtrack.
	NOPE! We don't need the stack, because we will check each direction of each cell anyways.

#6.5 Pick random cell direction ✔
	Using getRandDir ✔
	We need to take next cell from direction ✔

#7 Get next cell from chosen direction ✔
	- check if next cell is not outside of maze borders ✔
	- check if next cell is not visited yet ✔
	- on failure pick next direction (return to 3 ⬆) ✔
	- on success continue ⬇ ✔

#8 Make path
	- put the next cell in a Stack slice (for backtracking) ✔
	- delete walls between current cell and next cell. ✔
	- delete selected direction from current cell (for backtracking) ✔

#9 Call recursion ✔
	Repeat steps from 3 ⬆ ✔

#10 Backtracking
	If we arrived in the dead end (checks from #4 haven't passed on all dirs) and we need to go back:
	- Mark current (dead end) cell as visited.
	- Delete current cell from Stack. (pop the slice)
	- Return to 3 ⬆
	Backtrack to start cell
	By the end we will backtrack back to starting cell, filling the gaps we left earlier.
	By the end the Stack will be empty and all cells will be visited.

#12 Destroy Wall function
	Set selected direction to true on cur cell ✔
	Set opposite direction to next cell ✔
*/

// ## CELL STRUCT ##
type cellStruct struct {
	n, e, s, w bool
	visited    bool
}

// ## CELL COORDINATES ##
type cellCoord struct {
	y, x int
}

func main() {

	rand.Seed(time.Now().UnixNano()) //Making seed for random always be different

	// ## INITIALIZING 2D CELL SLICE ##
	var cols, rows int
	fmt.Scanln(&cols, &rows) // reading maze length
	//cols, rows := 10, 50                //for debug
	cells := make([][]cellStruct, cols) // creating cell slice with given columns
	for i := range cells {
		cells[i] = make([]cellStruct, rows) // editing cell slice to have given rows
	}

	//## MAZE ##
	generateMaze(cells)
	maze := createMaze(cells)
	//fmt.Print(cells)//for degub
	solveMaze(maze, cellCoord{1, 1}, cellCoord{len(maze) - 2, len(maze[0]) - 2})
	printMaze(maze)
}

// ## GENERATE MAZE ##
func generateMaze(cells [][]cellStruct) {
	//initCell := cells[0][0] // declaring start/initial cell

	walk(cells, cellCoord{0, 0}) // start walking (recursive function)
}

// ## WALK ##
func walk(cells [][]cellStruct, curCoord cellCoord) {
	//Get cell ✔
	//Get cell directions ✔
	//Get random cell direction ✔
	//Get next cell ✔
	//Make checks (borders and visited) ✔
	//Delete walls... ✔
	//Struggles:
	// How to store cell directions? ✔ - Get directions function!
	// I don't know how to walk. ✔ - you need next cell!

	cell := &cells[curCoord.y][curCoord.x] // take cell from given coordinates
	cell.visited = true

	dirs := getDirs(cell) // take available cell directions

	for len(dirs) != 0 { // till we haven't checked, thus deleted, all the directions of cell
		//taking directions from current cell. Taking random direction. Taking next cell coordinates from selected direction.
		var dir byte                            // random direction variable
		dirs, dir = getRandDir(dirs)            // saving random direction and updating cell direction slice
		nextCoord := getNextCell(dir, curCoord) // taking next cell coordinates from selected direction

		if 0 > nextCoord.y || nextCoord.y >= len(cells) || 0 > nextCoord.x || nextCoord.x >= len(cells[0]) { // if we went outside of maze borders
			continue //starting new iteration of the loop
		}

		nextCell := &cells[nextCoord.y][nextCoord.x] // taking next cell from coordinates
		//(I know that it is extra variable declaration, but it's more readable for me this way)
		if nextCell.visited { // if the next cell is visited
			continue
		}

		//if we are here it means the next cell is the available empty cell and we are going there

		destroyWallz(cell, nextCell, dir) //break walls. Set selected direction to true (path)
		walk(cells, nextCoord)            // calling recursion, switching to next cell

	}

}

// ## GET RANDOM DIRECTION ##
func getRandDir(dirs []byte) ([]byte, byte) {
	// we need dir slice as input to be able to not count used dirs.
	// Return n,e,s or w.
	i := rand.Intn(len(dirs))              // take random dir index
	d := dirs[i]                           // to return this exact dir
	dirs = append(dirs[:i], dirs[i+1:]...) // delete selected dir from dir slice
	return dirs, d
}

// ## GET NEXT CELL FROM DIRECTION ##
func getNextCell(dir byte, curCoord cellCoord) (nextCoord cellCoord) {
	switch dir {
	case 'n': //north
		nextCoord = cellCoord{curCoord.y + 1, curCoord.x}
	case 'e': //east
		nextCoord = cellCoord{curCoord.y, curCoord.x + 1}
	case 's': //south
		nextCoord = cellCoord{curCoord.y - 1, curCoord.x}
	case 'w': //west
		nextCoord = cellCoord{curCoord.y, curCoord.x - 1}
	}
	return nextCoord
}

// ## GET AVAILABLE CELL DIRECTIONS ##
func getDirs(cell *cellStruct) (dirs []byte) {
	if !cell.n { // true - empty (available), false - wall (unavailable)
		dirs = append(dirs, 'n')
	}
	if !cell.e {
		dirs = append(dirs, 'e')
	}
	if !cell.s {
		dirs = append(dirs, 's')
	}
	if !cell.w {
		dirs = append(dirs, 'w')
	}
	return dirs
}

// ## DESTROY WALLS ON GIVEN DIRECTION ##
func destroyWallz(cell *cellStruct, nextCell *cellStruct, dir byte) {
	switch dir {
	case 'n': //north
		cell.n = true     // setting wall on dir for current cell
		nextCell.s = true // setting wall on opposite dir for next cell
	case 'e': //east
		cell.e = true
		nextCell.w = true
	case 's': //south
		cell.s = true
		nextCell.n = true
	case 'w': //west
		cell.w = true
		nextCell.e = true
	}
}

// ## PRINTING MAZE IN CONSOLE ##
func printMaze(maze [][]rune) {
	for i := range maze {
		for j := range maze[i] {
			fmt.Print(string(maze[i][j]))
			//if you will not parse maze to strings, you will return numbers of chars.
		}
		fmt.Println()
	}
}

// ## CREATING MAZE FROM CELL STRUCT SLICE ##
func createMaze(cells [][]cellStruct) [][]rune {
	// Creating maze 2d byte slice
	maze := make([][]rune, len(cells)*3)
	//maze slice is 3x bigger then cell slice, because
	//one cell is 3x3 element in maze
	for i := range maze {
		maze[i] = make([]rune, len(cells[0])*3)
	}

	// Filling output array:
	for i, y := 0, 1; i < len(cells); i, y = i+1, y+3 {
		for j, x := 0, 1; j < len(cells[i]); j, x = j+1, x+3 {
			maze[y][x] = ' '     // Center
			maze[y+1][x+1] = '┌' // South-Eash
			maze[y-1][x+1] = '└' // North-East
			maze[y+1][x-1] = '┐' // South-West
			maze[y-1][x-1] = '┘' // North-West

			cell := cells[i][j]
			if cell.n { //North
				maze[y+1][x] = ' '
			} else {
				maze[y+1][x] = '─'
			}
			if cell.e { //East
				maze[y][x+1] = ' '
			} else {
				maze[y][x+1] = '│'
			}
			if cell.s { //South
				maze[y-1][x] = ' '
			} else {
				maze[y-1][x] = '─'
			}
			if cell.w { //West
				maze[y][x-1] = ' '
			} else {
				maze[y][x-1] = '│'
			}

		}
	}

	return maze
}

func solveMaze(maze [][]rune, startCoord, endCoord cellCoord) {

	//Initializing visited 2d slice
	visited := make([][]bool, len(maze)) // creating cell slice with given columns
	for i := range maze {
		visited[i] = make([]bool, len(maze[i])) // editing cell slice to have given rows
	}

	var solved bool
	passMaze(maze, startCoord, endCoord, visited, &solved)

}

// function to reccursively walk in the maze
func passMaze(maze [][]rune, curCoord, endCoord cellCoord, visited [][]bool, solved *bool) {

	//check if it is exit ✔
	//mark cur cell as visited, set cell symbol to '•' ✔
	//check available directions (for loop) ✔
	//move along all available directions clockwise from n ✔
	//delete checked direction ✔
	//check maze borders and visited ✔
	//check if it is wall or emptyness ✔
	//if all checks passed and it's emptyness, we are going that way (calling recursion) ✔
	//after loop, if we are backtracking when maze is not solved yet, set cell symbol to 'o' ✔

	if curCoord == endCoord {
		maze[curCoord.y][curCoord.x] = '✔'
		*solved = true
		return
	}

	cell := &maze[curCoord.y][curCoord.x]
	*cell = '•'
	visited[curCoord.y][curCoord.x] = true

	dirs := []byte{'n', 'e', 's', 'w'} //initializing all 4 directions for each cell (checks later)
	for len(dirs) != 0 && !*solved {
		nextCoord := getNextCell(dirs[0], curCoord)                                                        //getting next cell from direction
		dirs = dirs[1:]                                                                                    //deleting first direction from dirs slice
		if 0 > nextCoord.y || nextCoord.y >= len(maze) || 0 > nextCoord.x || nextCoord.x >= len(maze[0]) { // if we went outside of maze borders
			continue //starting new iteration of the loop
		}
		if visited[nextCoord.y][nextCoord.x] { // if the next cell is visited
			continue
		}
		if maze[nextCoord.y][nextCoord.x] != ' ' { // if next cell is no matter which type of wall
			continue
		}
		passMaze(maze, nextCoord, endCoord, visited, solved)
	}
	//Backtracking
	if !*solved {
		//if we are backtracking when maze is not solved, we appeared in dead end.
		//If we are backtracking when maze is solved, we are not in dead end.
		*cell = 'o'
	}

}
