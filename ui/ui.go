package ui

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	deeplapi "github.com/leschuster/deepl-cli/pkg/deepl-api"
	"github.com/leschuster/deepl-cli/ui/context"
	"github.com/leschuster/deepl-cli/ui/utils"
	mainview "github.com/leschuster/deepl-cli/ui/views/main-view"
	srclangview "github.com/leschuster/deepl-cli/ui/views/src-lang-view"
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
	err      error
	loaded   bool
	quitting bool
}

func InitialModel(api *deeplapi.DeeplAPI) Model {
	ctx := context.New(api)

	// TODO
	views := []tea.Model{
		mainview.InitialModel(ctx),
		srclangview.InitialModel(ctx),
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
	var cmd tea.Cmd
	var cmds []tea.Cmd

	switch msg := msg.(type) {

	// Did a custom error occur?
	case utils.ErrMsg:
		m.err = msg.Err
		return m, tea.Quit

	// Did the window size change?
	case tea.WindowSizeMsg:
		m.ctx.ScreenWidth = msg.Width
		m.ctx.ScreenHeight = msg.Height

		for i, view := range m.views {
			model, cmd := view.Update(msg)
			m.views[i] = model
			cmds = append(cmds, cmd)
		}

		m.loaded = true

		return m, tea.Batch(cmds...)

	// Is it a key press?
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			m.quitting = true
			return m, tea.Quit
		case "1":
			m.currView = mainViewIdx

			cmds = append(cmds, m.views[m.currView].Init())

			return m, tea.Batch(cmds...)
		case "2":
			m.currView = srcLangViewIdx
			return m, m.views[m.currView].Init()
		}

	// Did the user select a source language?
	case utils.SrcLangSelected:
		m.ctx.SourceLanguage = &msg.Language
		m.currView = mainViewIdx
	}

	model, cmd := m.views[m.currView].Update(msg)
	cmds = append(cmds, cmd)
	m.views[m.currView] = model

	return m, tea.Batch(cmds...)
}

func (m Model) View() string {
	if m.err != nil {
		return m.err.Error()
	}

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
