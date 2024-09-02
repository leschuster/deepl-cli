package loginview

import (
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/leschuster/deepl-cli/ui/context"
)

type Model struct {
	ctx   *context.ProgramContext
	input textinput.Model
}

func InitialModel(ctx *context.ProgramContext) Model {
	ti := textinput.New()
	ti.Placeholder = "API key"
	ti.Focus()

	return Model{
		ctx:   ctx,
		input: ti,
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd

	switch msg.(type) {
	case tea.WindowSizeMsg:

	}

	ti, cmd := m.input.Update(msg)
	m.input = ti
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}

func (m Model) View() string {
	style := m.ctx.Styles.LoginView.Style

	return style.Render(m.input.View())
}
