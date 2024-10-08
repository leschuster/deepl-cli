// Package keys provides the keymap of the keys
// used in this application

package keys

import (
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
)

// Define key bindings
type KeyMap struct {
	// Select an option
	Select   key.Binding
	Unselect key.Binding

	// Navigation
	Up    key.Binding
	Right key.Binding
	Down  key.Binding
	Left  key.Binding

	// Keybindings used when browsing a list.
	NextPage    key.Binding
	PrevPage    key.Binding
	GoToStart   key.Binding
	GoToEnd     key.Binding
	Filter      key.Binding
	ClearFilter key.Binding

	// Keybindings used when setting a filter.
	CancelWhileFiltering key.Binding
	AcceptWhileFiltering key.Binding

	// Help toggle keybindings.
	ShowFullHelp  key.Binding
	CloseFullHelp key.Binding

	// The quit keybinding. This won't be caught when filtering.
	Quit key.Binding

	// The quit-no-matter-what keybinding. This will be caught when filtering.
	ForceQuit key.Binding
}

func (k KeyMap) ShortHelp() []key.Binding {
	return []key.Binding{
		k.Select, k.Unselect, k.Up, k.Down, k.Right, k.Left, k.ShowFullHelp, k.Quit,
	}
}

func (k KeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Select, k.Unselect, k.CloseFullHelp, k.Quit},
		{k.Up, k.Down, k.Right, k.Left},
		{k.NextPage, k.PrevPage, k.Filter, k.ClearFilter},
	}
}

func (k *KeyMap) ConvertToListKeyMap() list.KeyMap {
	return list.KeyMap{
		CursorUp:             k.Up,
		CursorDown:           k.Down,
		NextPage:             k.NextPage,
		PrevPage:             k.PrevPage,
		GoToStart:            k.GoToStart,
		GoToEnd:              k.GoToEnd,
		Filter:               k.Filter,
		ClearFilter:          k.ClearFilter,
		CancelWhileFiltering: k.CancelWhileFiltering,
		AcceptWhileFiltering: k.AcceptWhileFiltering,
		ShowFullHelp:         k.ShowFullHelp,
		CloseFullHelp:        k.CloseFullHelp,
		Quit:                 k.Quit,
		ForceQuit:            k.ForceQuit,
	}
}
