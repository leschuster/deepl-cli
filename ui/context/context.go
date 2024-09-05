// Package context provides a way of sharing data and functionality
// between all models of the application.

package context

import (
	deeplapi "github.com/leschuster/deepl-cli/pkg/deepl-api"
	"github.com/leschuster/deepl-cli/ui/keys"
	"github.com/leschuster/deepl-cli/ui/styles"
	"github.com/leschuster/deepl-cli/ui/utils"
)

type ProgramContext struct {
	Api                            *deeplapi.DeeplAPI
	Keys                           keys.KeyMap
	ScreenWidth, ScreenHeight      int // Size of entire screen
	ContentWidth, ContentHeight    int // Size of the space that is available to a view
	Styles                         *styles.Styles
	SourceLanguage, TargetLanguage *deeplapi.Language
	SourceText                     string
	Formality                      string
	TranslationResult              *deeplapi.TranslateResp
	AvailableLanguages             utils.AvailableLanguages
	InsertMode                     bool
}

func New() *ProgramContext {
	return &ProgramContext{
		Keys:               keys.DefaultKeyMap(),
		Styles:             styles.New(),
		AvailableLanguages: utils.NewAvailableLanguages(),
	}
}
