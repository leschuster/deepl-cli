package mainview

import (
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/leschuster/deepl-cli/ui/components/button"
	"github.com/leschuster/deepl-cli/ui/components/layout"
	"github.com/leschuster/deepl-cli/ui/components/textarea"
	textareadelimiter "github.com/leschuster/deepl-cli/ui/components/textarea-delimiter"
	"github.com/leschuster/deepl-cli/ui/context"
	"github.com/leschuster/deepl-cli/ui/utils"
)

type Model struct {
	ctx                      *context.ProgramContext
	lay                      *layout.Layout
	insertMode               bool
	srcTextArea, tarTextArea textarea.Model
	textareaDelimiter        textareadelimiter.Model
}

func InitialModel(ctx *context.ProgramContext) Model {
	var srcLangBtn, tarLangBtn, formalityBth layout.LayoutModel
	srcLangBtn = button.InitialModel(ctx, "Source Language", "auto", onSrcLangBtnClick)
	tarLangBtn = button.InitialModel(ctx, "Target Language", "SELECT", onTarLangBtnClick)
	formalityBth = button.InitialModel(ctx, "Formality", "default", onFormalityBtnClick)

	srcTextArea := textarea.InitialModel(ctx, "Type to translate.", false)
	tarTextArea := textarea.InitialModel(ctx, "", true)
	delimiter := textareadelimiter.InitialModel(ctx)

	var srcTextAreaLay, tarTextAreaLay, delimiterLay layout.LayoutModel
	srcTextAreaLay = srcTextArea
	tarTextAreaLay = tarTextArea
	delimiterLay = delimiter

	var translateBtn layout.LayoutModel
	translateBtn = button.InitialModel(ctx, "", "Translate", onTranslateBtnClick)

	lay, err := layout.NewLayout(
		layout.NewRow(
			layout.Fill(&srcLangBtn, layout.Left, 0.5),
			layout.Empty(),
			layout.Fill(&tarLangBtn, layout.Left, 0.25),
			layout.Fill(&formalityBth, layout.Right, 0.25),
		),
		layout.NewRow(
			layout.FillAuto(&srcTextAreaLay, layout.Left),
			layout.Fixed(&delimiterLay, layout.Center, 5).NotSelectable(),
			layout.FillAuto(&tarTextAreaLay, layout.Left),
			layout.Empty(),
		),
		layout.NewRow(
			layout.FillAuto(&translateBtn, layout.Center),
			layout.Empty(),
			layout.Empty(),
			layout.Empty(),
		),
	)
	if err != nil {
		// TODO
	}

	return Model{
		ctx:               ctx,
		lay:               lay,
		srcTextArea:       srcTextArea,
		tarTextArea:       tarTextArea,
		textareaDelimiter: delimiter,
	}
}

func (m Model) Init() tea.Cmd {
	return m.lay.Init()
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.lay.Resize(msg.Width-4, msg.Height)
		return m, m.lay.UpdateAll(msg)
	case utils.EnteredInsertMode:
		m.insertMode = true
	case utils.ExitedInsertMode:
		m.insertMode = false
	case tea.KeyMsg:
		switch {
		case m.insertMode:
			// Ignore Keystrokes
		case key.Matches(msg, m.ctx.Keys.Up):
			m.lay.NavigateUp()
		case key.Matches(msg, m.ctx.Keys.Right):
			m.lay.NavigateRight()
		case key.Matches(msg, m.ctx.Keys.Down):
			m.lay.NavigateDown()
		case key.Matches(msg, m.ctx.Keys.Left):
			m.lay.NavigateLeft()
		}
	}

	cmd = m.lay.UpdateActive(msg)
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}

func (m Model) View() string {
	style := m.ctx.Styles.MainView.Style
	return style.Render(m.lay.View())
}

func onSrcLangBtnClick() tea.Msg {
	return nil
}

func onTarLangBtnClick() tea.Msg {
	return nil
}

func onTranslateBtnClick() tea.Msg {
	return nil
}

func onFormalityBtnClick() tea.Msg {
	return nil
}
