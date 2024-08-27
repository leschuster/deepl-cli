package deeplapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

const (
	FormalityDefault    = "default"
	FormalityMore       = "more"
	FormalityLess       = "less"
	FormalityPreferMore = "prefer_more"
	FormalityPreferLess = "prefer_less"
)

const baseURLFree = "https://api-free.deepl.com/v2"
const baseURLPro = "https://api.deepl.com/v2"

type DeeplAPI struct {
	apiKey  string
	baseURL string
	client  *http.Client
}

func New(apiKey string) *DeeplAPI {
	// API keys having ":fx" at the end are free tier
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

type TranslateParams struct {
	Text       []string `json:"text"`
	SourceLang string   `json:"source_lang"`
	TargetLang string   `json:"target_lang"`
	Context    string   `json:"context"`
	Formality  string   `json:"formality"`
}

type TranslateResp struct {
	Translations []struct {
		DetectedSourceLanguage string `json:"detected_source_language"`
		Text                   string `json:"text"`
	} `json:"translations"`
}

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

type Language struct {
	Language          string `json:"language"`
	Name              string `json:"name"`
	SupportsFormality string `json:"supports_formality"`
}

type GetLanguagesResp struct {
	Source []Language
	Target []Language
}

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

	// Do request
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
