package ecobee

import (
	"encoding/json"
	"net/http"
	"time"

	"golang.org/x/oauth2"
)

const (
	apiBaseURL = "https://api.ecobee.com/"
	apiVersion = 1
)

type client struct {
	apiBaseURL        string
	apiVersion        int
	httpClient        *http.Client
	customHTTPHeaders map[string]string
}

type clientOptionalParameters func(*Client)

// A Client is an ecobee API client. Its zero value is not a usable ecobee client.
type Client struct {
	client
}

// NewClient initializes a new API client with default values. It takes functors
// to modify values when creating it
func NewClient(optionalParameters ...clientOptionalParameters) *Client {
	client := &Client{
		client: client{
			apiBaseURL: apiBaseURL,
			apiVersion: apiVersion,
			httpClient: &http.Client{},
		},
	}

	for _, optionalParameter := range optionalParameters {
		optionalParameter(client)
	}

	return client
}

// WithAPIBaseURL returns a function that initializes a Client with an
// API base URL.
func WithAPIBaseURL(apiBaseURL string) func(*Client) {
	return func(c *Client) {
		c.apiBaseURL = apiBaseURL
	}
}

// WithAPIVersion returns a function that initializes a Client with an
// API version.
func WithAPIVersion(apiVersion int) func(*Client) {
	return func(c *Client) {
		c.apiVersion = apiVersion
	}
}

// WithCustomHTTPHeaders returns a function that initializes a Client with
// custom HTTP headers.
func WithCustomHTTPHeaders(customHTTPHeaders map[string]string) func(*Client) {
	return func(c *Client) {
		c.customHTTPHeaders = customHTTPHeaders
	}
}

// WithHTTPClient returns a function that initializes a Client with an HTTP
// client.
func WithHTTPClient(httpClient *http.Client) func(*Client) {
	return func(c *Client) {
		c.httpClient = httpClient
	}
}

// String returns a string representing an indented JSON encoding of the client.
func (c *Client) String() string {
	temp := struct {
		APIBaseURL           string    `json:""`
		APIVersion           int       `json:""`
		TokenType            string    `json:""`
		AccessToken          string    `json:""`
		AccessTokenExpiresOn time.Time `json:""`
		RefreshToken         string    `json:""`
	}{
		APIBaseURL: c.apiBaseURL,
		APIVersion: c.apiVersion,
	}

	if transport, ok := c.httpClient.Transport.(*oauth2.Transport); ok {
		if token, err := transport.Source.Token(); err == nil {
			temp.TokenType = token.TokenType
			temp.AccessToken = token.AccessToken
			temp.AccessTokenExpiresOn = token.Expiry.UTC()
			temp.RefreshToken = token.RefreshToken
		}
	}

	data, _ := json.MarshalIndent(&temp, "", "    ")

	return string(data)
}

// APIBaseURL returns the client's API base URL.
func (c *Client) APIBaseURL() string {
	return c.apiBaseURL
}

// APIVersion returns the client's API version.
func (c *Client) APIVersion() int {
	return c.apiVersion
}

// CustomHTTPHeaders returns the client's custom HTTP headers.
func (c *Client) CustomHTTPHeaders() map[string]string {
	customHTTPHeaders := make(map[string]string)

	for k, v := range c.customHTTPHeaders {
		customHTTPHeaders[k] = v
	}

	return customHTTPHeaders
}

// OAuth2Token returns the client's OAuth2 token.
func (c *Client) OAuth2Token() *oauth2.Token {
	transport, ok := c.httpClient.Transport.(*oauth2.Transport)
	if !ok {
		return nil
	}

	token, err := transport.Source.Token()
	if err != nil {
		return nil
	}

	tokenCopy := *token

	return &tokenCopy
}
