package validation

import (
	"crunch02/output"
)

func ValidateGridSize(height, width int) bool {
	if height < 3 || width < 3 {
		output.PrintLine("Error: invalid input. The grid size min 3x3")
		return false
	}
	return true
}

func ValidateGridContent(grid []string, width int) bool {
	for _, row := range grid {
		if len(row) != width {
			output.PrintLine("Row size must match the specified width!")
			return false
		}
		for _, char := range row {
			if char != '.' && char != '*' {
				output.PrintLine("Error: Invalid characters in grid. Only '.' and '*' are allowed.")
				return false
			}
		}
	}
	return true
}

func ValidateMineCount(grid []string, height, width int) bool {
	mineCount := 0
	for _, row := range grid {
		for _, cell := range row {
			if cell == '*' {
				mineCount++
			}
		}
	}

	if mineCount < 2 {
		output.PrintLine("The grid must contain at least 2 mines ('*').")
		return false
	}
	if mineCount >= height*width {
		output.PrintLine("Error: The number of bombs cannot be equal to or greater than the total number of cells.")
		return false
	}
	return true
}
