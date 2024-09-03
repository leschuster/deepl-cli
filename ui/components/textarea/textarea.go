package textarea

import (
	"github.com/charmbracelet/bubbles/textarea"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/leschuster/deepl-cli/ui/com"
	"github.com/leschuster/deepl-cli/ui/components/layout"
	"github.com/leschuster/deepl-cli/ui/context"
)

type Model struct {
	ctx      *context.ProgramContext
	textarea textarea.Model
	active   bool
}

func InitialModel(ctx *context.ProgramContext, placeholder string) Model {
	ti := textarea.New()
	ti.Placeholder = placeholder
	ti.Prompt = ""
	ti.Blur()

	return Model{
		ctx:      ctx,
		textarea: ti,
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case com.ContentSizeMsg:
		textareaHeight := msg.Height - 11
		m.textarea.SetHeight(textareaHeight)
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
	m.textarea.SetWidth(width - 10)
	return m
}

func (m *Model) SetPlaceholder(text string) {
	m.textarea.Placeholder = text
}

func (m *Model) Focus() {
	m.textarea.Focus()
}

func (m *Model) Blur() {
	m.textarea.Blur()
}

func (m *Model) Value() string {
	return m.textarea.Value()
}

func (m *Model) SetValue(text string) {
	m.textarea.SetValue(text)
}
