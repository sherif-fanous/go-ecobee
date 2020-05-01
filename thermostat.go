package ecobee

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/sfanous/go-ecobee/objects"
	"io"
	"io/ioutil"
	"net/url"
)

const (
	thermostatEndpoint        = "thermostat"
	thermostatSummaryEndpoint = "thermostatSummary"
)

// Thermostat retrieves a selection of thermostat data for one or more
// thermostats.
//
// For more information see: https://www.ecobee.com/home/developer/api/documentation/v1/operations/get-thermostats.shtml
func (c *Client) Thermostat(ctx context.Context, selection *objects.Selection, page *objects.Page) (*ThermostatSuccessResponse, error) {
	data, err := json.Marshal(struct {
		Selection *objects.Selection `json:"selection,omitempty"`
		Page      *objects.Page      `json:"page,omitempty"`
	}{
		Selection: selection,
		Page:      page,
	})
	if err != nil {
		return nil, fmt.Errorf("%s: %w", thermostatEndpoint, err)
	}

	queryParameters := url.Values{}
	queryParameters.Set("json", string(data))

	resp, err := c.get(ctx, fmt.Sprintf("%s%d/%s", c.apiBaseURL, c.apiVersion, thermostatEndpoint), queryParameters, nil)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", thermostatEndpoint, err)
	}

	defer func() {
		if resp != nil {
			_, _ = io.Copy(ioutil.Discard, resp.Body)
			_ = resp.Body.Close()
		}
	}()

	thermostatSuccessResponse := ThermostatSuccessResponse{}

	if err := processAPIResponse(thermostatEndpoint, resp, &thermostatSuccessResponse); err != nil {
		return nil, err
	}

	return &thermostatSuccessResponse, nil
}

// ThermostatSummary retrieves a list of thermostat configuration and state
// revisions.
//
// For more information see: https://www.ecobee.com/home/developer/api/documentation/v1/operations/get-thermostat-summary.shtml
func (c *Client) ThermostatSummary(ctx context.Context, selection *objects.Selection) (*ThermostatSummarySuccessResponse, error) {
	data, err := json.Marshal(struct {
		Selection *objects.Selection `json:"selection,omitempty"`
	}{
		Selection: selection,
	})
	if err != nil {
		return nil, fmt.Errorf("%s: %w", thermostatSummaryEndpoint, err)
	}

	queryParameters := url.Values{}
	queryParameters.Set("json", string(data))

	resp, err := c.get(ctx, fmt.Sprintf("%s%d/%s", c.apiBaseURL, c.apiVersion, thermostatSummaryEndpoint), queryParameters, nil)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", thermostatSummaryEndpoint, err)
	}

	defer func() {
		if resp != nil {
			_, _ = io.Copy(ioutil.Discard, resp.Body)
			_ = resp.Body.Close()
		}
	}()

	thermostatSummaryResponse := ThermostatSummarySuccessResponse{}

	if err := processAPIResponse(thermostatSummaryEndpoint, resp, &thermostatSummaryResponse); err != nil {
		return nil, err
	}

	return &thermostatSummaryResponse, nil
}

// UpdateThermostat permits the modification of any writable thermostat or
// sub-object property.
//
// For more information see: https://www.ecobee.com/home/developer/api/documentation/v1/operations/post-update-thermostats.shtml
func (c *Client) UpdateThermostat(ctx context.Context, selection *objects.Selection, thermostat *objects.Thermostat, functions []objects.Function) (*APIStatusResponse, error) {
	data, err := json.Marshal(struct {
		Selection  *objects.Selection  `json:"selection,omitempty"`
		Thermostat *objects.Thermostat `json:"thermostat,omitempty"`
		Functions  []objects.Function  `json:"functions,omitempty"`
	}{
		Selection:  selection,
		Thermostat: thermostat,
		Functions:  functions,
	})
	if err != nil {
		return nil, fmt.Errorf("%s: %w", thermostatEndpoint, err)
	}

	queryParameters := url.Values{}
	queryParameters.Set("format", "json")

	resp, err := c.post(ctx, fmt.Sprintf("%s%d/%s", c.apiBaseURL, c.apiVersion, thermostatEndpoint), queryParameters, nil, bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("%s: %w", thermostatEndpoint, err)
	}

	defer func() {
		if resp != nil {
			_, _ = io.Copy(ioutil.Discard, resp.Body)
			_ = resp.Body.Close()
		}
	}()

	updateThermostatResponse := APIStatusResponse{}

	if err := processAPIResponse(thermostatEndpoint, resp, &updateThermostatResponse); err != nil {
		return nil, err
	}

	return &updateThermostatResponse, nil
}
