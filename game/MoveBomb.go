package game

import (
	"math/rand"
)

func MoveBomb(grid []string, x, y int) {
	height := len(grid)
	width := len(grid[0])

	emptyCells := []struct{ row, col int }{}
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			if grid[i][j] == '.' {
				// Создаётся срез emptyCells для хранения координат всех пустых клеток.
				emptyCells = append(emptyCells, struct{ row, col int }{i, j})
			}
		}
	}

	if len(emptyCells) == 0 {
		return
	}

	randIndex := rand.Intn(len(emptyCells))
	target := emptyCells[randIndex]

	row := []rune(grid[target.row]) // Преобразуем строку в массив рун
	row[target.col] = '*'
	grid[target.row] = string(row) // Преобразуем обратно в строку и записываем в grid

	row = []rune(grid[x]) // Преобразуем строку в массив рун
	row[y] = '.'          // Меняем символ в исходной позиции на '.'
	grid[x] = string(row) // Преобразуем обратно в строку и записываем в grid
}
