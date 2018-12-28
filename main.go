package pdfpress

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const apiBase = "https://api.pdfpress.app"

// Key is account's api key
var Key string

// Client is the http client used to make api requests
var Client http.Client = http.Client{}

// Margins are the page margins
type Margins struct {
	Top    float64 `json:"top"`
	Bottom float64 `json:"bottom"`
	Left   float64 `json:"left"`
	Right  float64 `json:"right"`
}

// GenerateParams are the standard pdf generation options
type GenerateParams struct {
	Engine   string            `json:"engine"`
	PageSize string            `json:"page_size"`
	Margins  Margins           `json:"margins"`
	Headers  map[string]string `json:"headers"`
	HTML     string            `json:"html"`
	URL      string            `json:"url"`
}

// Generate generates a pdf from url or html
func Generate(ctx context.Context, body *GenerateParams) (io.ReadCloser, error) {
	b, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	resp, err := call(ctx, fmt.Sprintf("%s/generate", apiBase), "application/json", bytes.NewBuffer(b))
	return resp.Body, err
}

func call(ctx context.Context, method string, url string, body io.Reader) (*http.Response, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	req.SetBasicAuth(Key, "")
	req.Header.Set("Content-Type", "application/json")
	resp, err := Client.Do(req)
	return resp, err
}

// GenerateFromTemplate generates a PDF using the template and provided params.
func GenerateFromTemplate(ctx context.Context, templateID string, params interface{}) (io.ReadCloser, error) {
	b, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}
	resp, err := call(ctx, "POST", fmt.Sprintf("%s/templates/%s/generate", apiBase, templateID), bytes.NewBuffer(b))
	return resp.Body, err
}
