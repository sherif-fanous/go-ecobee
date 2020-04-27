package ecobee

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"net/http"
	"time"
)

const (
	apiBaseURL = "https://api.ecobee.com/"
	apiVersion = 1
)

type client struct {
	apiBaseURL            string
	apiVersion            int
	applicationKey        string
	authorizationToken    string
	tokenType             string
	accessToken           string
	accessTokenExpiresOn  time.Time
	refreshToken          string
	refreshTokenExpiresOn time.Time
	httpClient            *http.Client
	customHTTPHeaders     map[string]string
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

// WithApplicationKey returns a function that initializes a Client with an
// application key.
func WithApplicationKey(applicationKey string) func(*Client) {
	return func(c *Client) {
		c.applicationKey = applicationKey
	}
}

// WithAuthorizationToken returns a function that initializes a Client with an
// authorization token.
func WithAuthorizationToken(authorizationToken string) func(*Client) {
	return func(c *Client) {
		c.authorizationToken = authorizationToken
	}
}

// WithTokenType returns a function that initializes a Client with a token
// type.
func WithTokenType(tokenType string) func(*Client) {
	return func(c *Client) {
		c.tokenType = tokenType
	}
}

// WithAccessToken returns a function that initializes a Client with an access
// token.
func WithAccessToken(accessToken string) func(*Client) {
	return func(c *Client) {
		c.accessToken = accessToken
	}
}

// WithAccessTokenExpiresOn returns a function that initializes a Client with
// a time.Time specifying when the access token expires.
func WithAccessTokenExpiresOn(accessTokenExpiresOn time.Time) func(*Client) {
	return func(c *Client) {
		c.accessTokenExpiresOn = accessTokenExpiresOn
	}
}

// WithRefreshToken returns a function that initializes a Client with an
// refresh token.
func WithRefreshToken(refreshToken string) func(*Client) {
	return func(c *Client) {
		c.refreshToken = refreshToken
	}
}

// WithRefreshTokenExpiresOn returns a function that initializes a Client with
// a time.Time specifying when the refresh token expires.
func WithRefreshTokenExpiresOn(refreshTokenExpiresOn time.Time) func(*Client) {
	return func(c *Client) {
		c.refreshTokenExpiresOn = refreshTokenExpiresOn
	}
}

// WithHTTPClient returns a function that initializes a Client with an HTTP
// client.
func WithHTTPClient(httpClient *http.Client) func(*Client) {
	return func(c *Client) {
		c.httpClient = httpClient
	}
}

// WithCustomHTTPHeaders returns a function that initializes a Client with
// custom HTTP headers.
func WithCustomHTTPHeaders(customHTTPHeaders map[string]string) func(*Client) {
	return func(c *Client) {
		c.customHTTPHeaders = customHTTPHeaders
	}
}

// GobDecode implements the gob.GobDecoder interface.
func (c *Client) GobDecode(b []byte) error {
	temp := struct {
		APIBaseURL            string
		APIVersion            int
		ApplicationKey        string
		AuthorizationToken    string
		TokenType             string
		AccessToken           string
		AccessTokenExpiresOn  time.Time
		RefreshToken          string
		RefreshTokenExpiresOn time.Time
		CustomHTTPHeaders     map[string]string
	}{}

	dec := gob.NewDecoder(bytes.NewBuffer(b))

	_ = dec.Decode(&temp)

	c.apiBaseURL = temp.APIBaseURL
	c.apiVersion = temp.APIVersion
	c.applicationKey = temp.ApplicationKey
	c.authorizationToken = temp.AuthorizationToken
	c.tokenType = temp.TokenType
	c.accessToken = temp.AccessToken
	c.accessTokenExpiresOn = temp.AccessTokenExpiresOn
	c.refreshToken = temp.RefreshToken
	c.refreshTokenExpiresOn = temp.RefreshTokenExpiresOn
	c.customHTTPHeaders = temp.CustomHTTPHeaders

	return nil
}

// GobEncode implements the gob.GobEncoder interface.
func (c *Client) GobEncode() ([]byte, error) {
	temp := struct {
		APIBaseURL            string
		APIVersion            int
		ApplicationKey        string
		AuthorizationToken    string
		TokenType             string
		AccessToken           string
		AccessTokenExpiresOn  time.Time
		RefreshToken          string
		RefreshTokenExpiresOn time.Time
		CustomHTTPHeaders     map[string]string
	}{
		APIBaseURL:            c.apiBaseURL,
		APIVersion:            c.apiVersion,
		ApplicationKey:        c.applicationKey,
		AuthorizationToken:    c.authorizationToken,
		TokenType:             c.tokenType,
		AccessToken:           c.accessToken,
		AccessTokenExpiresOn:  c.accessTokenExpiresOn,
		RefreshToken:          c.refreshToken,
		RefreshTokenExpiresOn: c.refreshTokenExpiresOn,
		CustomHTTPHeaders:     c.customHTTPHeaders,
	}

	b := bytes.Buffer{}
	enc := gob.NewEncoder(&b)

	_ = enc.Encode(&temp)

	return b.Bytes(), nil
}

// String returns a string representing an indented JSON encoding of the client.
func (c *Client) String() string {
	temp := struct {
		APIBaseURL            string    `json:""`
		APIVersion            int       `json:""`
		ApplicationKey        string    `json:""`
		AuthorizationToken    string    `json:""`
		TokenType             string    `json:""`
		AccessToken           string    `json:""`
		AccessTokenExpiresOn  time.Time `json:""`
		RefreshToken          string    `json:""`
		RefreshTokenExpiresOn time.Time `json:""`
	}{
		APIBaseURL:            c.apiBaseURL,
		APIVersion:            c.apiVersion,
		ApplicationKey:        c.applicationKey,
		AuthorizationToken:    c.authorizationToken,
		TokenType:             c.tokenType,
		AccessToken:           c.accessToken,
		AccessTokenExpiresOn:  c.accessTokenExpiresOn,
		RefreshToken:          c.refreshToken,
		RefreshTokenExpiresOn: c.refreshTokenExpiresOn,
	}

	data, _ := json.MarshalIndent(&temp, "", "    ")

	return string(data)
}

// APIBaseURL returns the client's API base URL.
func (c *Client) APIBaseURL() string {
	return c.apiBaseURL
}

// SetAPIBaseURL sets the client's API base URL.
func (c *Client) SetAPIBaseURL(apiBaseURL string) {
	c.apiBaseURL = apiBaseURL
}

// APIVersion returns the client's API version.
func (c *Client) APIVersion() int {
	return c.apiVersion
}

// SetAPIVersion sets the client's API version.
func (c *Client) SetAPIVersion(apiVersion int) {
	c.apiVersion = apiVersion
}

// ApplicationKey returns the client's application key.
func (c *Client) ApplicationKey() string {
	return c.applicationKey
}

// AuthorizationToken returns the client's authorization token.
func (c *Client) AuthorizationToken() string {
	return c.authorizationToken
}

// TokenType returns the client's token type.
func (c *Client) TokenType() string {
	return c.tokenType
}

// AccessToken returns the client's access token.
func (c *Client) AccessToken() string {
	return c.accessToken
}

// AccessTokenExpiresOn returns when the client's access token expires.
func (c *Client) AccessTokenExpiresOn() time.Time {
	return c.accessTokenExpiresOn
}

// RefreshToken returns the client's refresh token.
func (c *Client) RefreshToken() string {
	return c.refreshToken
}

// RefreshTokenExpiresOn returns when the client's refresh token expires.
func (c *Client) RefreshTokenExpiresOn() time.Time {
	return c.refreshTokenExpiresOn
}

// CustomHTTPHeaders returns the client's custom HTTP headers.
func (c *Client) CustomHTTPHeaders() map[string]string {
	customHTTPHeaders := make(map[string]string)

	for k, v := range c.customHTTPHeaders {
		customHTTPHeaders[k] = v
	}

	return customHTTPHeaders
}

// SetHTTPClient sets the client's HTTP client.
func (c *Client) SetHTTPClient(httpClient *http.Client) {
	c.httpClient = httpClient
}

// SetCustomHTTPHeaders sets the client's custom HTTP headers.
func (c *Client) SetCustomHTTPHeaders(customHTTPHeaders map[string]string) {
	c.customHTTPHeaders = customHTTPHeaders
}
