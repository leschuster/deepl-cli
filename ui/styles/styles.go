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
		Active                 lipgloss.AdaptiveColor
		ButtonForeground       lipgloss.AdaptiveColor
		ButtonBackground       lipgloss.AdaptiveColor
		ButtonActiveForeground lipgloss.AdaptiveColor
		ButtonActiveBackground lipgloss.AdaptiveColor
	}

	Textarea struct {
		Style       lipgloss.Style
		ActiveStyle lipgloss.Style
	}

	TextareaDelimiter struct {
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

	LoginView struct {
		Style lipgloss.Style
	}

	MainView struct {
		Style lipgloss.Style
	}

	Button struct {
		Style       lipgloss.Style
		ActiveStyle lipgloss.Style
	}

	Topbar struct {
		Style                       lipgloss.Style
		LeftSide, RightSide, Spacer lipgloss.Style
	}
}

func New() *Styles {
	s := Styles{}

	s.Colors.Primary = lipgloss.AdaptiveColor{Light: "#EE6FF8", Dark: "#EE6FF8"}
	s.Colors.Active = lipgloss.AdaptiveColor{Light: "200", Dark: "200"}
	s.Colors.Text = lipgloss.AdaptiveColor{Light: "#1a1a1a", Dark: "#dddddd"}
	s.Colors.Title.Foreground = lipgloss.AdaptiveColor{Light: "62", Dark: "62"}
	s.Colors.Title.Background = lipgloss.AdaptiveColor{Light: "230", Dark: "230"}
	s.Colors.ButtonForeground = lipgloss.AdaptiveColor{Light: "3", Dark: "3"}
	s.Colors.ButtonBackground = lipgloss.AdaptiveColor{Light: "360", Dark: "360"}
	s.Colors.ButtonActiveForeground = lipgloss.AdaptiveColor{Light: "82", Dark: "82"}
	s.Colors.ButtonActiveBackground = s.Colors.Active

	s.Textarea.Style = lipgloss.NewStyle().
		Margin(2, 0)

	s.Textarea.ActiveStyle = lipgloss.NewStyle().
		Border(lipgloss.HiddenBorder(), false, false, false, true).
		BorderBackground(s.Colors.Active).
		Inherit(s.Textarea.Style).
		Margin(2, 0)

	s.TextareaDelimiter.Style = lipgloss.NewStyle().
		Margin(2, 0)

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

	s.LoginView.Style = lipgloss.NewStyle().
		Padding(1, 2).
		Border(lipgloss.RoundedBorder()).
		Foreground(s.Colors.Text)

	s.MainView.Style = lipgloss.NewStyle().
		Margin(2, 2)

	s.Button.Style = lipgloss.NewStyle().
		Padding(0, 2).
		Foreground(s.Colors.ButtonForeground).
		Background(s.Colors.ButtonBackground)

	s.Button.ActiveStyle = lipgloss.NewStyle().
		Padding(0, 1).
		Border(lipgloss.Border{Left: ">", Right: "<"}, false, true).
		Foreground(s.Colors.ButtonActiveForeground).
		Background(s.Colors.ButtonActiveBackground).
		Inherit(s.Button.Style)

	s.Topbar.Style = lipgloss.NewStyle().
		MarginBottom(1)

	s.Topbar.LeftSide = lipgloss.NewStyle().
		Align(lipgloss.Left).
		Padding(0, 1).
		Background(s.Colors.Primary).
		Foreground(s.Colors.Text)

	s.Topbar.RightSide = lipgloss.NewStyle().
		Align(lipgloss.Right).
		Padding(0, 1).
		Background(s.Colors.Primary).
		Foreground(s.Colors.Text)

	s.Topbar.Spacer = lipgloss.NewStyle()

	return &s
}
