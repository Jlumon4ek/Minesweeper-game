package main

import (
	"crunch02/game"
	gridPkg "crunch02/grid"
	"crunch02/inputhandler"
	"crunch02/output"
	"crunch02/statistics"
	"crunch02/validation"
	"fmt"
)

func main() {
	output.PrintLine("Choose a mode:")
	output.PrintLine("1. Enter a custom map")
	output.PrintLine("2. Generate a random map")
	output.PrintString("Enter your choice: ")
	var mode int
	fmt.Scanf("%d\n", &mode)

	height, width, grid := inputhandler.HandleInput(mode)
	if height == 0 || width == 0 || grid == nil {
		return
	}

	if !validation.ValidateMineCount(grid, height, width) {
		return
	}

	revealed := make([][]bool, height)
	for i := range revealed {
		revealed[i] = make([]bool, width)
	}

	moveCount := 0
	firstMove := true

	for {
		gridPkg.PrintGrid(grid, revealed)

		output.PrintString("Enter your coordinates (x and y): ")
		var x, y int
		fmt.Scanf("%d %d\n", &x, &y)

		x, y = x-1, y-1
		if x < 0 || x >= height || y < 0 || y >= width {
			output.PrintLine("Invalid coordinates! Please try again.")
			continue
		}

		moveCount++
		if firstMove && grid[x][y] == '*' {
			output.PrintLine("You hit a bomb on your first move! Moving the bomb to a safe location...")
			game.MoveBomb(grid, x, y)
		}
		firstMove = false

		if grid[x][y] == '*' {
			game.RevealAllCells(grid, revealed)
			gridPkg.PrintGrid(grid, revealed)
			statistics.PrintStatistics(height, width, len(grid), moveCount, "Game Over!")
			break
		}

		game.RevealCell(grid, revealed, x, y)

		if game.CheckWin(revealed, grid) {
			game.RevealAllCells(grid, revealed)
			gridPkg.PrintGrid(grid, revealed)
			statistics.PrintStatistics(height, width, len(grid), moveCount, "You Win!")
			break
		}
	}
}
