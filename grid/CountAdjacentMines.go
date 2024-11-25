package grid

func CountAdjacentMines(grid []string, x, y int) int {
	height := len(grid)
	width := len(grid[0])
	count := 0

	directions := []struct{ dx, dy int }{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1}, {0, 1},
		{1, -1}, {1, 0}, {1, 1},
	}

	for _, dir := range directions {
		// Рассчитываем координаты соседней клетки
		nx, ny := x+dir.dx, y+dir.dy
		if nx >= 0 && nx < height && ny >= 0 && ny < width && grid[nx][ny] == '*' {
			count++
		}
	}

	return count
}
