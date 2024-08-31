package button

import (
	"fmt"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/leschuster/deepl-cli/ui/context"
	"github.com/leschuster/deepl-cli/ui/navigator"
)

type Model struct {
	ctx    *context.ProgramContext
	label  string
	text   string
	active bool
	cmd    tea.Cmd
}

func InitialModel(ctx *context.ProgramContext, label, text string, onClick tea.Cmd) Model {
	return Model{
		ctx:   ctx,
		label: label,
		text:  text,
		cmd:   onClick,
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.ctx.Keys.Select):
			return nil, m.cmd
		}
	}

	return m, tea.Batch(cmds...)
}

func (m Model) View() string {
	fn := m.ctx.Styles.Button.Style.Render

	if m.active {
		fn = m.ctx.Styles.Button.ActiveStyle.Render
	}

	return fmt.Sprintf("%s: %s", m.label, fn(m.text))
}

func (m *Model) SetLabel(label string) {
	m.label = label
}

func (m *Model) SetText(text string) {
	m.text = text
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
