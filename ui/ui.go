package ui

import (
	"fmt"

	"github.com/gookit/color"
)

// PrintRed prints a white title formatted with a red message.
func PrintRed(title string, message string) {
	red := color.FgRed.Render("%s\n")
	fmt.Printf("%-15s: "+red, title, message)
}

// PrintCyan prints a white title formatted with a cyan message.
func PrintCyan(title string, message string) {
	cyan := color.FgCyan.Render("%s\n")
	fmt.Printf("%-15s: "+cyan, title, message)
}
