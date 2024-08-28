package context

import (
	deeplapi "github.com/leschuster/deepl-cli/pkg/deepl-api"
	"github.com/leschuster/deepl-cli/ui/styles"
)

type ProgramContext struct {
	Styles             *styles.Styles
	ScreenWidth        int
	ScreenHeight       int
	inputText          string
	outputText         string
	SourceLanguage     *deeplapi.Language
	TargetLanguage     *deeplapi.Language
	AvailableLanguages *deeplapi.GetLanguagesResp
}

func New() *ProgramContext {
	ctx := &ProgramContext{}

	ctx.Styles = styles.New()

	return ctx
}
