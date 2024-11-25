package statistics

import "crunch02/output"

func PrintStatistics(height, width, mineCount, moveCount int, message string) {
	output.PrintLine(message)
	output.PrintLine("\nYour statistics:")
	output.PrintString(" - Field size: ")
	output.PrintInt(height)
	output.PrintString("x")
	output.PrintInt(width)
	output.PrintString("\n - Number of bombs: ")
	output.PrintInt(mineCount)
	output.PrintString("\n - Number of moves: ")
	output.PrintInt(moveCount)
	output.PrintChar('\n')
}
