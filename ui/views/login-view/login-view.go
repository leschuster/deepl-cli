package loginview

import (
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/leschuster/deepl-cli/ui/com"
	"github.com/leschuster/deepl-cli/ui/context"
)

type Model struct {
	ctx                         *context.ProgramContext
	input                       textinput.Model
	contentWidth, contentHeight int
}

func InitialModel(ctx *context.ProgramContext) Model {
	ti := textinput.New()
	ti.Placeholder = "Please enter your DeepL API key"
	ti.Focus()

	return Model{
		ctx:   ctx,
		input: ti,
	}
}

func (m Model) Init() tea.Cmd {
	return com.InsertModeEnteredCmd()
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case com.ContentSizeMsg:
		m.contentWidth, m.contentHeight = msg.Width, msg.Height
		m.input.Width = min(m.contentWidth-4, 50)
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.ctx.Keys.Select):
			apiKey := m.input.Value()
			_ = apiKey

			cmds = append(cmds, com.InsertModeExitedCmd())
		}
	}

	ti, cmd := m.input.Update(msg)
	m.input = ti
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}

func (m Model) View() string {
	inputRendered := m.input.View()

	style := m.ctx.Styles.LoginView.Style.Width(lipgloss.Width(inputRendered) + 4)

	content := lipgloss.JoinVertical(
		lipgloss.Left,
		inputRendered,
		"\nHint: The key will be saved in your systems keyring",
	)

	return lipgloss.Place(
		m.contentWidth, m.contentHeight,
		lipgloss.Center, lipgloss.Center,
		style.Render(content),
		lipgloss.WithWhitespaceChars(" "),
	)
}
