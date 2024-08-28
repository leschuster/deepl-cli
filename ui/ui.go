package ui

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	deeplapi "github.com/leschuster/deepl-cli/pkg/deepl-api"
	"github.com/leschuster/deepl-cli/ui/context"
	mainview "github.com/leschuster/deepl-cli/ui/views/main-view"
)

const (
	mainViewIdx = iota
	srcLangViewIdx
	tarLangViewIdx
	loginViewIdx
)

type index int

type Model struct {
	ctx      *context.ProgramContext
	views    []tea.Model
	currView index
	loaded   bool
	quitting bool
}

func InitialModel(api *deeplapi.DeeplAPI) Model {
	ctx := context.New()
	ctx.Api = api

	// TODO
	views := []tea.Model{
		mainview.InitialModel(ctx),
	}

	return Model{
		ctx:      ctx,
		views:    views,
		currView: mainViewIdx,
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	// var cmd tea.Cmd
	// var cmds []tea.Cmd

	switch msg := msg.(type) {

	// Did the window size change?
	case tea.WindowSizeMsg:
		m.ctx.ScreenWidth = msg.Width
		m.ctx.ScreenHeight = msg.Height

		m.loaded = true

		return m, nil

	// Is it a key press?
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			m.quitting = true
			return m, tea.Quit
		}
	}

	return m, nil
}

func (m Model) View() string {
	if m.quitting {
		return ""
	}

	if !m.loaded {
		return "Loading..."
	}

	return m.views[m.currView].View()
}

func Run(api *deeplapi.DeeplAPI) {
	p := tea.NewProgram(InitialModel(api), tea.WithAltScreen())

	if _, err := p.Run(); err != nil {
		fmt.Printf("There has been an error: %v\n", err)
		os.Exit(1)
	}
}
