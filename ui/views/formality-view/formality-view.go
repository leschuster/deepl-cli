// Package formalityview provides the view where the user is able to select a formality.

package formalityview

import (
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	deeplapi "github.com/leschuster/deepl-cli/pkg/deepl-api"
	"github.com/leschuster/deepl-cli/ui/com"
	"github.com/leschuster/deepl-cli/ui/components/list"
	"github.com/leschuster/deepl-cli/ui/context"
)

type Model struct {
	ctx                         *context.ProgramContext
	list                        list.Model[string]
	contentWidth, contentHeight int
}

func InitialModel(ctx *context.ProgramContext) Model {
	li := list.InitialModel[string](ctx, "Select Formality:")

	formalities := []string{
		deeplapi.FormalityLess,
		deeplapi.FormalityPreferLess,
		deeplapi.FormalityDefault,
		deeplapi.FormalityPreferMore,
		deeplapi.FormalityMore,
	}

	items := make([]list.Item[string], len(formalities))
	for i, form := range formalities {
		items[i] = list.NewItem(form, "", form)
	}

	li.SetItems(items)

	return Model{
		ctx:  ctx,
		list: li,
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case com.ContentSizeMsg:
		m.contentWidth, m.contentHeight = m.ctx.ContentWidth, m.ctx.ContentHeight
		w, h := m.calcListSize()
		m.list.Resize(w, h)
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.ctx.Keys.Select):
			// User selected a formality
			item, ok := m.list.GetSelected()
			if !ok || item == nil {
				return m, nil
			}

			return m, com.FormalitySelectedCmd((*item).Data())
		}
	}

	l, cmd := m.list.Update(msg)
	m.list = l.(list.Model[string])
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}

func (m Model) View() string {
	style := m.ctx.Styles.LangView.Style

	content := style.Render(m.list.View())

	// Place content in the center of the screen
	return lipgloss.Place(
		m.contentWidth, m.contentHeight,
		lipgloss.Center, lipgloss.Center,
		content,
		lipgloss.WithWhitespaceChars(" "),
	)
}

func (m *Model) calcListSize() (width, height int) {
	width = min(42, m.contentWidth)
	height = max(20, int(0.75*float32(m.contentHeight))-4)
	return
}
