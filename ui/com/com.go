// Provides a set of internally used commands and messages
// for the Bubbletea framework

package com

import (
	tea "github.com/charmbracelet/bubbletea"
	deeplapi "github.com/leschuster/deepl-cli/pkg/deepl-api"
)

// Represents the event that an error occured
type Err struct {
	Err error
}

// Get error message as string
func (e Err) Error() string {
	if e.Err == nil {
		return ""
	}
	return e.Err.Error()
}

// tea command to throw an error
func ThrowErr(err error) func() tea.Msg {
	return func() tea.Msg {
		return Err{
			Err: err,
		}
	}
}

// Describe the action that the user entered an API key
type APIKeyEnteredMsg struct {
	Key string
}

// Command to trigger APIKeyEntered
func APIKeyEnteredCmd(key string) func() tea.Msg {
	return func() tea.Msg {
		return APIKeyEnteredMsg{
			Key: key,
		}
	}
}

// Describes that the size available to a view changed
// Difference to tea.WindowSizeMsg: this one contains the screen height substracted
// by header and footer height, and the screen width substracted by global margins
// The new size is accessible to the context.
type ContentSizeMsg struct{}

// Command to trogger ContentSizeMsg
func ContentSizeCmd() func() tea.Msg {
	return func() tea.Msg {
		return ContentSizeMsg{}
	}
}

// Describes the action the of the user selecting a source language
type SrcLangSelectedMsg struct {
	Language deeplapi.Language
}

// Command to trigger SrcLangSelected
func SrcLangSelectedCmd(language deeplapi.Language) func() tea.Msg {
	return func() tea.Msg {
		return SrcLangSelectedMsg{
			Language: language,
		}
	}
}

// Describes the action of the user selecting a target language
type TarLangSelectedMsg struct {
	Language deeplapi.Language
}

// Command to trigger TarLangSelected
func TarLangSelectedCmd(language deeplapi.Language) func() tea.Msg {
	return func() tea.Msg {
		return TarLangSelectedMsg{
			Language: language,
		}
	}
}

// Describes the action of the user selecting a formality
// for the translation
type FormalitySelectedMsg struct {
	Formality string
}

// Command to trigger FormalitySelected
func FormalitySelectedCmd(formality string) func() tea.Msg {
	return func() tea.Msg {
		return FormalitySelectedMsg{
			Formality: formality,
		}
	}
}

// Describes the action of the user selecting the source language button
type SrcLangBtnSelectedMsg struct{}

// Command to trigger SrcLangBtnSelected
func SrcLangBtnSelectedCmd() func() tea.Msg {
	return func() tea.Msg {
		return SrcLangBtnSelectedMsg{}
	}
}

// Describes the action of the user selecting the target language button
type TarLangBtnSelectedMsg struct{}

// Command to trigger TarLangBtnSelected
func TarLangBtnSelectedCmd() func() tea.Msg {
	return func() tea.Msg {
		return TarLangBtnSelectedMsg{}
	}
}

// Describes the action of the user selecting the formality button
type FormalityBtnSelectedMsg struct{}

// Command to trigger FormalityBtnSelected
func FormalityBtnSelectedCmd() func() tea.Msg {
	return func() tea.Msg {
		return FormalityBtnSelectedMsg{}
	}
}

// Describes the action of the user selecting the translate button
type TranslateBtnSelectedMsg struct{}

// Command to trigger TranslateBtnSelected
func TranslateBtnSelectedCmd() func() tea.Msg {
	return func() tea.Msg {
		return TranslateBtnSelectedMsg{}
	}
}

// Describes the action of the user entering insert mode
// All navigational inputs shall be ignored for the time being
type InsertModeEnteredMsg struct{}

// Command to trigger insert mode
func InsertModeEnteredCmd() func() tea.Msg {
	return func() tea.Msg {
		return InsertModeEnteredMsg{}
	}
}

// Describes the action of the user exiting insert mode
type InsertModeExitedMsg struct{}

// Command to exit insert mode
func InsertModeExitedCmd() func() tea.Msg {
	return func() tea.Msg {
		return InsertModeExitedMsg{}
	}
}

// Previously requested available languages have been received
type APILanguagesReceivedMsg struct{}

// Command to trigger APILanguagesReceived
func APILanguagesReceivedCmd() func() tea.Msg {
	return func() tea.Msg {
		return APILanguagesReceivedMsg{}
	}
}

// Previously requested translation has been received
type APITranslationReceivedMsg struct{}

// Command to trigger APITranslationReceived
func APITranslationReceivedCmd() func() tea.Msg {
	return func() tea.Msg {
		return APITranslationReceivedMsg{}
	}
}

// Describes that a loading process started in the background
type StartLoadingMsg struct{}

// Command to trigger StartLoading
func StartLoadingCmd() func() tea.Msg {
	return func() tea.Msg {
		return StartLoadingMsg{}
	}
}

// Describes that a loading process stopped in the background
type StopLoadingMsg struct{}

// Command to trigger StopLoading
func StopLoadingCmd() func() tea.Msg {
	return func() tea.Msg {
		return StopLoadingMsg{}
	}
}
