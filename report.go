package ecobee

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/url"
	"time"

	"github.com/sfanous/go-ecobee/objects"
)

const (
	meterReportEndpoint   = "meterReport"
	runtimeReportEndpoint = "runtimeReport"
)

// A MeterType specifies a meter type.
type MeterType string

// Supported MeterType values.
const (
	MeterTypeEnergy MeterType = "energy"
)

// An MeterReportParameters specifies the request parameters of the
// MeterReport method.
type MeterReportParameters struct {
	// The report UTC start date.
	StartDate *time.Time
	// The report start interval.
	StartInterval *int
	// The report UTC end date.
	EndDate *time.Time
	// The report end interval.
	EndInterval *int
	// A CSV string of meter types.
	Meters []MeterType
}

// MeterReport retrieves the historical meter reading information for a
// selection of thermostats.
//
// For more information see: https://www.ecobee.com/home/developer/api/documentation/v1/operations/get-meter-report.shtml
func (c *Client) MeterReport(ctx context.Context, selection *objects.Selection, parameters *MeterReportParameters) (*MeterReportSuccessResponse, error) {
	data, err := json.Marshal(struct {
		Selection     *objects.Selection `json:"selection,omitempty"`
		StartDate     *string            `json:"startDate,omitempty"`
		StartInterval *int               `json:"startInterval,omitempty"`
		EndDate       *string            `json:"endDate,omitempty"`
		EndInterval   *int               `json:"endInterval,omitempty"`
		Meters        *string            `json:"meters,omitempty"`
	}{
		Selection:     selection,
		StartDate:     String((*parameters.StartDate).Format("2006-01-02")),
		StartInterval: parameters.StartInterval,
		EndDate:       String((*parameters.EndDate).Format("2006-01-02")),
		EndInterval:   parameters.EndInterval,
		Meters:        String(string(parameters.Meters[0])),
	})
	if err != nil {
		return nil, fmt.Errorf("%s: %w", meterReportEndpoint, err)
	}

	queryParameters := url.Values{}
	queryParameters.Set("format", "json")
	queryParameters.Set("body", string(data))

	resp, err := c.get(ctx, fmt.Sprintf("%s%d/%s", apiBaseURL, apiVersion, meterReportEndpoint), queryParameters, map[string][]string{"Authorization": {fmt.Sprintf("%s %s", c.tokenType, c.accessToken)}})
	if err != nil {
		return nil, fmt.Errorf("%s: %w", meterReportEndpoint, err)
	}

	defer func() {
		if resp != nil {
			_, _ = io.Copy(ioutil.Discard, resp.Body)
			_ = resp.Body.Close()
		}
	}()

	meterReportResponse := MeterReportSuccessResponse{}

	if err := processAPIResponse(meterReportEndpoint, resp, &meterReportResponse); err != nil {
		return nil, err
	}

	return &meterReportResponse, nil
}

// An RuntimeReportParameters specifies the request parameters of the
// RuntimeReport method.
type RuntimeReportParameters struct {
	// The report UTC start date.
	StartDate *time.Time
	// The report start interval.
	StartInterval *int
	// The report UTC end date.
	EndDate *time.Time
	// The report end interval.
	EndInterval *int
	// A CSV string of column names.
	Columns *string
	// Whether to include sensor runtime report data
	IncludeSensor *bool
}

// RuntimeReport retrieves the historical runtime report information for a
// selection of thermostats.
//
// For more information see: https://www.ecobee.com/home/developer/api/documentation/v1/operations/get-runtime-report.shtml
func (c *Client) RuntimeReport(ctx context.Context, selection *objects.Selection, parameters *RuntimeReportParameters) (*RuntimeReportSuccessResponse, error) {
	data, err := json.Marshal(struct {
		Selection     *objects.Selection `json:"selection,omitempty"`
		StartDate     *string            `json:"startDate,omitempty"`
		StartInterval *int               `json:"startInterval,omitempty"`
		EndDate       *string            `json:"endDate,omitempty"`
		EndInterval   *int               `json:"endInterval,omitempty"`
		Columns       *string            `json:"columns,omitempty"`
		IncludeSensor *bool              `json:"includeSensor,omitempty"`
	}{
		Selection:     selection,
		StartDate:     String((*parameters.StartDate).Format("2006-01-02")),
		StartInterval: parameters.StartInterval,
		EndDate:       String((*parameters.EndDate).Format("2006-01-02")),
		EndInterval:   parameters.EndInterval,
		Columns:       parameters.Columns,
		IncludeSensor: parameters.IncludeSensor,
	})
	if err != nil {
		return nil, fmt.Errorf("%s: %w", runtimeReportEndpoint, err)
	}

	queryParameters := url.Values{}
	queryParameters.Set("format", "json")
	queryParameters.Set("body", string(data))

	resp, err := c.get(ctx, fmt.Sprintf("%s%d/%s", apiBaseURL, apiVersion, runtimeReportEndpoint), queryParameters, map[string][]string{"Authorization": {fmt.Sprintf("%s %s", c.tokenType, c.accessToken)}})
	if err != nil {
		return nil, fmt.Errorf("%s: %w", runtimeReportEndpoint, err)
	}

	defer func() {
		if resp != nil {
			_, _ = io.Copy(ioutil.Discard, resp.Body)
			_ = resp.Body.Close()
		}
	}()

	runtimeReportResponse := RuntimeReportSuccessResponse{}

	if err := processAPIResponse(runtimeReportEndpoint, resp, &runtimeReportResponse); err != nil {
		return nil, err
	}

	return &runtimeReportResponse, nil
}
