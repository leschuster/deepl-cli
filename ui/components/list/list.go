package list

import (
	"fmt"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/leschuster/deepl-cli/ui/context"
)

type Item[T interface{}] struct {
	title, prefix string
	data          T
}

func NewItem[T interface{}](title, prefix string, data T) Item[T] {
	return Item[T]{
		title:  title,
		prefix: prefix,
		data:   data,
	}
}

func (i Item[T]) Title() string {
	return fmt.Sprintf("%s - %s", i.prefix, i.title)
}

func (i Item[T]) Description() string {
	return ""
}

func (i Item[T]) Data() T {
	return i.data
}

func (i Item[T]) Prefix() string {
	return i.prefix
}

func (i Item[T]) FilterValue() string {
	return fmt.Sprintf("%s - %s", i.prefix, i.title)
}

type Model[T interface{}] struct {
	ctx           *context.ProgramContext
	list          list.Model
	width, height int
}

func InitialModel[T interface{}](ctx *context.ProgramContext) Model[T] {
	delegate := list.NewDefaultDelegate()
	delegate.ShowDescription = false
	delegate.Styles.NormalTitle = ctx.Styles.List.NormalTitleStyle
	delegate.Styles.SelectedTitle = ctx.Styles.List.SelectedTitleStyle

	li := list.New([]list.Item{}, delegate, 0, 0)
	li.Styles = ctx.Styles.List.Style
	li.KeyMap = ctx.Keys.ConvertToListKeyMap()

	return Model[T]{
		ctx:  ctx,
		list: li,
	}
}

func (m Model[T]) Init() tea.Cmd {
	return nil
}

func (m Model[T]) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd

	m.list, cmd = m.list.Update(msg)
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}

func (m Model[T]) View() string {
	items := m.list.Items()
	_ = items
	return m.list.View()
}

func (m *Model[T]) Resize(width, height int) {
	m.width, m.height = width, height
	m.list.SetWidth(width)
	m.list.SetHeight(height)
}

func (m *Model[T]) GetSelected() (*Item[T], bool) {
	i, ok := m.list.SelectedItem().(Item[T])
	if !ok {
		return nil, false
	}

	return &i, true
}

func (m *Model[T]) SetItems(items []Item[T]) {
	var i []list.Item
	for _, item := range items {
		i = append(i, item)
	}
	m.list.SetItems(i)
}
