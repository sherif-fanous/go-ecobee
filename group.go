package ecobee

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/sherif-fanous/go-ecobee/objects"
	"io"
	"io/ioutil"
	"net/url"
)

const (
	groupEndpoint = "group"
)

// Group retrieves the grouping data for the thermostats registered to
// the particular user.
//
// For more information see: https://www.ecobee.com/home/developer/api/documentation/v1/operations/get-group.shtml
func (c *Client) Group(ctx context.Context, selection *objects.Selection) (*GroupSuccessResponse, error) {
	data, err := json.Marshal(struct {
		Selection *objects.Selection `json:"selection,omitempty"`
	}{
		Selection: selection,
	})
	if err != nil {
		return nil, fmt.Errorf("%s: %w", groupEndpoint, err)
	}

	queryParameters := url.Values{}
	queryParameters.Set("format", "json")
	queryParameters.Set("body", string(data))

	resp, err := c.get(ctx, fmt.Sprintf("%s%d/%s", c.apiBaseURL, c.apiVersion, groupEndpoint), queryParameters, nil)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", groupEndpoint, err)
	}

	defer func() {
		if resp != nil {
			_, _ = io.Copy(ioutil.Discard, resp.Body)
			_ = resp.Body.Close()
		}
	}()

	groupResponse := GroupSuccessResponse{}

	if err := processAPIResponse(groupEndpoint, resp, &groupResponse); err != nil {
		return nil, err
	}

	return &groupResponse, nil
}

// UpdateGroup updates the grouping data for the thermostats registered to
// the particular user.
//
// For more information see: https://www.ecobee.com/home/developer/api/documentation/v1/operations/post-group-update.shtml
func (c *Client) UpdateGroup(ctx context.Context, selection *objects.Selection, groups []objects.Group) (*GroupSuccessResponse, error) {
	data, err := json.Marshal(struct {
		Selection *objects.Selection `json:"selection,omitempty"`
		Groups    []objects.Group    `json:"groups,omitempty"`
	}{
		Selection: selection,
		Groups:    groups,
	})
	if err != nil {
		return nil, fmt.Errorf("%s: %w", groupEndpoint, err)
	}

	queryParameters := url.Values{}
	queryParameters.Set("format", "json")

	resp, err := c.post(ctx, fmt.Sprintf("%s%d/%s", c.apiBaseURL, c.apiVersion, groupEndpoint), queryParameters, nil, bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("%s: %w", groupEndpoint, err)
	}

	defer func() {
		if resp != nil {
			_, _ = io.Copy(ioutil.Discard, resp.Body)
			_ = resp.Body.Close()
		}
	}()

	updateGroupResponse := GroupSuccessResponse{}

	if err := processAPIResponse(groupEndpoint, resp, &updateGroupResponse); err != nil {
		return nil, err
	}

	return &updateGroupResponse, nil
}
