// Package ui provides the entry point for the user interface.

package ui

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/leschuster/deepl-cli/pkg/auth"
	deeplapi "github.com/leschuster/deepl-cli/pkg/deepl-api"
	"github.com/leschuster/deepl-cli/ui/com"
	"github.com/leschuster/deepl-cli/ui/components/header"
	"github.com/leschuster/deepl-cli/ui/components/help"
	"github.com/leschuster/deepl-cli/ui/context"
	errorview "github.com/leschuster/deepl-cli/ui/views/error-view"
	formalityview "github.com/leschuster/deepl-cli/ui/views/formality-view"
	loginview "github.com/leschuster/deepl-cli/ui/views/login-view"
	mainview "github.com/leschuster/deepl-cli/ui/views/main-view"
	srclangview "github.com/leschuster/deepl-cli/ui/views/src-lang-view"
	tarlangview "github.com/leschuster/deepl-cli/ui/views/tar-lang-view"
)

type ViewIdx int

const (
	mainViewIdx ViewIdx = iota
	srcLangViewIdx
	tarLangViewIdx
	formalityViewIdx
	loginViewIdx
	errorViewIdx
)

// Configure some layout constants
const (
	headerHeight = 2
	helpHeight   = 6
)

// The ui model is at the root of the application.
// It is responsible for managing different views
// and rendering the header and help.
type Model struct {
	auth     auth.Auth
	ctx      *context.ProgramContext
	views    []tea.Model
	currView ViewIdx
	loaded   bool
	quitting bool
	header   header.Model
	help     help.Model
}

// Get a new ui model
func InitialModel(auth auth.Auth) Model {
	ctx := context.New()

	// Setup available views
	views := []tea.Model{
		mainview.InitialModel(ctx),
		srclangview.InitialModel(ctx),
		tarlangview.InitialModel(ctx),
		formalityview.InitialModel(ctx),
		loginview.InitialModel(ctx),
		errorview.InitialModel(ctx),
	}

	currView := mainViewIdx

	if apiKey, err := auth.GetAPIKey(); err == nil {
		// User is already signed in
		ctx.Api = deeplapi.New(apiKey)
	} else {
		// User is not signed in
		// Redirect to login view
		currView = loginViewIdx
	}

	return Model{
		auth:     auth,
		ctx:      ctx,
		views:    views,
		currView: currView,
		header:   header.InitialModel(ctx),
		help:     help.InitialModel(ctx, helpHeight),
	}
}

// Initialize ui model
func (m Model) Init() tea.Cmd {
	cmds := []tea.Cmd{
		tea.SetWindowTitle("DeepL CLI (Unofficial)"), // Set Title
		m.views[m.currView].Init(),                   // Initialize active view
	}

	return tea.Batch(cmds...)
}

// Update ui model
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd

	switch msg := msg.(type) {

	// Did an error occur?
	case com.Err:
		// Redirect to errorView
		m.currView = errorViewIdx

	// Did the window size change?
	case tea.WindowSizeMsg:
		m.ctx.ScreenWidth = msg.Width
		m.ctx.ScreenHeight = msg.Height

		// content size is the available size for a view
		m.ctx.ContentWidth = msg.Width
		m.ctx.ContentHeight = msg.Height - headerHeight - helpHeight - 1

		// Pass on to header
		h, cmd := m.header.Update(msg)
		m.header = h.(header.Model)
		cmds = append(cmds, cmd)

		// We throw a separate contentSizeMsg that other models will react on
		cmds = append(cmds, com.ContentSizeCmd())

		return m, tea.Batch(cmds...)

	// Did the content size change?
	case com.ContentSizeMsg:
		// This message gets thrown after a tea.WindowSizeMsg
		// We catch it here to distribute it among ALL views,
		// not just the active one

		for i, view := range m.views {
			model, cmd := view.Update(msg)
			m.views[i] = model
			cmds = append(cmds, cmd)
		}

		// UI is loaded
		m.loaded = true

		return m, tea.Batch(cmds...)

	// Did the user enter an API key?
	case com.APIKeyEnteredMsg:
		m.ctx.Api = deeplapi.New(msg.Key)

		// Switch to main view
		m.currView = mainViewIdx
		cmds = append(cmds, m.views[m.currView].Init())

		// Define a command to save apikey locally
		// Bubbletea will run it asynchronously
		cmd = func() tea.Msg {
			err := m.auth.SetApiKey(msg.Key)
			if err != nil {
				return com.ThrowErr(err)
			}
			return nil
		}
		cmds = append(cmds, cmd)

		// Exit early so that no other component receives this message
		return m, tea.Batch(cmds...)

	// Is it a key press?
	case tea.KeyMsg:
		// Pass it on to help
		h, cmd := m.help.Update(msg)
		m.help = h.(help.Model)
		cmds = append(cmds, cmd)

		switch {
		case key.Matches(msg, m.ctx.Keys.Quit) && !m.ctx.InsertMode:
			fallthrough
		case key.Matches(msg, m.ctx.Keys.ForceQuit):
			m.quitting = true
			return m, tea.Quit
		}

	// Did the available languages request complete?
	case com.APILanguagesReceivedMsg:
		cmds = append(cmds, com.StopLoadingCmd())

	// Did the translation request complete?
	case com.APITranslationReceivedMsg:
		cmds = append(cmds, com.StopLoadingCmd())

	// Did the user press the source language button?
	case com.SrcLangBtnSelectedMsg:
		m.currView = srcLangViewIdx
		return m, m.views[m.currView].Init()

	// Did the user select a source language?
	case com.SrcLangSelectedMsg:
		m.ctx.SourceLanguage = &msg.Language
		m.currView = mainViewIdx

	// Did the user press the target language button?
	case com.TarLangBtnSelectedMsg:
		m.currView = tarLangViewIdx
		return m, m.views[m.currView].Init()

	// Did the user select a target language?
	case com.TarLangSelectedMsg:
		m.ctx.TargetLanguage = &msg.Language
		m.currView = mainViewIdx

	// Did the user press the formality button?
	case com.FormalityBtnSelectedMsg:
		m.currView = formalityViewIdx
		return m, m.views[m.currView].Init()

	// Did the user select a formality?
	case com.FormalitySelectedMsg:
		m.ctx.Formality = msg.Formality
		m.currView = mainViewIdx

	// Did we enter insert mode?
	case com.InsertModeEnteredMsg:
		m.ctx.InsertMode = true

	// Did we exit insert mode?
	case com.InsertModeExitedMsg:
		m.ctx.InsertMode = false

	// Did the user press the translate button?
	case com.TranslateBtnSelectedMsg:
		if m.ctx.Api == nil {
			return m, com.ThrowErr(fmt.Errorf("ctx.api is nil"))
		}

		cmds = append(cmds, com.StartLoadingCmd())

		// Define a command that will fetch the translation
		// We return this command because Bubbletea handles
		// commands asynchronously
		cmd = func() tea.Msg {
			srcLang := "" // if empty, DeepL will try to detect it
			if m.ctx.SourceLanguage != nil {
				srcLang = m.ctx.SourceLanguage.Language
			}

			if m.ctx.TargetLanguage == nil {
				return com.Err{
					Err: fmt.Errorf("no target language selected"),
				}
			}
			tarLang := m.ctx.TargetLanguage.Language

			formality := ""
			if m.ctx.TargetLanguage.SupportsFormality {
				formality = m.ctx.Formality
			}

			params := deeplapi.TranslateParams{
				Text:       []string{m.ctx.SourceText},
				SourceLang: srcLang,
				TargetLang: tarLang,
				Context:    "",
				Formality:  formality,
			}

			resp, err := m.ctx.Api.Translate(params)
			if err != nil {
				return com.Err{
					Err: fmt.Errorf("failed to fetch translation: %v", err),
				}
			}

			m.ctx.TranslationResult = resp

			return com.APITranslationReceivedMsg{}
		}
		cmds = append(cmds, cmd)

		return m, tea.Batch(cmds...)
	}

	// Pass msg to header
	headerModel, cmd := m.header.Update(msg)
	cmds = append(cmds, cmd)
	m.header = headerModel.(header.Model)

	// Pass msg to activ view
	model, cmd := m.views[m.currView].Update(msg)
	cmds = append(cmds, cmd)
	m.views[m.currView] = model

	// Pass msg to help
	helpModel, cmd := m.help.Update(msg)
	cmds = append(cmds, cmd)
	m.help = helpModel.(help.Model)

	return m, tea.Batch(cmds...)
}

// View ui model
func (m Model) View() string {
	if m.quitting {
		// Clear screen before exiting the program
		return ""
	}

	if !m.loaded {
		return "Loading..."
	}

	return lipgloss.JoinVertical(
		lipgloss.Left,
		m.header.View(),
		m.views[m.currView].View(),
		m.help.View(),
	)
}

// Start the application and show the user interface
func Run(auth auth.Auth) {
	// Create a new program occupying the whole screen
	p := tea.NewProgram(InitialModel(auth), tea.WithAltScreen())

	if _, err := p.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "There has been an error: %v\n", err)
		os.Exit(1)
	}
}
