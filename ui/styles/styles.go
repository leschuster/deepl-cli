package styles

import (
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/lipgloss"
)

type Styles struct {
	Colors struct{}

	Textarea struct {
		Style lipgloss.Style
	}

	LanguageSelect struct {
		PromptStyle lipgloss.Style
		CursorStyle lipgloss.Style
		Width       int
	}

	List struct {
		Styles             list.Styles
		NormalTitleStyle   lipgloss.Style
		SelectedTitleStyle lipgloss.Style
	}

	LangView struct {
		Style lipgloss.Style
	}
}

func New() *Styles {
	s := Styles{}

	s.Textarea.Style = lipgloss.NewStyle().
		Padding(1, 2).
		Border(lipgloss.ThickBorder()).
		Margin(2, 2)

	s.LangView.Style = lipgloss.NewStyle().
		Padding(1, 2).
		Border(lipgloss.RoundedBorder())

	return &s
}
