package deeplapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
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
	data, err := api.post("/translate", body)
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

func (api *DeeplAPI) post(endpoint string, body []byte) ([]byte, error) {
	// Join path
	reqURL, err := url.JoinPath(api.baseURL, endpoint)
	if err != nil {
		return nil, fmt.Errorf("could not join paths: %v", err)
	}

	// Create request
	req, err := http.NewRequest(http.MethodGet, reqURL, bytes.NewBuffer(body))
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
		return nil, fmt.Errorf("request to '%s' failed with status %s: %v", reqURL, resp.Status, err)
	}

	// Read response body
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("request to '%s' failed: could not read response body", reqURL)
	}

	return respBody, nil
}
