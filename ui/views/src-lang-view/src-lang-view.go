package srclangview

import (
	"strings"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	deeplapi "github.com/leschuster/deepl-cli/pkg/deepl-api"
	"github.com/leschuster/deepl-cli/ui/components/list"
	"github.com/leschuster/deepl-cli/ui/context"
	"github.com/leschuster/deepl-cli/ui/utils"
)

type Model struct {
	ctx  *context.ProgramContext
	list list.Model[deeplapi.Language]
}

func InitialModel(ctx *context.ProgramContext) Model {
	return Model{
		ctx:  ctx,
		list: list.InitialModel[deeplapi.Language](ctx),
	}
}

func (m Model) Init() tea.Cmd {
	cmds := []tea.Cmd{
		m.ctx.AvailableLanguages.LoadInitial, // Load available languages
	}

	return tea.Batch(cmds...)
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.list.Resize(m.calcListSize())
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.ctx.Keys.Select):
			return m, m.onSelect
		}

	case utils.LoadedNewLanguagesMsg:
		langs, err := m.ctx.AvailableLanguages.GetSourceLanguages()
		if err != nil {
			cmd = utils.ErrCmd(
				utils.ErrMsg{Err: err},
			)
			return m, cmd
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
	doc := strings.Builder{}
	style := m.ctx.Styles.LangView.Style

	doc.WriteString(
		style.Render(m.list.View()),
	)

	return lipgloss.Place(
		m.ctx.ScreenWidth, m.ctx.ScreenHeight,
		lipgloss.Center, lipgloss.Center,
		doc.String(),
	)
}

func (m *Model) onSelect() tea.Msg {
	item, ok := m.list.GetSelected()
	if !ok {
		return nil
	}

	return utils.SrcLangSelected{
		Language: item.Data(),
	}
}

func (m *Model) calcListSize() (width, height int) {
	width = min(42, m.ctx.ScreenWidth)
	height = max(10, int(0.75*float32(m.ctx.ScreenHeight))-4)
	return
}
