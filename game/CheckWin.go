package game

func CheckWin(revealed [][]bool, grid []string) bool {
	for i := range grid {
		for j := range grid[i] {
			// Проверка безопасная ли клетка и открыта ли она
			if grid[i][j] == '.' && !revealed[i][j] {
				return false
			}
		}
	}
	return true
}

// revealed [][]bool - массив булевых значений, показывающий, какие клетки уже открыты
