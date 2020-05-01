package ecobee

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"net/url"
)

const (
	authorizeEndpoint = "authorize"
	tokenEndpoint     = "token"
)

// Scope details the intent of the API application towards the user's account.
type Scope string

// Supported Scope values.
const (
	ScopeSmartRead  Scope = "smartRead"
	ScopeSmartWrite Scope = "smartWrite"
	ScopeEMS        Scope = "ems"
)

// PINAuthorization requests an authorization code and a PIN to enable the
// authorization of an application within the ecobee Web Portal.
//
// For more information see: https://www.ecobee.com/home/developer/api/documentation/v1/auth/pin-api-authorization.shtml
func (c *Client) PINAuthorization(ctx context.Context, applicationKey string, scope Scope) (*PINAuthorizationSuccessResponse, error) {
	queryParameters := url.Values{}
	queryParameters.Set("response_type", "ecobeePin")
	queryParameters.Set("client_id", applicationKey)
	queryParameters.Set("scope", string(scope))

	resp, err := c.get(ctx, fmt.Sprintf("%s%s", c.apiBaseURL, authorizeEndpoint), queryParameters, nil)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", authorizeEndpoint, err)
	}

	defer func() {
		if resp != nil {
			_, _ = io.Copy(ioutil.Discard, resp.Body)
			_ = resp.Body.Close()
		}
	}()

	authorizeResponse := PINAuthorizationSuccessResponse{}

	if err := processAuthorizationResponse(tokenEndpoint, resp, &authorizeResponse); err != nil {
		return &authorizeResponse, err
	}

	return &authorizeResponse, nil
}

// RequestTokens requests access and refresh tokens once the application
// is authorized within the ecobee Web Portal.
//
// For more information see: https://www.ecobee.com/home/developer/api/documentation/v1/auth/pin-api-authorization.shtml
func (c *Client) RequestTokens(ctx context.Context, applicationKey string, authorizationToken string) (*TokensSuccessResponse, error) {
	queryParameters := url.Values{}
	queryParameters.Set("grant_type", "ecobeePin")
	queryParameters.Set("code", authorizationToken)
	queryParameters.Set("client_id", applicationKey)

	resp, err := c.post(ctx, fmt.Sprintf("%s%s", c.apiBaseURL, tokenEndpoint), queryParameters, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", tokenEndpoint, err)
	}

	defer func() {
		if resp != nil {
			_, _ = io.Copy(ioutil.Discard, resp.Body)
			_ = resp.Body.Close()
		}
	}()

	tokensResponse := TokensSuccessResponse{}

	if err := processAuthorizationResponse(tokenEndpoint, resp, &tokensResponse); err != nil {
		return &tokensResponse, nil
	}

	return &tokensResponse, err
}
