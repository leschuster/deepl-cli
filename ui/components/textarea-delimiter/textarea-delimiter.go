// Package textareadelimiter provides the vertical line shown between both textareas.
// It doesn't do anything and is just there for the visuals.

package textareadelimiter

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/leschuster/deepl-cli/ui/com"
	"github.com/leschuster/deepl-cli/ui/components/layout"
	"github.com/leschuster/deepl-cli/ui/context"
)

// Delimiter Model
type Model struct {
	ctx    *context.ProgramContext
	height int
	active bool
}

// Get new delimiter
func InitialModel(ctx *context.ProgramContext) Model {
	return Model{
		ctx: ctx,
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg.(type) {
	case com.ContentSizeMsg:
		m.height = m.ctx.ContentHeight - 10
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
