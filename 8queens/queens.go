package main

import "fmt"

const n = 8

var board = [n][n]int{}
var numSols = 0

func underAttack(x int, y int, dx int) bool {
	x, y = x+dx, y-1
	for y >= 0 && x >= 0 && x < n {
		if board[y][x] == 1 {
			return true
		}
		x = x + dx
		y = y - 1
	}
	return false
}

func solve(y int) {
	if y == n {
		numSols = numSols + 1
		for _, row := range board {
			fmt.Println(row)
		}
		fmt.Println()
	}
	for x := range [n]int{} {
		if !(underAttack(x, y, 0) || underAttack(x, y, -1) || underAttack(x, y, 1)) {
			board[y][x] = 1
			solve(y + 1)
			board[y][x] = 0
		}
	}
}

func main() {
	solve(0)
	fmt.Printf("Found %d solution(s)\n", numSols)
}
