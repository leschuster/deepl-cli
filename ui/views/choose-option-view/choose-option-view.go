package chooseoptionview

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/leschuster/deepl-cli/ui/components/list"
	"github.com/leschuster/deepl-cli/ui/context"
)

type Model struct {
	ctx  *context.ProgramContext
	list list.Model
}

func InitialModel(ctx *context.ProgramContext) Model {
	li := list.InitialModel(ctx)

	return Model{
		ctx:  ctx,
		list: li,
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
		switch msg.String() {
		case "ctrl+c", "q":
			// TODO QuitCMD
			return m, tea.Quit
		}
	}

	l, cmd := m.list.Update(msg)
	m.list = l.(list.Model)
	cmds = append(cmds, cmd)
	return m, tea.Batch(cmds...)
}

func (m Model) View() string {
	style := m.ctx.Styles.ChooseOptionView.Style

	return style.Render(m.list.View())
}
