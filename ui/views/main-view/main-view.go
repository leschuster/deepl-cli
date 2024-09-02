package mainview

import (
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/leschuster/deepl-cli/ui/com"
	formalitybtn "github.com/leschuster/deepl-cli/ui/components/button/formality-btn"
	srclangbtn "github.com/leschuster/deepl-cli/ui/components/button/src-lang-btn"
	tarlangbtn "github.com/leschuster/deepl-cli/ui/components/button/tar-lang-btn"
	translatebtn "github.com/leschuster/deepl-cli/ui/components/button/translate-btn"
	"github.com/leschuster/deepl-cli/ui/components/help"
	"github.com/leschuster/deepl-cli/ui/components/layout"
	"github.com/leschuster/deepl-cli/ui/components/textarea"
	textareadelimiter "github.com/leschuster/deepl-cli/ui/components/textarea-delimiter"
	"github.com/leschuster/deepl-cli/ui/components/topbar"
	"github.com/leschuster/deepl-cli/ui/context"
)

type Model struct {
	ctx                      *context.ProgramContext
	lay                      *layout.Layout
	insertMode               bool
	topbar                   topbar.Model
	help                     help.Model
	srcTextArea, tarTextArea textarea.Model
	textareaDelimiter        textareadelimiter.Model
}

func InitialModel(ctx *context.ProgramContext) Model {
	var srcLangBtn, tarLangBtn, formalityBtn, translateBtn layout.LayoutModel
	srcLangBtn = srclangbtn.InitialModel(ctx)
	tarLangBtn = tarlangbtn.InitialModel(ctx)
	formalityBtn = formalitybtn.InitialModel(ctx)
	translateBtn = translatebtn.InitialModel(ctx)

	srcTextArea := textarea.InitialModel(ctx, "Type to translate.", false)
	tarTextArea := textarea.InitialModel(ctx, "", true)
	delimiter := textareadelimiter.InitialModel(ctx)

	var srcTextAreaLay, tarTextAreaLay, delimiterLay layout.LayoutModel
	srcTextAreaLay = srcTextArea
	tarTextAreaLay = tarTextArea
	delimiterLay = delimiter

	// Define the general structure of the view
	lay := layout.NewLayout(
		layout.NewRow(
			layout.Fill(&srcLangBtn, layout.Left, 0.5),
			layout.Empty(), // Each row needs to have the same amount of elements
			layout.Fill(&tarLangBtn, layout.Left, 0.25),
			layout.Fill(&formalityBtn, layout.Right, 0.25),
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

	return Model{
		ctx:               ctx,
		lay:               lay,
		topbar:            topbar.InitialModel(ctx),
		help:              help.InitialModel(ctx, 8),
		srcTextArea:       srcTextArea,
		tarTextArea:       tarTextArea,
		textareaDelimiter: delimiter,
	}
}

func (m Model) Init() tea.Cmd {
	return m.lay.Init()
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		t, cmd := m.topbar.Update(msg)
		m.topbar = t.(topbar.Model)
		cmds = append(cmds, cmd)

		m.lay.Resize(msg.Width-4, msg.Height)
	case com.InsertModeEnteredMsg:
		m.insertMode = true
	case com.InsertModeExitedMsg:
		m.insertMode = false
	case tea.KeyMsg:
		h, cmd := m.help.Update(msg)
		m.help = h.(help.Model)
		cmds = append(cmds, cmd)

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

		// In the case of a key-press, we only want the
		// message to be forwarded to the active element
		cmds = append(cmds, m.lay.UpdateActive(msg))
		return m, tea.Batch(cmds...)
	}

	// In general, all components should receive the message
	cmds = append(cmds, m.lay.UpdateAll(msg))

	return m, tea.Batch(cmds...)
}

func (m Model) View() string {
	style := m.ctx.Styles.MainView.Style

	return lipgloss.JoinVertical(
		lipgloss.Left,
		m.topbar.View(),
		style.Render(m.lay.View()),
		m.help.View(),
	)
}
