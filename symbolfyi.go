// Package symbolfyi provides a Go client for the SymbolFYI API.
//
// SymbolFYI (symbolfyi.com) — Symbol encoding in 11 formats and Unicode property lookup.
// This client requires no authentication and has zero external dependencies.
//
// Usage:
//
//	client := symbolfyi.NewClient()
//	result, err := client.Search("example")
package symbolfyi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// DefaultBaseURL is the default base URL for the SymbolFYI API.
const DefaultBaseURL = "https://symbolfyi.com/api"

// Client is a SymbolFYI API client.
type Client struct {
	BaseURL    string
	HTTPClient *http.Client
}

// NewClient creates a new SymbolFYI API client with default settings.
func NewClient() *Client {
	return &Client{
		BaseURL:    DefaultBaseURL,
		HTTPClient: &http.Client{},
	}
}

func (c *Client) get(path string, result interface{}) error {
	resp, err := c.HTTPClient.Get(c.BaseURL + path)
	if err != nil {
		return fmt.Errorf("symbolfyi: request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("symbolfyi: HTTP %d: %s", resp.StatusCode, string(body))
	}

	if err := json.NewDecoder(resp.Body).Decode(result); err != nil {
		return fmt.Errorf("symbolfyi: decode failed: %w", err)
	}
	return nil
}

// Search searches across all content.
func (c *Client) Search(query string) (*SearchResult, error) {
	var result SearchResult
	path := fmt.Sprintf("/search/?q=%s", url.QueryEscape(query))
	if err := c.get(path, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Entity returns details for a symbol by slug.
func (c *Client) Entity(slug string) (*EntityDetail, error) {
	var result EntityDetail
	if err := c.get("/symbol/"+slug+"/", &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// GlossaryTerm returns a glossary term by slug.
func (c *Client) GlossaryTerm(slug string) (*GlossaryTerm, error) {
	var result GlossaryTerm
	if err := c.get("/glossary/"+slug+"/", &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Random returns a random symbol.
func (c *Client) Random() (*EntityDetail, error) {
	var result EntityDetail
	if err := c.get("/random/", &result); err != nil {
		return nil, err
	}
	return &result, nil
}
