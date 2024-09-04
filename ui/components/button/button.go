package button

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/leschuster/deepl-cli/ui/components/layout"
	"github.com/leschuster/deepl-cli/ui/context"
)

// Provides a base button compontent
type Model struct {
	ctx         *context.ProgramContext
	label, text string
	active      bool
}

// Get a new button
func InitialModel(ctx *context.ProgramContext, label, text string) Model {
	return Model{
		ctx:   ctx,
		label: label,
		text:  text,
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	//! This is a base component. To implement logic
	//! please use a wrapper component.

	return m, nil
}

func (m Model) View() string {
	fn := m.ctx.Styles.Button.Style.Render

	if m.active {
		fn = m.ctx.Styles.Button.ActiveStyle.Render
	}

	label := ""
	if m.label != "" {
		label = fmt.Sprintf("%s: ", m.label)
	}

	return fmt.Sprintf("%s%s", label, fn(m.text))
}

func (m *Model) SetLabel(label string) {
	m.label = label
}

func (m *Model) SetText(text string) {
	m.text = text
}

// Implement LayoutModel interface

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
