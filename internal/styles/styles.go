package styles

import "github.com/charmbracelet/lipgloss"

var (
	MainColor    = lipgloss.Color("5")
	SuccessColor = lipgloss.Color("10")
	ErrorColor   = lipgloss.Color("9")
	DimColor     = lipgloss.Color("240")
	ActiveColor  = lipgloss.Color("255")

	TitleStyle = lipgloss.NewStyle().Foreground(MainColor).Bold(true)
	StatsStyle = lipgloss.NewStyle().Foreground(MainColor)
	DimStyle   = lipgloss.NewStyle().Foreground(DimColor)
	
	CorrectStyle = lipgloss.NewStyle().Foreground(SuccessColor)
	WrongStyle   = lipgloss.NewStyle().Foreground(ErrorColor).Underline(true)
	ActiveStyle  = lipgloss.NewStyle().Foreground(ActiveColor).Underline(true)
)

func GetVibeColor(accuracy float64) lipgloss.Color {
	if accuracy < 70 {
		return ErrorColor
	}
	if accuracy < 90 {
		return lipgloss.Color("3") // Yellow/Orange
	}
	return MainColor
}
