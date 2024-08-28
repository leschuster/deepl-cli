package utils

import tea "github.com/charmbracelet/bubbletea"

type ErrMsg struct {
	Err error
}

func (e ErrMsg) Error() string {
	return e.Err.Error()
}

func ErrCmd(msg ErrMsg) func() tea.Msg {
	return func() tea.Msg {
		return msg
	}
}

type LoadedNewLanguagesMsg struct{}
