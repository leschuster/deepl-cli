// Package formalitybtn provides the UI button that
// redirecty the user to the formalityView

package formalitybtn

import (
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/leschuster/deepl-cli/ui/com"
	"github.com/leschuster/deepl-cli/ui/components/button"
	"github.com/leschuster/deepl-cli/ui/components/layout"
	"github.com/leschuster/deepl-cli/ui/context"
)

/*
 * This is just a tight wrapper around the base button.
 * We need to to this in order to listen to different messages
 * and output different commands. With the Layout package used,
 * we do not have outside access to the Model instances.
 */

// Button to redirect the user to the formalityView
type Model struct {
	ctx *context.ProgramContext
	btn button.Model // Just a wrapper around the base button
}

// Get a new button
func InitialModel(ctx *context.ProgramContext) Model {
	return Model{
		ctx: ctx,
		btn: button.InitialModel(ctx, "Formality", "default"),
	}
}

// Implement tea.Model interface

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case com.FormalitySelectedMsg:
		m.btn.SetText(msg.Formality)

	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.ctx.Keys.Select):
			return m, com.FormalityBtnSelectedCmd()
		}
	}

	return m, nil
}

func (m Model) View() string {
	return m.btn.View()
}

// Implement layout.LayoutModel interface

func (m Model) IsActive() bool {
	return m.btn.IsActive()
}

func (m Model) SetActive() layout.LayoutModel {
	model := m.btn.SetActive()
	m.btn = model.(button.Model)
	return m
}

func (m Model) UnsetActive() layout.LayoutModel {
	model := m.btn.UnsetActive()
	m.btn = model.(button.Model)
	return m
}

func (m Model) OnAvailWidthChange(width int) layout.LayoutModel {
	return m
}
