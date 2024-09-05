// Package srctextarea provides the text area for the user to type in
// the text that should be translated.

package srctextarea

import (
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/leschuster/deepl-cli/ui/com"
	"github.com/leschuster/deepl-cli/ui/components/layout"
	"github.com/leschuster/deepl-cli/ui/components/textarea"
	"github.com/leschuster/deepl-cli/ui/context"
)

/*
 * This is just a tight wrapper around the base textarea component.
 * We need to to this in order to listen to different messages
 * and output different commands. With the Layout package used,
 * we do not have outside access to the Model instances.
 */

// Provides the textarea where the user can type the text,
// that should be translated
type Model struct {
	ctx        *context.ProgramContext
	textarea   textarea.Model
	insertMode bool
}

func InitialModel(ctx *context.ProgramContext) Model {
	return Model{
		ctx:      ctx,
		textarea: textarea.InitialModel(ctx, "Type to Translate."),
	}
}

// Implement tea.Model interface

func (m Model) Init() tea.Cmd {
	return m.textarea.Init()
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd

	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch {

		// User can start to type after entering insert mode
		case key.Matches(msg, m.ctx.Keys.Select) && m.textarea.IsActive() && !m.insertMode:
			m.textarea.Focus()
			m.insertMode = true
			cmds = append(cmds, com.InsertModeEnteredCmd())

			// Return early because the textarea shall
			// not receive the 'enter' key that activated insert mode
			return m, tea.Batch(cmds...)

		// User can no longer type after exiting insert mode
		// Saving the users text
		case key.Matches(msg, m.ctx.Keys.Unselect):
			m.textarea.Blur()
			m.insertMode = false
			m.ctx.SourceText = m.textarea.Value()
			cmds = append(cmds, com.InsertModeExitedCmd())
		}

	}

	ta, cmd := m.textarea.Update(msg)
	m.textarea = ta.(textarea.Model)
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}

func (m Model) View() string {
	return m.textarea.View()
}

// Implement layout.LayoutModel interface

func (m Model) IsActive() bool {
	return m.textarea.IsActive()
}
func (m Model) SetActive() layout.LayoutModel {
	mod := m.textarea.SetActive()
	m.textarea = mod.(textarea.Model)
	return m
}
func (m Model) UnsetActive() layout.LayoutModel {
	mod := m.textarea.UnsetActive()
	m.textarea = mod.(textarea.Model)
	return m
}
func (m Model) OnAvailWidthChange(width int) layout.LayoutModel {
	mod := m.textarea.OnAvailWidthChange(width)
	m.textarea = mod.(textarea.Model)
	return m
}
