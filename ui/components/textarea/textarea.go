package textarea

import (
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/textarea"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/leschuster/deepl-cli/ui/context"
	"github.com/leschuster/deepl-cli/ui/navigator"
	"github.com/leschuster/deepl-cli/ui/utils"
)

type Model struct {
	ctx      *context.ProgramContext
	textarea textarea.Model
	readonly bool
	active   bool
}

func InitialModel(ctx *context.ProgramContext, placeholder string, readonly bool) Model {
	ti := textarea.New()
	ti.Placeholder = placeholder
	ti.Blur()

	return Model{
		ctx:      ctx,
		textarea: ti,
		readonly: readonly,
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.ctx.Keys.Select) && m.active:
			m.textarea.Focus()
			cmds = append(cmds, utils.EnteredInsertModeCmd)
		case key.Matches(msg, m.ctx.Keys.Unselect):
			m.textarea.Blur()
			cmds = append(cmds, utils.ExitedInsertModeCmd)
		}
	}

	m.textarea, cmd = m.textarea.Update(msg)
	cmds = append(cmds, cmd)
	return m, tea.Batch(cmds...)
}

func (m Model) View() string {
	fn := m.ctx.Styles.Textarea.Style.Render

	if m.active {
		fn = m.ctx.Styles.Textarea.ActiveStyle.Render
	}

	return fn(m.textarea.View())
}

// Implement NavModal interface
func (m Model) IsActive() bool {
	return m.active
}
func (m Model) SetActive() navigator.NavModal {
	m.active = true
	return m
}
func (m Model) UnsetActive() navigator.NavModal {
	m.active = false
	return m
}

func (m *Model) Resize(width, height int) {
	m.textarea.SetWidth(width)
	m.textarea.SetHeight(height)
}
