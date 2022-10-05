package main

import (
	"fmt"
	"math/rand"
	"time"
)

var a = [][]byte{
	[]byte("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"),
	[]byte("X..E..X.....X........X........X"),
	[]byte("X..XXXX..XXXXXXX..XXXXXXX..XXXX"),
	[]byte("X..X.....X.....X..X..X..X..X..X"),
	[]byte("X..XXXX..X..X..X..X..X..X..X..X"),
	[]byte("X.....X.....X.....X........X..X"),
	[]byte("X..XXXXXXXXXX..XXXX..XXXX..X..X"),
	[]byte("X..X..X..............X.....X..X"),
	[]byte("X..X..XXXX..X..XXXX..X..XXXX..X"),
	[]byte("X...........X..X.....X..X.....X"),
	[]byte("XXXX..XXXXXXX..X..XXXXXXX..XXXX"),
	[]byte("X..X..X.....X..X..X.....X..X..X"),
	[]byte("X..X..X..XXXX..X..X..X..X..X..X"),
	[]byte("X..X.....X.....X.....X..X.....X"),
	[]byte("X..XXXXXXXXXXXXXXXX..X..XXXX..X"),
	[]byte("X.....X........X..X..X........X"),
	[]byte("X..XXXX..X..X..X..X..XXXXXXX..X"),
	[]byte("X........X..X..............X..X"),
	[]byte("XXXXXXXXXXXXX..XXXXXXXXXXXXX..X"),
	[]byte("X..............X.............SX"),
	[]byte("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"),
}

type coord struct {
	x uint
	y uint
}

func FindPath(m [][]byte, current coord) (successful bool, path []coord) {
	if m[current.y][current.x] == 'E' {
		return true, append(path, current)
	}
	m[current.y][current.x] = byte('*')
	path = append(path, current)
	if m[current.y+1][current.x] != 'X' && m[current.y+1][current.x] != '*' {
		suc, path := FindPath(m, coord{x: current.x, y: current.y + 1})
		if suc {
			m[current.y][current.x] = '#'
			return suc, append(path, current)
		}
	}
	if m[current.y-1][current.x] != 'X' && m[current.y-1][current.x] != '*' {
		suc, path := FindPath(m, coord{x: current.x, y: current.y - 1})
		if suc {
			m[current.y][current.x] = '#'

			return suc, append(path, current)
		}
	}
	if m[current.y][current.x+1] != 'X' && m[current.y][current.x+1] != '*' {
		suc, path := FindPath(m, coord{x: current.x + 1, y: current.y})
		if suc {
			m[current.y][current.x] = '#'
			return suc, append(path, current)
		}
	}
	if m[current.y][current.x-1] != 'X' && m[current.y][current.x-1] != '*' {
		suc, path := FindPath(m, coord{x: current.x - 1, y: current.y})
		if suc {
			m[current.y][current.x] = '#'
			return suc, append(path, current)
		}
	}
	return false, path
}

func contains(s []int, a int) bool {
	for _, v := range s {
		if a == v {
			return true
		}
	}
	return false
}

func randomChoice(length int) (res []int) {
	for i := random(1, length+1); i >= 0; i-- {
		tmp := random(1, length+1)
		if !contains(res, tmp) {
			res = append(res, tmp)
		}
	}
	return
}

func generateMaze(x, y int) [][]Sector {
	maze := make([][]Sector, y)
	var id int
	for i := range maze {
		maze[i] = make([]Sector, x)
		for i2 := range maze[i] {
			maze[i][i2].id = id
			id++
		}
	}
	for line := range maze {
		for cell := range maze[line] {
			if cell != len(maze[line])-1 {
				if line == len(maze)-1 {
					maze[line][cell].canWalkRight = true
				} else {
					maze[line][cell].canWalkRight = rand.Intn(2) == 0
				}
				if maze[line][cell].canWalkRight {
					maze[line][cell+1].id = maze[line][cell].id
				}
			}
		}

		if line != len(maze)-1 {
			for i := 0; i < id; i++ {
				var cont []int
				for cell := range maze[line] {
					if maze[line][cell].id == i {
						cont = append(cont, cell)
					}
				}
				if len(cont) > 1 {
					tmp := randomChoice(len(cont) - 1)
					for _, v := range tmp {
						maze[line+1][cont[v]].canWalkUp = true
						maze[line+1][cont[v]].id = maze[line][cont[v]].id
					}
				} else if len(cont) == 1 {
					maze[line+1][cont[0]].canWalkUp = true
					maze[line+1][cont[0]].id = maze[line][cont[0]].id
				}
			}
			// var changed bool
			/*			for cell, watchID, start := 0, maze[line][0].id, 0; cell <= len(maze[line]); cell++ {
								if cell == len(maze[line]) || watchID != maze[line][cell].id {
									if !changed {
										//mt.Println(start, cell, random(start, cell))
										tmp := random(start, cell)
										maze[line+1][tmp].canWalkUp = true
										maze[line+1][tmp].id = maze[line][tmp].id
									}
									if cell != len(maze[line]) {
										changed = false
										watchID = maze[line][cell].id
										start = cell
									}
								}
							}
						}
					}*/
		}
	}
	return maze
}

func random(min, max int) int {
	return rand.Intn(max-min) + min
}

type Sector struct {
	id           int
	canWalkRight bool
	canWalkUp    bool
}

func main() {
	// for i, v := range a {
	// 	for i2, v2 := range v {
	// 		if v2 == 'S' {
	// 			fmt.Println(i2, i)
	// 		}
	// 	}
	// }
	// fmt.Println(FindPath(a, coord{x: 29, y: 19}))
	// for _, v := range a {
	// 	for _, v2 := range v {
	// 		fmt.Print(string(v2))
	// 	}
	// 	fmt.Println()
	// }
	rand.Seed(time.Now().UnixNano())
	maze := generateMaze(10, 10)
	fmt.Println(maze)
	for line := range maze {
		fmt.Print("|")
		for _, cell := range maze[line] {
			if cell.canWalkUp {
				if cell.canWalkRight {
					fmt.Print("  ")
				} else {
					fmt.Print(" |")
				}
			} else {
				if cell.canWalkRight {
					fmt.Print("‾ ")
				} else {
					fmt.Print("‾|")
				}
			}
		}
		fmt.Println()
	}
}
