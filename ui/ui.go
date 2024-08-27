package ui

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	deeplapi "github.com/leschuster/deepl-cli/pkg/deepl-api"
)

type model struct {
	loading            bool
	inputText          string
	outputText         string
	srcLang            *deeplapi.Language
	tarLang            *deeplapi.Language
	availableLanguages *deeplapi.GetLanguagesResp
	api                *deeplapi.DeeplAPI
}

func initMainModel(api *deeplapi.DeeplAPI) model {
	return model{
		loading:            false,
		inputText:          "",
		outputText:         "",
		srcLang:            nil,
		tarLang:            nil,
		availableLanguages: nil,
		api:                api,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	// Is it a key press?
	case tea.KeyMsg:
		switch msg.String() {
		// Exit program
		case "ctrl+c", "q":
			return m, tea.Quit
		}
	}

	return m, nil
}

func (m model) View() string {
	return "DeepL CLI"
}

func Run(api *deeplapi.DeeplAPI) {
	p := tea.NewProgram(initMainModel(api))
	if _, err := p.Run(); err != nil {
		fmt.Printf("There has been an error: %v\n", err)
		os.Exit(1)
	}
}
