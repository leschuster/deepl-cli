package errorview

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/leschuster/deepl-cli/ui/com"
	"github.com/leschuster/deepl-cli/ui/context"
)

type Model struct {
	ctx                         *context.ProgramContext
	err                         string
	contentWidth, contentHeight int
}

func InitialModel(ctx *context.ProgramContext) Model {
	return Model{
		ctx: ctx,
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case com.ContentSizeMsg:
		m.contentWidth, m.contentHeight = msg.Width, msg.Height
	case com.Err:
		m.err = msg.Error()
	}

	return m, tea.Batch(cmds...)
}

func (m Model) View() string {
	content := lipgloss.JoinVertical(
		lipgloss.Center,

		lipgloss.NewStyle().
			Foreground(m.ctx.Styles.Colors.Error).
			Align(lipgloss.Center).
			Bold(true).
			Underline(true).
			Width(min(m.contentWidth-10, 44)).
			Render("Error:\n"),

		lipgloss.NewStyle().
			Align(lipgloss.Left).
			Render(m.err),
	)

	style := m.ctx.Styles.ErrorView.Style.Width(min(m.contentWidth-4, 50))

	return lipgloss.Place(
		m.contentWidth, m.contentHeight,
		lipgloss.Center, lipgloss.Center,
		style.Render(content),
		lipgloss.WithWhitespaceChars(" "),
	)
}
