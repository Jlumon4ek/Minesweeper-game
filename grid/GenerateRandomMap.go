package grid

import (
	"math/rand"
)

func GenerateRandomMap(height, width, mineCount int) []string {
	grid := make([][]rune, height)
	for i := range grid {
		grid[i] = make([]rune, width)
		for j := range grid[i] {
			grid[i][j] = '.'
		}
	}

	for m := 0; m < mineCount; {
		x := rand.Intn(height)
		y := rand.Intn(width)
		if grid[x][y] == '.' {
			grid[x][y] = '*'
			m++
		}
	}

	stringGrid := make([]string, height)
	for i := range grid {
		stringGrid[i] = string(grid[i])
	}
	return stringGrid
}
