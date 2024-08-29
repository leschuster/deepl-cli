package list

import (
	"fmt"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/leschuster/deepl-cli/ui/context"
)

type Item struct {
	title, prefix string
}

func NewItem(title, prefix string) Item {
	return Item{
		title:  title,
		prefix: prefix,
	}
}

func (i Item) Title() string {
	return fmt.Sprintf("%s - %s", i.prefix, i.title)
}

func (i Item) Description() string {
	return ""
}

func (i Item) Prefix() string {
	return i.prefix
}

func (i Item) FilterValue() string {
	return fmt.Sprintf("%s - %s", i.prefix, i.title)
}

type Model struct {
	ctx           *context.ProgramContext
	list          list.Model
	width, height int
}

func InitialModel(ctx *context.ProgramContext) Model {
	delegate := list.NewDefaultDelegate()
	delegate.ShowDescription = false
	delegate.Styles.NormalTitle = ctx.Styles.List.NormalTitleStyle
	delegate.Styles.SelectedTitle = ctx.Styles.List.SelectedTitleStyle

	li := list.New([]list.Item{}, delegate, 0, 0)
	li.Styles = ctx.Styles.List.Style

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
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m Model) View() string {
	items := m.list.Items()
	_ = items
	return m.list.View()
}

func (m *Model) Resize(width, height int) {
	m.width, m.height = width, height
	m.list.SetWidth(width)
	m.list.SetHeight(height)
}

func (m *Model) GetSelected() (string, error) {
	i, ok := m.list.SelectedItem().(Item)
	if !ok {
		return "", fmt.Errorf("could not get selected item")
	}

	return i.Title(), nil
}

func (m *Model) SetItems(items []Item) {
	var i []list.Item
	for _, item := range items {
		i = append(i, item)
	}
	m.list.SetItems(i)
}
