package output

import (
	"github.com/alem-platform/ap"
)

func PrintChar(r rune) {
	_, err := ap.PutRune(r)
	if err != nil {
		panic("Invalid rune detected during output")
	}
}

func PrintString(s string) {
	for _, r := range s {
		PrintChar(r)
	}
}

func PrintLine(s string) {
	PrintString(s)
	PrintChar('\n')
}

func PrintInt(n int) {
	PrintString(IntToString(n))
}

func IntToString(n int) string {
	if n == 0 {
		return "0"
	}

	var isNegative bool
	if n < 0 {
		isNegative = true
		n = -n
	}

	digits := []rune{}
	for n > 0 {
		digit := n % 10
		digits = append([]rune{'0' + rune(digit)}, digits...)
		n /= 10
	}

	if isNegative {
		digits = append([]rune{'-'}, digits...)
	}

	return string(digits)
}

func PadString(s string, length int) string {
	for len(s) < length {
		s += " "
	}
	return s
}
