package srclangview

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
	list                        list.Model[deeplapi.Language]
	contentWidth, contentHeight int
}

func InitialModel(ctx *context.ProgramContext) Model {
	return Model{
		ctx:  ctx,
		list: list.InitialModel[deeplapi.Language](ctx, "Select Source Language:"),
	}
}

func (m Model) Init() tea.Cmd {
	cmds := []tea.Cmd{
		m.ctx.AvailableLanguages.LoadInitial(), // Load available languages
	}

	return tea.Batch(cmds...)
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case com.ContentSizeMsg:
		m.contentWidth, m.contentHeight = msg.Width, msg.Height
		w, h := m.calcListSize()
		m.list.Resize(w, h)
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.ctx.Keys.Select):
			// User selected a language
			item, ok := m.list.GetSelected()
			if !ok || item == nil {
				return m, nil
			}

			return m, com.SrcLangSelectedCmd((*item).Data())
		}

	case com.APILanguagesReceivedMsg:
		langs, err := m.ctx.AvailableLanguages.GetSourceLanguages()
		if err != nil {
			return m, com.ThrowErr(err)
		}

		items := make([]list.Item[deeplapi.Language], len(langs))
		for i, lang := range langs {
			items[i] = list.NewItem(lang.Name, lang.Language, lang)
		}
		m.list.SetItems(items)
	}

	l, cmd := m.list.Update(msg)
	m.list = l.(list.Model[deeplapi.Language])
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
	height = max(10, int(0.75*float32(m.contentHeight))-4)
	return
}
