package inputhandler

import (
	"crunch02/grid"
	"crunch02/output"
	"crunch02/validation"
	"fmt"
)

func HandleInput(mode int) (int, int, []string) {
	var height, width int
	if mode == 1 {
		output.PrintString("Enter height and width: ")
		fmt.Scanf("%d %d\n", &height, &width)
		if !validation.ValidateGridSize(height, width) {
			return 0, 0, nil
		}

		grid := make([]string, height)
		output.PrintLine("Enter the grid (use '.' for empty and '*' for mines):")
		for i := 0; i < height; i++ {
			var row string
			fmt.Scanf("%s\n", &row)
			grid[i] = row
		}

		if !validation.ValidateGridContent(grid, width) {
			return 0, 0, nil
		}
		return height, width, grid
	} else if mode == 2 {
		output.PrintString("Enter height and width: ")
		fmt.Scanf("%d %d\n", &height, &width)
		if !validation.ValidateGridSize(height, width) {
			return 0, 0, nil
		}

		mineCount := (height * width) / 4
		return height, width, grid.GenerateRandomMap(height, width, mineCount)
	} else {
		output.PrintLine("Invalid choice!")
		return 0, 0, nil
	}
}
