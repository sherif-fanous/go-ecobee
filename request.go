package ecobee

import (
	"context"
	"io"
	"net/http"
	"net/url"

	"golang.org/x/net/context/ctxhttp"
)

func (c *Client) get(ctx context.Context, endpoint string, queryParameters url.Values, headers map[string][]string) (*http.Response, error) {
	return c.sendRequest(ctx, http.MethodGet, endpoint, queryParameters, headers, nil)
}

func (c *Client) post(ctx context.Context, endpoint string, queryParameters url.Values, headers map[string][]string, body io.Reader) (*http.Response, error) {
	return c.sendRequest(ctx, http.MethodPost, endpoint, queryParameters, headers, body)
}

func (c *Client) sendRequest(ctx context.Context, method string, endpoint string, queryParameters url.Values, headers map[string][]string, body io.Reader) (*http.Response, error) {
	req, err := http.NewRequest(method, endpoint, body)
	if err != nil {
		return nil, err
	}

	req.URL.RawQuery = queryParameters.Encode()

	for k, v := range c.customHTTPHeaders {
		req.Header.Set(k, v)
	}

	for k, v := range headers {
		req.Header[k] = v
	}

	resp, err := c.doRequest(ctx, req)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

func (c *Client) doRequest(ctx context.Context, req *http.Request) (*http.Response, error) {
	resp, err := ctxhttp.Do(ctx, c.httpClient, req)
	if err != nil {
		return resp, err
	}

	return resp, nil
}
