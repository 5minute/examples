package main

import "fmt"

var field = [9][9]int{
	{5, 3, 0, 0, 7, 0, 0, 0, 0},
	{6, 0, 0, 1, 9, 5, 0, 0, 0},
	{0, 9, 8, 0, 0, 0, 0, 6, 0},
	{8, 0, 0, 0, 6, 0, 0, 0, 3},
	{4, 0, 0, 8, 0, 3, 0, 0, 1},
	{7, 0, 0, 0, 2, 0, 0, 0, 6},
	{0, 6, 0, 0, 0, 0, 2, 8, 0},
	{0, 0, 0, 4, 1, 9, 0, 0, 5},
	{0, 0, 0, 0, 8, 0, 0, 7, 9},
}

func draw() {
	for _, row := range field {
		fmt.Println(row)
	}
	fmt.Println()
}

func canPut(x int, y int, value int) bool {
	return !alreadyInVertical(x, y, value) &&
		!alreadyInHorizontal(x, y, value) &&
		!alreadyInSquare(x, y, value)
}

func alreadyInVertical(x int, y int, value int) bool {
	for i := range [9]int{} {
		if field[i][x] == value {
			return true
		}
	}
	return false
}

func alreadyInHorizontal(x int, y int, value int) bool {
	for i := range [9]int{} {
		if field[y][i] == value {
			return true
		}
	}
	return false
}

func alreadyInSquare(x int, y int, value int) bool {
	sx, sy := int(x/3)*3, int(y/3)*3
	for dy := range [3]int{} {
		for dx := range [3]int{} {
			if field[sy+dy][sx+dx] == value {
				return true
			}
		}
	}
	return false
}

func next(x int, y int) (int, int) {
	nextX, nextY := (x+1)%9, y
	if nextX == 0 {
		nextY = y + 1
	}
	return nextX, nextY
}

func solve(x int, y int) bool {
	if y == 9 {
		return true
	}
	if field[y][x] != 0 {
		return solve(next(x, y))
	} else {
		for i := range [9]int{} {
			var v = i + 1
			if canPut(x, y, v) {
				field[y][x] = v
				if solve(next(x, y)) {
					return true
				}
				field[y][x] = 0
			}
		}
		return false
	}
}

func main() {
	draw()
	if solve(0, 0) {
		fmt.Println("Found solution")
		draw()
	} else {
		fmt.Println("No solution")
	}
}
