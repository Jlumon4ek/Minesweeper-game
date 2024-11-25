package game

import (
	gridPkg "crunch02/grid"
)

func RevealCell(grid []string, revealed [][]bool, x, y int) {
	height := len(grid)
	width := len(grid[0])

	if x < 0 || x >= height || y < 0 || y >= width || revealed[x][y] {
		return
	}

	revealed[x][y] = true

	if grid[x][y] == '.' {
		adjacentMines := gridPkg.CountAdjacentMines(grid, x, y)
		if adjacentMines == 0 {
			directions := []struct{ dx, dy int }{
				{-1, -1}, {-1, 0}, {-1, 1},
				{0, -1}, {0, 1},
				{1, -1}, {1, 0}, {1, 1},
			}

			for _, dir := range directions {
				RevealCell(grid, revealed, x+dir.dx, y+dir.dy)
			}
		}
	}
}

// -1, -1  |  -1,  0  |  -1,  1
// --------|-----------|---------
//  0, -1  |    X      |   0,  1
// --------|-----------|---------
//  1, -1  |   1,  0   |   1,  1
// Для каждого направления из directions вычисляются новые координаты и запускается рекурсивный вызов RevealCell для соседней клетки.
