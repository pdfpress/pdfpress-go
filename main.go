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

// GenerateParams are params for generating via html or url
type GenerateParams struct {
	URL  string `json:"url,omitempty"`
	HTML string `json:"html,omitempty"`
}

// Generate generates a pdf from url or html
func Generate(ctx context.Context, body *GenerateParams) (io.ReadCloser, error) {
	b, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	resp, err := http.Post(fmt.Sprintf("%s/generate", apiBase), "application/json", bytes.NewBuffer(b))
	if err != nil {
		return nil, err
	}
	return resp.Body, nil
}

// GenerateFromTemplate generates a PDF using the template and provided params.
func GenerateFromTemplate(templateID string, params interface{}) (io.ReadCloser, error) {
	return nil, nil
}
