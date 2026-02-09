package ui

import (
	"fmt"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/nathfavour/vibetype/internal/engine"
	"github.com/nathfavour/vibetype/internal/data"
	"github.com/nathfavour/vibetype/internal/styles"
)

type tickMsg time.Time

func tick() tea.Cmd {
	return tea.Tick(time.Millisecond*100, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}

type keyMap struct {
	Quit key.Binding
}

func (k keyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Quit}
}

func (k keyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{{k.Quit}}
}

var keys = keyMap{
	Quit: key.NewBinding(
		key.WithKeys("ctrl+c", "esc"),
		key.WithHelp("esc", "quit"),
	),
}

type Model struct {
	engine *engine.Engine
	help   help.Model
	width  int
	height int
}

func NewModel(text string) Model {
	return Model{
		engine: engine.NewEngine(text),
		help:   help.New(),
	}
}

func (m Model) Init() tea.Cmd {
	return tick()
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tickMsg:
		if !m.engine.Finished {
			return m, tick()
		}
		return m, nil

	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		m.help.Width = msg.Width

	case tea.KeyMsg:
		switch {
		case key.Matches(msg, keys.Quit):
			return m, tea.Quit
		case msg.Type == tea.KeyBackspace:
			m.engine.Pop()
		case msg.Type == tea.KeyEnter:
			if m.engine.Finished {
				m.engine = engine.NewEngine(data.GetRandomQuote())
				return m, tick()
			}
		case msg.Type == tea.KeyRunes:
			for _, r := range msg.Runes {
				m.engine.Push(r)
			}
			if m.engine.Active {
				return m, tick()
			}
		case msg.Type == tea.KeySpace:
			m.engine.Push(' ')
			if m.engine.Active {
				return m, tick()
			}
		}
	}
	return m, nil
}

func (m Model) View() string {

	if m.width == 0 || m.height == 0 {

		return "Initializing..."

	}



	var s strings.Builder



	if m.engine.Finished {

		s.WriteString(styles.TitleStyle.Render("Vibetype Results"))

		s.WriteString("\n\n")

		s.WriteString(fmt.Sprintf("WPM:      %.0f\n", m.engine.WPM()))

				s.WriteString(fmt.Sprintf("Accuracy: %.0f%%\n", m.engine.Accuracy()))

				s.WriteString("\n\n")

				s.WriteString(styles.DimStyle.Render("Press enter to restart • esc to quit"))

			} else {

		// Stats

		accuracy := m.engine.Accuracy()

		vibeColor := styles.GetVibeColor(accuracy)



		stats := fmt.Sprintf("WPM: %.0f | Accuracy: %.0f%%", m.engine.WPM(), accuracy)

		s.WriteString(lipgloss.NewStyle().Foreground(vibeColor).Bold(true).Render(stats))

		s.WriteString("\n\n")



		// Text

		var text strings.Builder

		for i, targetRune := range m.engine.TargetText {

			if i < len(m.engine.InputText) {

				inputRune := m.engine.InputText[i]

				if inputRune == targetRune {

					text.WriteString(styles.CorrectStyle.Render(string(targetRune)))

				} else {

					text.WriteString(styles.WrongStyle.Render(string(targetRune)))

				}

			} else if i == len(m.engine.InputText) {

				text.WriteString(styles.ActiveStyle.Render(string(targetRune)))

			} else {

				text.WriteString(styles.DimStyle.Render(string(targetRune)))

			}

		}



		// Wrap text based on width

		wrappedText := lipgloss.NewStyle().Width(m.width / 2).Render(text.String())

		s.WriteString(wrappedText)

		s.WriteString("\n\n")

		s.WriteString(m.help.View(keys))

	}



	return lipgloss.Place(m.width, m.height, lipgloss.Center, lipgloss.Center, s.String())

}
