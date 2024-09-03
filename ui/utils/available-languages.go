package utils

import (
	"fmt"
	"sync"

	tea "github.com/charmbracelet/bubbletea"
	deeplapi "github.com/leschuster/deepl-cli/pkg/deepl-api"
	"github.com/leschuster/deepl-cli/ui/com"
)

type AvailableLanguages struct {
	api      *deeplapi.DeeplAPI
	srcLangs []deeplapi.Language
	tarLangs []deeplapi.Language
	mu       *sync.RWMutex
}

func NewAvailableLanguages(api *deeplapi.DeeplAPI) AvailableLanguages {
	return AvailableLanguages{
		api: api,
		mu:  &sync.RWMutex{},
	}
}

func (al *AvailableLanguages) LoadInitial() tea.Msg {
	al.mu.Lock()
	defer al.mu.Unlock()

	if al.srcLangs != nil && al.tarLangs != nil {
		// Already fetched data

		// We need to execute the cmd again so that newly created components
		// will fetch the data
		return com.APILanguagesReceivedMsg{}
	}

	resp, err := al.api.GetLanguages()
	if err != nil {
		return com.ThrowErr(err)
	}

	al.srcLangs = resp.Source
	al.tarLangs = resp.Target

	return com.APILanguagesReceivedMsg{}
}

func (al *AvailableLanguages) GetSourceLanguages() ([]deeplapi.Language, error) {
	al.mu.RLock()
	defer al.mu.RUnlock()

	if al.srcLangs == nil {
		return nil, fmt.Errorf("could not get source languages: srcLangs is nil")
	}

	return al.srcLangs, nil
}

func (al *AvailableLanguages) GetTargetLanguages() ([]deeplapi.Language, error) {
	al.mu.RLock()
	defer al.mu.RUnlock()

	if al.tarLangs == nil {
		return nil, fmt.Errorf("could not get target languages: tarLangs is nil")
	}

	return al.tarLangs, nil
}
