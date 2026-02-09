package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/nathfavour/vibetype/internal/ui"
)

const defaultText = "The quick brown fox jumps over the lazy dog. A journey of a thousand miles begins with a single step. To be or not to be, that is the question."

func main() {
	m := ui.NewModel(defaultText)
	p := tea.NewProgram(m, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(1)
	}
}