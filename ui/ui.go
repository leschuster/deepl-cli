package ui

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	deeplapi "github.com/leschuster/deepl-cli/pkg/deepl-api"
	"github.com/leschuster/deepl-cli/ui/components/textarea"
	"github.com/leschuster/deepl-cli/ui/context"
)

type model struct {
	ctx             *context.ProgramContext
	api             *deeplapi.DeeplAPI
	inputTextModel  textarea.Model
	outputTextModel textarea.Model
	loaded          bool
	quitting        bool
}

func initialModel(api *deeplapi.DeeplAPI) model {
	ctx := context.New()

	return model{
		ctx:             ctx,
		api:             api,
		inputTextModel:  textarea.InitialModel(ctx),
		outputTextModel: textarea.InitialModel(ctx),
	}
}

func (m model) Init() tea.Cmd {
	return tea.Batch(tea.EnterAltScreen)
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.WindowSizeMsg:
		m.ctx.ScreenWidth = msg.Width
		m.ctx.ScreenHeight = msg.Height

		m.inputTextModel.Resize(msg.Width/2-10, msg.Height/2)
		m.outputTextModel.Resize(msg.Width/2-10, msg.Height/2)

		m.loaded = true

		return m, nil

	// Is it a key press?
	case tea.KeyMsg:
		switch msg.String() {
		// Exit program
		case "ctrl+c", "q":
			m.quitting = true
			return m, tea.Quit
		}
	}

	return m, nil
}

func (m model) View() string {
	if m.quitting {
		return ""
	}
	if !m.loaded {
		return "Loading..."
	}

	return lipgloss.JoinHorizontal(
		lipgloss.Center,
		m.inputTextModel.View(),
		m.outputTextModel.View(),
	)
}

func Run(api *deeplapi.DeeplAPI) {
	p := tea.NewProgram(initialModel(api))
	if _, err := p.Run(); err != nil {
		fmt.Printf("There has been an error: %v\n", err)
		os.Exit(1)
	}
}
