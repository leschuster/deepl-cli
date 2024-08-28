package languageselect

import (
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/leschuster/deepl-cli/ui/context"
)

type Model struct {
	ctx       *context.ProgramContext
	textinput textinput.Model
}

func InitialModel(ctx *context.ProgramContext, defaultInput string) Model {
	ti := textinput.New()
	ti.Placeholder = defaultInput
	ti.Width = ctx.Styles.LanguageSelect.Width
	ti.Prompt = "Hello"
	ti.PromptStyle = ctx.Styles.LanguageSelect.PromptStyle
	ti.Cursor.Style = ctx.Styles.LanguageSelect.CursorStyle
	ti.ShowSuggestions = true

	// TODO
	ti.SetSuggestions([]string{"DE", "EN", "FR"})

	return Model{
		ctx:       ctx,
		textinput: ti,
	}
}

func (m Model) Init() tea.Cmd {
	return textinput.Blink
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			if m.textinput.Focused() {
				m.textinput.Blur()
			}
		}
	}

	m.textinput, cmd = m.textinput.Update(msg)
	cmds = append(cmds, cmd)
	return m, tea.Batch(cmds...)
}

func (m Model) View() string {
	return m.textinput.View()
}
