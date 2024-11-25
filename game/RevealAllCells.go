package game

func RevealAllCells(grid []string, revealed [][]bool) {
	for i := range grid {
		for j := range grid[i] {
			revealed[i][j] = true
		}
	}
}
