package grid

import (
	"crunch02/output"
)

func PrintGrid(grid []string, revealed [][]bool) {
	output.PrintLine("\nGrid:")

	width := len(grid[0])

	// Нумерация столбцов
	output.PrintString("      ")
	for col := 1; col <= width; col++ {
		output.PrintString(output.PadString(output.IntToString(col), 8))
	}
	output.PrintChar('\n')

	output.PrintString("    ")
	for col := 1; col <= width; col++ {
		if col == width {
			output.PrintString("_______ ")
		} else {
			output.PrintString("________")
		}
	}
	output.PrintString(" ")
	output.PrintChar('\n')

	for rowIdx, row := range grid {
		for line := 0; line < 3; line++ {
			if line == 1 {
				output.PrintString(output.PadString(output.IntToString(rowIdx+1), 3) + "|")
			} else {
				output.PrintString("   |")
			}

			for colIdx := range row {
				if revealed[rowIdx][colIdx] {
					if line == 1 {
						if grid[rowIdx][colIdx] == '*' {
							output.PrintString("   *   |")
						} else {
							adjacentMines := CountAdjacentMines(grid, rowIdx, colIdx)
							if adjacentMines == 0 {
								output.PrintString("       |")
							} else {
								output.PrintString("   " + output.IntToString(adjacentMines) + "   |") // Число
							}
						}
					} else if line == 2 {
						output.PrintString("_______|")
					} else {
						output.PrintString("       |")
					}
				} else {
					if line == 1 {
						output.PrintString("XXXXXXX|")
					} else {
						output.PrintString("XXXXXXX|")
					}
				}
			}
			output.PrintChar('\n')
		}
	}
}
