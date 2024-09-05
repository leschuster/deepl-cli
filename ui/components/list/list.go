// Package list provides a base list component.

package list

import (
	"fmt"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/leschuster/deepl-cli/ui/context"
)

// Item represents an item in the list
type Item[T interface{}] struct {
	title, prefix string
	data          T
}

// Create a new item
func NewItem[T interface{}](title, prefix string, data T) Item[T] {
	return Item[T]{
		title:  title,
		prefix: prefix,
		data:   data,
	}
}

// Get title
func (i Item[T]) Title() string {
	return fmt.Sprintf("%s - %s", i.prefix, i.title)
}

// Get description
func (i Item[T]) Description() string {
	return "" // not used, just there to implement the list.Item interface
}

// Get data
func (i Item[T]) Data() T {
	return i.data
}

// Get prefix
func (i Item[T]) Prefix() string {
	return i.prefix
}

// Get value to filter by
func (i Item[T]) FilterValue() string {
	return fmt.Sprintf("%s - %s", i.prefix, i.title)
}

// The list model provides a base list component with the ability to filter
type Model[T interface{}] struct {
	ctx           *context.ProgramContext
	list          list.Model
	width, height int
}

// Get new list
func InitialModel[T interface{}](ctx *context.ProgramContext, title string) Model[T] {
	delegate := list.NewDefaultDelegate()
	delegate.ShowDescription = false
	delegate.Styles.NormalTitle = ctx.Styles.List.NormalTitleStyle
	delegate.Styles.SelectedTitle = ctx.Styles.List.SelectedTitleStyle

	li := list.New([]list.Item{}, delegate, 0, 0)
	li.Styles = ctx.Styles.List.Style
	li.KeyMap = ctx.Keys.ConvertToListKeyMap()
	li.Title = title

	return Model[T]{
		ctx:  ctx,
		list: li,
	}
}

// Init list
func (m Model[T]) Init() tea.Cmd {
	return nil
}

// Update list
func (m Model[T]) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd

	// Send msg to m.list
	m.list, cmd = m.list.Update(msg)
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}

// Render list
func (m Model[T]) View() string {
	items := m.list.Items()
	_ = items
	return m.list.View()
}

// Set new with and height
func (m *Model[T]) Resize(width, height int) {
	m.width, m.height = width, height
	m.list.SetWidth(width)
	m.list.SetHeight(height)
}

// Return the selected element
func (m *Model[T]) GetSelected() (*Item[T], bool) {
	i, ok := m.list.SelectedItem().(Item[T])
	if !ok {
		return nil, false
	}

	return &i, true
}

// Set list items
func (m *Model[T]) SetItems(items []Item[T]) {
	var i []list.Item
	for _, item := range items {
		i = append(i, item)
	}
	m.list.SetItems(i)
}
