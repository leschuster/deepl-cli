package styles

import (
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/lipgloss"
)

type Styles struct {
	Colors struct {
		Primary lipgloss.AdaptiveColor
		Text    lipgloss.AdaptiveColor
		Title   struct {
			Foreground lipgloss.AdaptiveColor
			Background lipgloss.AdaptiveColor
		}
		ButtonForeground      lipgloss.AdaptiveColor
		ButtonBackground      lipgloss.AdaptiveColor
		ButtonHoverForeground lipgloss.AdaptiveColor
		ButtonHoverBackground lipgloss.AdaptiveColor
	}

	Textarea struct {
		Style lipgloss.Style
	}

	LanguageSelect struct {
		PromptStyle lipgloss.Style
		CursorStyle lipgloss.Style
		Width       int
	}

	List struct {
		Style              list.Styles
		NormalTitleStyle   lipgloss.Style
		SelectedTitleStyle lipgloss.Style
	}

	LangView struct {
		Style lipgloss.Style
	}

	LanguageButton struct {
		Style        lipgloss.Style
		HoveredStyle lipgloss.Style
	}
}

func New() *Styles {
	s := Styles{}

	s.Colors.Primary = lipgloss.AdaptiveColor{Light: "#EE6FF8", Dark: "#EE6FF8"}
	s.Colors.Text = lipgloss.AdaptiveColor{Light: "#1a1a1a", Dark: "#dddddd"}
	s.Colors.Title.Foreground = lipgloss.AdaptiveColor{Light: "62", Dark: "62"}
	s.Colors.Title.Background = lipgloss.AdaptiveColor{Light: "230", Dark: "230"}
	s.Colors.ButtonForeground = lipgloss.AdaptiveColor{Light: "3", Dark: "3"}
	s.Colors.ButtonBackground = lipgloss.AdaptiveColor{Light: "360", Dark: "360"}
	s.Colors.ButtonHoverForeground = lipgloss.AdaptiveColor{Light: "82", Dark: "82"}
	s.Colors.ButtonHoverBackground = lipgloss.AdaptiveColor{Light: "42", Dark: "42"}

	s.Textarea.Style = lipgloss.NewStyle().
		Padding(1, 2).
		Border(lipgloss.ThickBorder()).
		Margin(2, 2)

	s.List.Style = list.DefaultStyles()
	s.List.Style.TitleBar.Align(lipgloss.Center)
	s.List.Style.Title = lipgloss.NewStyle().
		Foreground(s.Colors.Title.Foreground).
		Background(s.Colors.Title.Background).
		Padding(0, 2)

	s.List.NormalTitleStyle = lipgloss.NewStyle().
		Foreground(s.Colors.Text).
		Padding(0, 0, 0, 2)

	s.List.SelectedTitleStyle = lipgloss.NewStyle().
		Border(lipgloss.Border{Left: ">"}, false, false, false, true).
		BorderForeground(s.Colors.Primary).
		Foreground(s.Colors.Primary).
		Padding(0, 0, 0, 1)

	s.LangView.Style = lipgloss.NewStyle().
		Padding(1, 2).
		Border(lipgloss.NormalBorder()).
		BorderBackground(lipgloss.Color("62"))

	s.LanguageButton.Style = lipgloss.NewStyle().
		Padding(0, 1).
		Border(lipgloss.Border{Left: ">"}, false, false, false, true).
		Foreground(s.Colors.ButtonForeground).
		Background(s.Colors.ButtonBackground)

	s.LanguageButton.HoveredStyle = lipgloss.NewStyle().
		Inherit(s.LanguageButton.Style).
		Foreground(s.Colors.ButtonHoverForeground).
		Background(s.Colors.ButtonHoverBackground)

	return &s
}
