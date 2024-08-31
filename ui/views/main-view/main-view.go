package mainview

import (
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/leschuster/deepl-cli/ui/components/button"
	"github.com/leschuster/deepl-cli/ui/components/textarea"
	"github.com/leschuster/deepl-cli/ui/context"
	"github.com/leschuster/deepl-cli/ui/navigator"
	"github.com/leschuster/deepl-cli/ui/utils"
)

type Model struct {
	ctx        *context.ProgramContext
	nav        navigator.Navigator
	insertMode bool
}

func InitialModel(ctx *context.ProgramContext) Model {
	var srcLangBtn, tarLangBtn navigator.NavModal
	srcLangBtn = button.InitialModel(ctx, "Source Language", "AUTO", onSrcLangBtnClick)
	tarLangBtn = button.InitialModel(ctx, "Target Language", "Select", onTarLangBtnClick)

	var srcTextArea, tarTextArea navigator.NavModal
	srcTextArea = textarea.InitialModel(ctx, "Type to translate.", false)
	tarTextArea = textarea.InitialModel(ctx, "", true)

	mat := navigator.Matrix{
		[](*navigator.NavModal){&srcLangBtn, &tarLangBtn},
		[](*navigator.NavModal){&srcTextArea, &tarTextArea},
	}

	nav := navigator.New(mat)

	return Model{
		ctx: ctx,
		nav: nav,
	}
}

func (m Model) Init() tea.Cmd {
	return m.nav.Init()
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		// TODO
	case utils.EnteredInsertMode:
		m.insertMode = true
	case utils.ExitedInsertMode:
		m.insertMode = false
	case tea.KeyMsg:
		switch {
		case m.insertMode:
			// Ignore Keystrokes
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
