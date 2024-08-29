package navigator

import tea "github.com/charmbracelet/bubbletea"

type NavModal interface {
	Init() tea.Cmd
	Update(tea.Msg) (tea.Model, tea.Cmd)
	View() string
	IsActive() bool
	SetActive() NavModal
	UnsetActive() NavModal
}
