package textareadelimiter

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/leschuster/deepl-cli/ui/components/layout"
	"github.com/leschuster/deepl-cli/ui/context"
)

type Model struct {
	ctx    *context.ProgramContext
	height int
	active bool
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
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.height = msg.Height - 20
	}
	return m, nil
}

func (m Model) View() string {
	content := strings.Trim(strings.Repeat("┃\n", m.height), "\n")
	return m.ctx.Styles.TextareaDelimiter.Style.Render(content)
}

func (m Model) IsActive() bool {
	return m.active
}

func (m Model) SetActive() layout.LayoutModel {
	m.active = true
	return m
}

func (m Model) UnsetActive() layout.LayoutModel {
	m.active = false
	return m
}

func (m Model) OnAvailWidthChange(width int) layout.LayoutModel {
	return m
}
