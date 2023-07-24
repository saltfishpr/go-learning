package main

import "fmt"

type Position struct {
	// Offset is the byte offset, starting at 0.
	Offset int
	// Line is the line number, starting at 1.
	Line int
	// Column is the column number, starting at 1 (character count per line).
	Column int
}

// String returns a string representation of the position on the format <line>:<column>.
func (p Position) String() string {
	return fmt.Sprintf("%d:%d", p.Line, p.Column)
}
