// Package header provides the top bar of the application.
// It displays the name of the app, the current version,
// as well as the current status (loading etc.).

package header

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/leschuster/deepl-cli/ui/com"
	"github.com/leschuster/deepl-cli/ui/context"
)

// Header model to display the top bar of the application
type Model struct {
	ctx         *context.ProgramContext
	width       int
	left, right string
	loading     bool
}

func InitialModel(ctx *context.ProgramContext) Model {
	return Model{
		ctx:   ctx,
		left:  "DeepL CLI (Unofficial)",
		right: "v0.1.2",
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
	case com.StartLoadingMsg:
		m.loading = true
	case com.StopLoadingMsg:
		m.loading = false
	}
	return m, nil
}

func (m Model) View() string {
	left := m.ctx.Styles.Header.LeftSide.Render(m.left)
	right := m.ctx.Styles.Header.RightSide.Render(m.right)

	middleContent := ""
	if m.loading {
		middleContent = " Loading..."
	}

	middle := m.ctx.Styles.Header.Spacer.
		Width(
			m.width - lipgloss.Width(left) - lipgloss.Width(right),
		).
		Align(lipgloss.Left).
		Render(middleContent)

	return lipgloss.JoinHorizontal(
		lipgloss.Top,
		left,
		middle,
		right,
	)
}
