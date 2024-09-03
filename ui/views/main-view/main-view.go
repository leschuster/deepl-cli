package mainview

import (
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/leschuster/deepl-cli/ui/com"
	formalitybtn "github.com/leschuster/deepl-cli/ui/components/button/formality-btn"
	srclangbtn "github.com/leschuster/deepl-cli/ui/components/button/src-lang-btn"
	tarlangbtn "github.com/leschuster/deepl-cli/ui/components/button/tar-lang-btn"
	translatebtn "github.com/leschuster/deepl-cli/ui/components/button/translate-btn"
	"github.com/leschuster/deepl-cli/ui/components/layout"
	textareadelimiter "github.com/leschuster/deepl-cli/ui/components/textarea-delimiter"
	srctextarea "github.com/leschuster/deepl-cli/ui/components/textarea/src-textarea"
	tartextarea "github.com/leschuster/deepl-cli/ui/components/textarea/tar-textarea"
	"github.com/leschuster/deepl-cli/ui/context"
)

type Model struct {
	ctx        *context.ProgramContext
	lay        *layout.Layout
	insertMode bool
}

func InitialModel(ctx *context.ProgramContext) Model {
	var srcLangBtn, tarLangBtn, formalityBtn, translateBtn layout.LayoutModel
	srcLangBtn = srclangbtn.InitialModel(ctx)
	tarLangBtn = tarlangbtn.InitialModel(ctx)
	formalityBtn = formalitybtn.InitialModel(ctx)
	translateBtn = translatebtn.InitialModel(ctx)

	var srcTextArea, tarTextArea, delimiter layout.LayoutModel
	srcTextArea = srctextarea.InitialModel(ctx)
	tarTextArea = tartextarea.InitialModel(ctx)
	delimiter = textareadelimiter.InitialModel(ctx)

	// Define the general structure of the view
	lay := layout.NewLayout(
		layout.NewRow(
			layout.Fill(&srcLangBtn, layout.Left, 0.5),
			layout.Empty(), // Each row needs to have the same amount of elements
			layout.Fill(&tarLangBtn, layout.Left, 0.25),
			layout.Fill(&formalityBtn, layout.Right, 0.25),
		),
		layout.NewRow(
			layout.FillAuto(&srcTextArea, layout.Left),
			layout.Fixed(&delimiter, layout.Center, 5).NotSelectable(),
			layout.FillAuto(&tarTextArea, layout.Left),
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
		ctx: ctx,
		lay: lay,
	}
}

func (m Model) Init() tea.Cmd {
	return m.lay.Init()
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd

	switch msg := msg.(type) {

	case com.ContentSizeMsg:
		m.lay.Resize(msg.Width-4, msg.Height)

	case com.InsertModeEnteredMsg:
		m.insertMode = true

	case com.InsertModeExitedMsg:
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

	return style.Render(m.lay.View())
}
