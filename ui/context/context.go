package context

import deeplapi "github.com/leschuster/deepl-cli/pkg/deepl-api"

type ProgramContext struct {
	ScreenWidth        int
	ScreenHeight       int
	inputText          string
	outputText         string
	SourceLanguage     *deeplapi.Language
	TargetLanguage     *deeplapi.Language
	AvailableLanguages *deeplapi.GetLanguagesResp
}
