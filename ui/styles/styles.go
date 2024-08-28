package styles

import "github.com/charmbracelet/lipgloss"

type Styles struct {
	Colors struct{}

	Textarea struct {
		Style lipgloss.Style
	}
}

func New() *Styles {
	s := Styles{}

	s.Textarea.Style = lipgloss.NewStyle().
		Padding(1, 2).
		Border(lipgloss.ThickBorder()).
		Margin(2, 2)

	return &s
}
