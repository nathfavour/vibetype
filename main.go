package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/nathfavour/vibetype/internal/data"
	"github.com/nathfavour/vibetype/internal/ui"
)

func main() {
	m := ui.NewModel(data.GetRandomQuote())
	p := tea.NewProgram(m, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(1)
	}
}