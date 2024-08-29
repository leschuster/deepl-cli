package mainview

import (
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	languagebutton "github.com/leschuster/deepl-cli/ui/components/language-button"
	"github.com/leschuster/deepl-cli/ui/context"
	"github.com/leschuster/deepl-cli/ui/navigator"
)

type Model struct {
	ctx *context.ProgramContext
	nav navigator.Navigator
}

func InitialModel(ctx *context.ProgramContext) Model {
	var srcLangBtn, tarLangBtn navigator.NavModal
	srcLangBtn = languagebutton.InitialModel(ctx, "Source Language", onSrcLangBtnClick)
	tarLangBtn = languagebutton.InitialModel(ctx, "Target Language", onTarLangBtnClick)

	mat := navigator.Matrix{
		[](*navigator.NavModal){&srcLangBtn, &tarLangBtn},
	}

	nav := navigator.New(mat)

	return Model{
		ctx: ctx,
		nav: nav,
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		// TODO
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.ctx.Keys.Up):
			cmd = m.nav.Up()
			cmds = append(cmds, cmd)
		case key.Matches(msg, m.ctx.Keys.Right):
			cmd = m.nav.Right()
			cmds = append(cmds, cmd)
		case key.Matches(msg, m.ctx.Keys.Down):
			cmd = m.nav.Down()
			cmds = append(cmds, cmd)
		case key.Matches(msg, m.ctx.Keys.Left):
			cmd = m.nav.Left()
			cmds = append(cmds, cmd)
		}
	}

	cmd = m.nav.UpdateActive(msg)
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}

func (m Model) View() string {
	return m.nav.View()
}

func onSrcLangBtnClick() tea.Msg {
	return nil
}

func onTarLangBtnClick() tea.Msg {
	return nil
}
