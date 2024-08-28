package textarea

import (
	"github.com/charmbracelet/bubbles/textarea"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/leschuster/deepl-cli/ui/context"
)

type Model struct {
	ctx      *context.ProgramContext
	textarea textarea.Model
}

func InitialModel(ctx *context.ProgramContext) Model {
	ti := textarea.New()
	ti.Placeholder = "Type to translate."
	ti.Focus()

	return Model{
		ctx:      ctx,
		textarea: ti,
	}
}

func (m Model) Init() tea.Cmd {
	return textarea.Blink
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			if m.textarea.Focused() {
				m.textarea.Blur()
			}
		}
	}

	m.textarea, cmd = m.textarea.Update(msg)
	cmds = append(cmds, cmd)
	return m, tea.Batch(cmds...)
}

func (m Model) View() string {
	style := m.ctx.Styles.Textarea.Style

	return style.Render(m.textarea.View())
}

func (m *Model) Resize(width, height int) {
	m.textarea.SetWidth(width)
	m.textarea.SetHeight(height)
}
