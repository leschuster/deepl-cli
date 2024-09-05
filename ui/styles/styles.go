// Package styles provides most styles used in this application.

package styles

import (
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/lipgloss"
)

type Styles struct {
	// COLORS

	Colors struct {
		Primary struct {
			Foreground lipgloss.AdaptiveColor
			Background lipgloss.AdaptiveColor
		}
		Active struct {
			Foreground lipgloss.AdaptiveColor
			Background lipgloss.AdaptiveColor
		}
		Error lipgloss.AdaptiveColor
	}

	// VIEWS

	MainView struct {
		Style lipgloss.Style
	}

	LangView struct {
		Style lipgloss.Style
	}

	LoginView struct {
		Style lipgloss.Style
	}

	ErrorView struct {
		Style lipgloss.Style
	}

	// COMPONENTS

	Header struct {
		Style                       lipgloss.Style
		LeftSide, RightSide, Spacer lipgloss.Style
	}

	Textarea struct {
		Style       lipgloss.Style
		ActiveStyle lipgloss.Style
	}

	TextareaDelimiter struct {
		Style lipgloss.Style
	}

	List struct {
		Style              list.Styles
		NormalTitleStyle   lipgloss.Style
		SelectedTitleStyle lipgloss.Style
	}

	Button struct {
		Style       lipgloss.Style
		ActiveStyle lipgloss.Style
	}
}

func New() *Styles {
	s := Styles{}

	// COLORS

	s.Colors.Primary.Foreground = lipgloss.AdaptiveColor{Light: "15", Dark: "15"}
	s.Colors.Primary.Background = lipgloss.AdaptiveColor{Light: "56", Dark: "56"}

	s.Colors.Active.Foreground = lipgloss.AdaptiveColor{Light: "12", Dark: "12"}
	s.Colors.Active.Background = lipgloss.AdaptiveColor{Light: "200", Dark: "200"}

	s.Colors.Error = lipgloss.AdaptiveColor{Light: "9", Dark: "9"}

	// VIEWS

	s.MainView.Style = lipgloss.NewStyle().
		Margin(2, 2)

	s.LangView.Style = lipgloss.NewStyle().
		Padding(1, 2).
		Border(lipgloss.NormalBorder())
		//BorderBackground(lipgloss.Color("62"))

	s.LoginView.Style = lipgloss.NewStyle().
		Padding(1, 2).
		Border(lipgloss.RoundedBorder())

	s.ErrorView.Style = lipgloss.NewStyle().
		Padding(1, 2).
		Border(lipgloss.RoundedBorder()).
		Foreground(s.Colors.Error)

	// COMPONENTS

	s.Header.Style = lipgloss.NewStyle().
		MarginBottom(1)
	s.Header.LeftSide = lipgloss.NewStyle().
		Align(lipgloss.Left).
		Padding(0, 1).
		Background(s.Colors.Primary.Background).
		Foreground(s.Colors.Primary.Foreground)
	s.Header.RightSide = lipgloss.NewStyle().
		Align(lipgloss.Right).
		Padding(0, 1).
		Background(s.Colors.Primary.Background).
		Foreground(s.Colors.Primary.Foreground)
	s.Header.Spacer = lipgloss.NewStyle()

	s.Textarea.Style = lipgloss.NewStyle().
		Margin(2, 0).
		Padding(0, 0, 0, 1)
	s.Textarea.ActiveStyle = lipgloss.NewStyle().
		Border(lipgloss.HiddenBorder(), false, false, false, true).
		BorderBackground(s.Colors.Active.Background).
		Inherit(s.Textarea.Style).
		Margin(2, 0)

	s.TextareaDelimiter.Style = lipgloss.NewStyle().
		Margin(2, 0)

	s.List.Style = list.DefaultStyles()
	s.List.Style.TitleBar.Align(lipgloss.Center)
	s.List.Style.Title = lipgloss.NewStyle().
		Foreground(s.Colors.Primary.Foreground).
		Background(s.Colors.Primary.Background).
		Padding(0, 2)
	s.List.NormalTitleStyle = lipgloss.NewStyle().
		Padding(0, 0, 0, 2)
	s.List.SelectedTitleStyle = lipgloss.NewStyle().
		Border(lipgloss.Border{Left: ">"}, false, false, false, true).
		BorderForeground(s.Colors.Active.Background).
		Foreground(s.Colors.Active.Background).
		Padding(0, 0, 0, 1)

	s.Button.Style = lipgloss.NewStyle().
		Padding(0, 2).
		Foreground(s.Colors.Primary.Foreground).
		Background(s.Colors.Primary.Background)
	s.Button.ActiveStyle = lipgloss.NewStyle().
		Padding(0, 1).
		Border(lipgloss.Border{Left: ">", Right: "<"}, false, true).
		Foreground(s.Colors.Active.Foreground).
		Background(s.Colors.Active.Background).
		Inherit(s.Button.Style)

	return &s
}
