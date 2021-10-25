package main

import (
	"fmt"
	"math"
)

var field [4][4]int
var used [18]bool

func ok(v1 int, v2 int) bool {
	return math.Abs(float64(v1-v2)) >= 2 || v1%v2 == 0 || v2%v1 == 0
}

func solve(p int) bool {
	if p == 16 {
		for _, row := range field {
			fmt.Println(row)
		}
		return true
	}
	x, y := p%4, int(p/4)
	for i := range [16]int{} {
		if used[i] {
			continue
		}
		v := i + 3
		if (x > 0 && !ok(v, field[y][x-1])) || (y > 0 && !ok(v, field[y-1][x])) {
			return false
		}
		used[i] = true
		field[y][x] = v
		if solve(p + 1) {
			return true
		}
		used[i] = false
	}
	return false

}

func main() {
	solve(0)
}
