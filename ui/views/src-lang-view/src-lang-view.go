package srclangview

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/leschuster/deepl-cli/ui/components/list"
	"github.com/leschuster/deepl-cli/ui/context"
	"github.com/leschuster/deepl-cli/ui/utils"
)

type Model struct {
	ctx  *context.ProgramContext
	list list.Model
}

func InitialModel(ctx *context.ProgramContext) Model {
	return Model{
		ctx:  ctx,
		list: list.InitialModel(ctx),
	}
}

func (m Model) Init() tea.Cmd {
	return m.ctx.AvailableLanguages.LoadInitial
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd

	switch msg.(type) {
	case utils.LoadedNewLanguagesMsg:
		langs, err := m.ctx.AvailableLanguages.GetSourceLanguages()
		if err != nil {
			cmd = utils.ErrCmd(
				utils.ErrMsg{Err: err},
			)
			return m, cmd
		}

		items := make([]list.Item, len(langs))
		for i, lang := range langs {
			items[i] = list.NewItem(lang.Name, lang.Language)
		}
		m.list.SetItems(items)
	}

	l, cmd := m.list.Update(msg)
	m.list = l.(list.Model)
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}

func (m Model) View() string {
	style := m.ctx.Styles.LangView.Style

	return style.Render(m.list.View())
}
