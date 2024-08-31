package utils

import (
	tea "github.com/charmbracelet/bubbletea"
	deeplapi "github.com/leschuster/deepl-cli/pkg/deepl-api"
)

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

type SrcLangSelected struct {
	Language deeplapi.Language
}

type EnteredInsertMode struct{}

func EnteredInsertModeCmd() tea.Msg {
	return EnteredInsertMode{}
}

type ExitedInsertMode struct{}

func ExitedInsertModeCmd() tea.Msg {
	return ExitedInsertMode{}
}
