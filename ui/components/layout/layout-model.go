package layout

import tea "github.com/charmbracelet/bubbletea"

// Extends the normal tea.Model
// with methods set and unset it as active element.
type LayoutModel interface {
	// tea.Model
	Init() tea.Cmd
	Update(tea.Msg) (tea.Model, tea.Cmd)
	View() string

	IsActive() bool
	SetActive() LayoutModel
	UnsetActive() LayoutModel

	// This method gets called with the actual with the component
	// will have available. This is useful when you use fill-mode
	// and do not have a fixed width.
	OnAvailWidthChange(width int) LayoutModel
}
