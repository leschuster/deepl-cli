// Package deeplapi provides a wrapper around the official DeepL API.
// You can translate text and retrieve available languages.
// An API key is required. Both free and paid tiers are supported.
package deeplapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

// Defines the formality of the translated text.
// Not supported by all languages.
const (
	FormalityDefault    = "default"
	FormalityMore       = "more"
	FormalityLess       = "less"
	FormalityPreferMore = "prefer_more"
	FormalityPreferLess = "prefer_less"
)

const baseURLFree = "https://api-free.deepl.com/v2"
const baseURLPro = "https://api.deepl.com/v2"

// DeeplAPI provides abstract access to the official DeepL API
type DeeplAPI struct {
	apiKey  string
	baseURL string
	client  *http.Client
}

// Creates a new DeeplAPI instance
func New(apiKey string) *DeeplAPI {
	// Free API keys have ":fx" as suffix
	isFreeTier := strings.HasSuffix(apiKey, ":fx")

	baseURL := baseURLPro
	if isFreeTier {
		baseURL = baseURLFree
	}

	return &DeeplAPI{
		apiKey:  apiKey,
		baseURL: baseURL,
		client:  &http.Client{},
	}
}

// Parameters for DeeplAPI.Translate
// Text and TargetLang are required
type TranslateParams struct {
	Text       []string `json:"text"`        // Text to translate, UTF-8
	SourceLang string   `json:"source_lang"` // Original language code, optional
	TargetLang string   `json:"target_lang"` // Target Language code
	Context    string   `json:"context"`     // Additional context that influences the translation, but is not translated itself, optional
	Formality  string   `json:"formality"`   // Define whether the text should be formal or more informal, not supported by all languages, optional
}

// Response type for DeeplAPI.Translate
type TranslateResp struct {
	Translations []struct {
		DetectedSourceLanguage string `json:"detected_source_language"`
		Text                   string `json:"text"`
	} `json:"translations"`
}

// The Translate function uses DeepL to translate params.Text into the specified language
func (api *DeeplAPI) Translate(params TranslateParams) (*TranslateResp, error) {

	// Marshal request body
	body, err := json.Marshal(params)
	if err != nil {
		return nil, fmt.Errorf("could not marshal options to JSON: %v", err)
	}

	// Make request
	data, err := api.request("/translate", http.MethodPost, body)
	if err != nil {
		return nil, err
	}

	// Unmarshal response
	responseObj := TranslateResp{}
	err = json.Unmarshal(data, &responseObj)
	if err != nil {
		return nil, fmt.Errorf("could not unmarshall response: %v", err)
	}

	return &responseObj, nil
}

// Language represents a language that is supported by DeepL
// Beware of the fact that DeepL does support different languages as source
// and target languages.
type Language struct {
	Language          string `json:"language"`           // Language code
	Name              string `json:"name"`               // Friendly name
	SupportsFormality bool   `json:"supports_formality"` // Whether DeepL supports specifying a formality or not (when used as target langauge)
}

// Response type for DeeplAPI.GetLanguages
type GetLanguagesResp struct {
	Source []Language
	Target []Language
}

// The GetLanguages function retrieves all languages that DeepL supports
// Because supported source and target languages may differ, the reponse differentiates
// between them
func (api *DeeplAPI) GetLanguages() (*GetLanguagesResp, error) {
	// Fetch source languages
	srcLangRaw, err := api.request("/languages?type=source", http.MethodGet, nil)
	if err != nil {
		return nil, err
	}

	srcLang := []Language{}
	err = json.Unmarshal(srcLangRaw, &srcLang)
	if err != nil {
		return nil, fmt.Errorf("could not unmarshall response: %v", err)
	}

	// Fetch target languages
	tarLangRaw, err := api.request("/languages?type=target", http.MethodGet, nil)
	if err != nil {
		return nil, err
	}

	tarLang := []Language{}
	err = json.Unmarshal(tarLangRaw, &tarLang)
	if err != nil {
		return nil, fmt.Errorf("could not unmarshall response: %v", err)
	}

	return &GetLanguagesResp{
		Source: srcLang,
		Target: tarLang,
	}, nil
}

// Helper function to perform a request to the DeepL API
func (api *DeeplAPI) request(endpoint, method string, body []byte) ([]byte, error) {
	// Join path
	reqURL := api.baseURL + endpoint

	// Create request
	req, err := http.NewRequest(method, reqURL, bytes.NewBuffer(body))
	if err != nil {
		return nil, fmt.Errorf("could not create request: %v", err)
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("DeepL-Auth-Key %s", api.apiKey))

	// Perform request
	resp, err := api.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request to '%s' failed", reqURL)
	}
	defer resp.Body.Close()
	if resp.StatusCode > 299 {
		return nil, fmt.Errorf("request to '%s' failed with status %s", reqURL, resp.Status)
	}

	// Read response body
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("request to '%s' failed: could not read response body", reqURL)
	}

	return respBody, nil
}
