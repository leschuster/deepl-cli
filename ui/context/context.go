package context

import (
	deeplapi "github.com/leschuster/deepl-cli/pkg/deepl-api"
	"github.com/leschuster/deepl-cli/ui/keys"
	"github.com/leschuster/deepl-cli/ui/styles"
	"github.com/leschuster/deepl-cli/ui/utils"
)

type ProgramContext struct {
	Api                *deeplapi.DeeplAPI
	Keys               keys.KeyMap
	ScreenWidth        int
	ScreenHeight       int
	Styles             *styles.Styles
	SourceLanguage     *deeplapi.Language
	TargetLanguage     *deeplapi.Language
	Formality          string
	AvailableLanguages utils.AvailableLanguages
}

func New(api *deeplapi.DeeplAPI) *ProgramContext {
	return &ProgramContext{
		Api:                api,
		Keys:               keys.DefaultKeyMap(),
		Styles:             styles.New(),
		AvailableLanguages: utils.NewAvailableLanguages(api),
	}
}
