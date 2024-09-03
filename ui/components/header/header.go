package header

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/leschuster/deepl-cli/ui/context"
)

type Model struct {
	ctx         *context.ProgramContext
	width       int
	left, right string
}

func InitialModel(ctx *context.ProgramContext) Model {
	return Model{
		ctx:   ctx,
		left:  "DeepL CLI (Unofficial)",
		right: "API Key saved in Keychain",
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
	}
	return m, nil
}

func (m Model) View() string {
	left := m.ctx.Styles.Topbar.LeftSide.Render(m.left)
	right := m.ctx.Styles.Topbar.RightSide.Render(m.right)
	spacer := m.ctx.Styles.Topbar.Spacer.Width(
		m.width - lipgloss.Width(left) - lipgloss.Width(right),
	).Render("")

	return lipgloss.JoinHorizontal(
		lipgloss.Top,
		left,
		spacer,
		right,
	)
}
