package layout

import tea "github.com/charmbracelet/bubbletea"

// Extends the normal tea.Model
type LayoutModel interface {
	Init() tea.Cmd
	Update(tea.Msg) (tea.Model, tea.Cmd)
	View() string

	IsActive() bool
	SetActive() LayoutModel
	UnsetActive() LayoutModel

	OnAvailWidthChange(width int) LayoutModel
}
