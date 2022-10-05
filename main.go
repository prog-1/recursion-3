package main

import "fmt"

type (
	Coord struct{ X, Y int }
)

func check(c []int, exit Coord) bool {
	for i := range c {
		if i%2 == 0 && exit.X == c[i] && exit.Y == c[i+1] {
			return true
		}
	}
	return false
}

func fill(m [][]string, x, y int, c []int, start, exit Coord) []int {
	if m[y][x] == "E" || check(c, exit) {
		return c
	}
	if m[y][x] == "x" || m[y][x] == "." {
		return c
	}
	c = append(c, x, y)
	m[y][x] = "."
	fmt.Println(c)
	fill(m, x, y+1, c, start, exit)
	fill(m, x-1, y, c, start, exit)
	fill(m, x, y-1, c, start, exit)
	fill(m, x+1, y, c, start, exit)
	return c
}

func main() {
	m := [][]string{
		{"x", "x", "x", "x", "x", "x", "x", "x", "x", "x", "x", "x", "x", "x"},
		{"x", "x", " ", "x", "x", "x", "x", "x", " ", "x", "x", " ", "E", "x"},
		{"x", " ", " ", "x", " ", " ", " ", " ", " ", " ", " ", " ", "x", "x"},
		{"x", " ", "x", "x", " ", "x", " ", "x", "x", " ", "x", " ", " ", "x"},
		{"x", " ", "x", " ", " ", "x", " ", " ", "x", " ", "x", "x", " ", "x"},
		{"x", " ", "x", " ", "x", "x", "x", " ", "x", " ", " ", "x", "x", "x"},
		{"x", " ", "x", " ", " ", " ", "x", " ", "x", "x", " ", " ", " ", "x"},
		{"x", " ", "x", "x", "x", " ", "x", " ", " ", "x", "x", "x", " ", "x"},
		{"x", " ", " ", " ", " ", " ", "x", "x", " ", " ", " ", "x", " ", "x"},
		{"x", " ", "x", " ", "x", "x", "x", "x", " ", "x", " ", "x", " ", "x"},
		{"x", " ", "x", " ", "x", " ", " ", "x", "x", "x", " ", "x", " ", "x"},
		{"x", "x", "x", " ", "x", "x", " ", "x", " ", " ", " ", "x", " ", "x"},
		{"x", "S", " ", " ", " ", " ", " ", "x", "x", "x", "x", "x", " ", "x"},
		{"x", "x", "x", "x", "x", "x", "x", "x", "x", "x", "x", "x", "x", "x"},
	}
	var start, exit Coord
	var s []int
	start.X, start.Y, exit.X, exit.Y = 1, 12, 12, 1
	fill(m, start.X, start.Y, s, start, exit)
	m[start.Y][start.X] = "S"
	for row := range m {
		for col := range m {
			fmt.Printf(m[row][col])
		}
		fmt.Println()
	}
}
