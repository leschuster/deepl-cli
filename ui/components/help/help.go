package help

import (
	"strings"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/leschuster/deepl-cli/ui/context"
)

type Model struct {
	ctx    *context.ProgramContext
	help   help.Model
	height int
}

func InitialModel(ctx *context.ProgramContext, height int) Model {
	return Model{
		ctx:    ctx,
		help:   help.New(),
		height: height,
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.ctx.Keys.ShowFullHelp):
			m.help.ShowAll = true
		case key.Matches(msg, m.ctx.Keys.CloseFullHelp):
			m.help.ShowAll = false
		}
	}
	return m, nil
}

func (m Model) View() string {
	helpView := m.help.View(m.ctx.Keys)
	height := m.height - strings.Count(helpView, "\n")

	return "\n" + strings.Repeat("\n", height) + helpView
}
