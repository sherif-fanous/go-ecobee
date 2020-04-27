package ecobee

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/sfanous/go-ecobee/objects"
	"io"
	"net/http"
)

type apiStatusResponse struct {
	status *objects.Status
}

// APIStatusResponse describes responses returned by the ecobee server that
// only contain the Status object
type APIStatusResponse struct {
	apiStatusResponse
}

// String implements the fmt.Stringer interface. It returns a string
// representing an indented JSON encoding of the response.
func (a *APIStatusResponse) String() string {
	temp := struct {
		Status *objects.Status `json:""`
	}{
		Status: a.status,
	}

	data, _ := json.MarshalIndent(&temp, "", "    ")

	return string(data)
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (a *APIStatusResponse) UnmarshalJSON(data []byte) error {
	temp := struct {
		Status *objects.Status `json:"status"`
	}{}

	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}

	a.status = temp.Status

	return nil
}

// Status returns the response's status.
func (a *APIStatusResponse) Status() *objects.Status {
	return a.status
}

type authorizationErrorResponse struct {
	errorType        string
	errorDescription string
	errorURI         string
}

// AuthorizationErrorResponse describes the error response returned by the
// ecobee server while authorizing.
type AuthorizationErrorResponse struct {
	authorizationErrorResponse
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (e *AuthorizationErrorResponse) UnmarshalJSON(data []byte) error {
	temp := struct {
		ErrorType        string `json:"error"`
		ErrorDescription string `json:"error_description"`
		ErrorURI         string `json:"error_uri"`
	}{}

	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}

	e.errorType = temp.ErrorType
	e.errorDescription = temp.ErrorDescription
	e.errorURI = temp.ErrorURI

	return nil
}

// ErrorType returns the response's error type.
func (e *AuthorizationErrorResponse) ErrorType() string {
	return e.errorType
}

// ErrorDescription returns the response's error description.
func (e *AuthorizationErrorResponse) ErrorDescription() string {
	return e.errorDescription
}

// ErrorURI returns the response's error URI.
func (e *AuthorizationErrorResponse) ErrorURI() string {
	return e.errorURI
}

type groupSuccessResponse struct {
	groups []objects.Group
	status *objects.Status
}

// GroupSuccessResponse describes the response returned by the ecobee
// server while retrieving or updating groups.
type GroupSuccessResponse struct {
	groupSuccessResponse
}

// String implements the fmt.Stringer interface. It returns a string
// representing an indented JSON encoding of the response.
func (g *GroupSuccessResponse) String() string {
	temp := struct {
		Groups []objects.Group `json:""`
		Status *objects.Status `json:""`
	}{
		Groups: g.groups,
		Status: g.status,
	}

	data, _ := json.MarshalIndent(&temp, "", "    ")

	return string(data)
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (g *GroupSuccessResponse) UnmarshalJSON(data []byte) error {
	temp := struct {
		Groups []objects.Group `json:"groups,omitempty"`
		Status *objects.Status `json:"status,omitempty"`
	}{}

	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}

	g.groups = temp.Groups
	g.status = temp.Status

	return nil
}

// Groups returns the response's groups.
func (g *GroupSuccessResponse) Groups() []objects.Group {
	groups := make([]objects.Group, len(g.groups), len(g.groups))
	copy(groups, g.groups)

	return groups
}

// Status returns the response's status.
func (g *GroupSuccessResponse) Status() *objects.Status {
	return g.status
}

type meterReportSuccessResponse struct {
	reportList []objects.MeterReport
	status     *objects.Status
}

// MeterReportSuccessResponse describes the success response returned by the
// ecobee server while retrieving a meter report.
type MeterReportSuccessResponse struct {
	meterReportSuccessResponse
}

// String implements the fmt.Stringer interface. It returns a string
// representing an indented JSON encoding of the response.
func (m *MeterReportSuccessResponse) String() string {
	temp := struct {
		ReportList []objects.MeterReport `json:""`
		Status     *objects.Status       `json:""`
	}{
		ReportList: m.reportList,
		Status:     m.status,
	}

	data, _ := json.MarshalIndent(&temp, "", "    ")

	return string(data)
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (m *MeterReportSuccessResponse) UnmarshalJSON(data []byte) error {
	temp := struct {
		ReportList []objects.MeterReport `json:"reportList,omitempty"`
		Status     *objects.Status       `json:"status,omitempty"`
	}{}

	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}

	m.reportList = temp.ReportList
	m.status = temp.Status

	return nil
}

// ReportList returns the response's status report list.
func (m *MeterReportSuccessResponse) ReportList() []objects.MeterReport {
	reportList := make([]objects.MeterReport, len(m.reportList), len(m.reportList))
	copy(reportList, m.reportList)

	return reportList
}

// Status returns the response's status.
func (m *MeterReportSuccessResponse) Status() *objects.Status {
	return m.status
}

type pinAuthorizationSuccessResponse struct {
	authorizationToken string
	expiresIn          int
	pin                string
	pollingInterval    int
	scope              Scope
}

// PINAuthorizationSuccessResponse describes the success response returned by
// the ecobee server while requesting the PIN.
type PINAuthorizationSuccessResponse struct {
	pinAuthorizationSuccessResponse
}

// String implements the fmt.Stringer interface. It returns a string
// representing an indented JSON encoding of the response.
func (a *PINAuthorizationSuccessResponse) String() string {
	temp := struct {
		AuthorizationToken string `json:""`
		ExpiresIn          int    `json:""`
		PIN                string `json:""`
		PollingInterval    int    `json:""`
		Scope              Scope  `json:""`
	}{
		AuthorizationToken: a.authorizationToken,
		ExpiresIn:          a.expiresIn,
		PIN:                a.pin,
		PollingInterval:    a.pollingInterval,
		Scope:              a.scope,
	}

	data, _ := json.MarshalIndent(&temp, "", "    ")

	return string(data)
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (a *PINAuthorizationSuccessResponse) UnmarshalJSON(data []byte) error {
	temp := struct {
		AuthorizationToken string `json:"code"`
		ExpiresIn          int    `json:"expires_in"`
		PIN                string `json:"ecobeePin"`
		PollingInterval    int    `json:"interval"`
		Scope              Scope  `json:"scope"`
	}{}

	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}

	a.authorizationToken = temp.AuthorizationToken
	a.expiresIn = temp.ExpiresIn
	a.pin = temp.PIN
	a.pollingInterval = temp.PollingInterval
	a.scope = temp.Scope

	return nil
}

// AuthorizationToken returns the response's authorization token.
func (a *PINAuthorizationSuccessResponse) AuthorizationToken() string {
	return a.authorizationToken
}

// ExpiresIn returns the response's expires in.
func (a *PINAuthorizationSuccessResponse) ExpiresIn() int {
	return a.expiresIn
}

// PIN returns the response's PIN.
func (a *PINAuthorizationSuccessResponse) PIN() string {
	return a.pin
}

// PollingInterval returns the response's polling interval.
func (a *PINAuthorizationSuccessResponse) PollingInterval() int {
	return a.pollingInterval
}

// Scope returns the response's scope.
func (a *PINAuthorizationSuccessResponse) Scope() Scope {
	return a.scope
}

type runtimeReportSuccessResponse struct {
	reportList []objects.RuntimeReport
	sensorList []objects.RuntimeSensorReport
	status     *objects.Status
}

// RuntimeReportSuccessResponse describes the success response returned by the
// ecobee server while retrieving a runtime report.
type RuntimeReportSuccessResponse struct {
	runtimeReportSuccessResponse
}

// String implements the fmt.Stringer interface. It returns a string
// representing an indented JSON encoding of the response.
func (m *RuntimeReportSuccessResponse) String() string {
	temp := struct {
		ReportList []objects.RuntimeReport       `json:""`
		SensorList []objects.RuntimeSensorReport `json:""`
		Status     *objects.Status               `json:""`
	}{
		ReportList: m.reportList,
		SensorList: m.sensorList,
		Status:     m.status,
	}

	data, _ := json.MarshalIndent(&temp, "", "    ")

	return string(data)
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (m *RuntimeReportSuccessResponse) UnmarshalJSON(data []byte) error {
	temp := struct {
		ReportList []objects.RuntimeReport       `json:"reportList,omitempty"`
		SensorList []objects.RuntimeSensorReport `json:"sensorList,omitempty"`
		Status     *objects.Status               `json:"status,omitempty"`
	}{}

	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}

	m.reportList = temp.ReportList
	m.sensorList = temp.SensorList
	m.status = temp.Status

	return nil
}

// ReportList returns the response's status report list.
func (m *RuntimeReportSuccessResponse) ReportList() []objects.RuntimeReport {
	reportList := make([]objects.RuntimeReport, len(m.reportList), len(m.reportList))
	copy(reportList, m.reportList)

	return reportList
}

// SensorList returns the response's status sensor list.
func (m *RuntimeReportSuccessResponse) SensorList() []objects.RuntimeSensorReport {
	sensorList := make([]objects.RuntimeSensorReport, len(m.sensorList), len(m.sensorList))
	copy(sensorList, m.sensorList)

	return sensorList
}

// Status returns the response's status.
func (m *RuntimeReportSuccessResponse) Status() *objects.Status {
	return m.status
}

type thermostatSuccessResponse struct {
	page           *objects.Page
	thermostatList []objects.Thermostat
	status         *objects.Status
}

// ThermostatSuccessResponse describes the success response returned by the
// ecobee server while retrieving thermostat data.
type ThermostatSuccessResponse struct {
	thermostatSuccessResponse
}

// String implements the fmt.Stringer interface. It returns a string
// representing an indented JSON encoding of the response.
func (t *ThermostatSuccessResponse) String() string {
	temp := struct {
		Page           *objects.Page        `json:""`
		ThermostatList []objects.Thermostat `json:""`
		Status         *objects.Status      `json:""`
	}{
		Page:           t.page,
		ThermostatList: t.thermostatList,
		Status:         t.status,
	}

	data, _ := json.MarshalIndent(&temp, "", "    ")

	return string(data)
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (t *ThermostatSuccessResponse) UnmarshalJSON(data []byte) error {
	temp := struct {
		Page           *objects.Page        `json:"page,omitempty"`
		ThermostatList []objects.Thermostat `json:"thermostatList,omitempty"`
		Status         *objects.Status      `json:"status,omitempty"`
	}{}

	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}

	t.page = temp.Page
	t.thermostatList = temp.ThermostatList
	t.status = temp.Status

	return nil
}

// Page returns the response's page.
func (t *ThermostatSuccessResponse) Page() *objects.Page {
	return t.page
}

// ThermostatList returns the response's thermostat list.
func (t *ThermostatSuccessResponse) ThermostatList() []objects.Thermostat {
	thermostatList := make([]objects.Thermostat, len(t.thermostatList), len(t.thermostatList))
	copy(thermostatList, t.thermostatList)

	return thermostatList
}

// Status returns the response's status.
func (t *ThermostatSuccessResponse) Status() *objects.Status {
	return t.status
}

type thermostatSummarySuccessResponse struct {
	revisionList    []string
	thermostatCount *int
	statusList      []string
	status          *objects.Status
}

// ThermostatSummarySuccessResponse describes the success response returned by
// the ecobee server while retrieving thermostat configuration and state
// revisions.
type ThermostatSummarySuccessResponse struct {
	thermostatSummarySuccessResponse
}

// String implements the fmt.Stringer interface. It returns a string
// representing an indented JSON encoding of the response.
func (t *ThermostatSummarySuccessResponse) String() string {
	temp := struct {
		RevisionList    []string        `json:""`
		ThermostatCount *int            `json:""`
		StatusList      []string        `json:""`
		Status          *objects.Status `json:""`
	}{
		RevisionList:    t.revisionList,
		ThermostatCount: t.thermostatCount,
		StatusList:      t.statusList,
		Status:          t.status,
	}

	data, _ := json.MarshalIndent(&temp, "", "    ")

	return string(data)
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (t *ThermostatSummarySuccessResponse) UnmarshalJSON(data []byte) error {
	temp := struct {
		RevisionList    []string        `json:"revisionList,omitempty"`
		ThermostatCount *int            `json:"thermostatCount,omitempty"`
		StatusList      []string        `json:"statusList,omitempty"`
		Status          *objects.Status `json:"status,omitempty"`
	}{}

	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}

	t.revisionList = temp.RevisionList
	t.thermostatCount = temp.ThermostatCount
	t.statusList = temp.StatusList
	t.status = temp.Status

	return nil
}

// RevisionList returns the response's revision list.
func (t *ThermostatSummarySuccessResponse) RevisionList() []string {
	revisionList := make([]string, len(t.revisionList), len(t.revisionList))
	copy(revisionList, t.revisionList)

	return revisionList
}

// ThermostatCount returns the response's thermostat count.
func (t *ThermostatSummarySuccessResponse) ThermostatCount() int {
	return *(t.thermostatCount)
}

// StatusList returns the response's status list.
func (t *ThermostatSummarySuccessResponse) StatusList() []string {
	statusList := make([]string, len(t.statusList), len(t.statusList))
	copy(statusList, t.statusList)

	return statusList
}

// Status returns the response's status.
func (t *ThermostatSummarySuccessResponse) Status() *objects.Status {
	return t.status
}

type tokensSuccessResponse struct {
	accessToken  string
	tokenType    string
	expiresIn    int
	refreshToken string
	scope        Scope
}

// TokensSuccessResponse describes the success response returned by the ecobee
// server while requesting tokens.
type TokensSuccessResponse struct {
	tokensSuccessResponse
}

// String implements the fmt.Stringer interface. It returns a string
// representing an indented JSON encoding of the response.
func (t *TokensSuccessResponse) String() string {
	temp := struct {
		AccessToken  string `json:"access_token"`
		TokenType    string `json:"token_type"`
		ExpiresIn    int    `json:"expires_in"`
		RefreshToken string `json:"refresh_token"`
		Scope        Scope  `json:"scope"`
	}{
		AccessToken:  t.accessToken,
		TokenType:    t.tokenType,
		ExpiresIn:    t.expiresIn,
		RefreshToken: t.refreshToken,
		Scope:        t.scope,
	}

	data, _ := json.MarshalIndent(&temp, "", "    ")

	return string(data)
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (t *TokensSuccessResponse) UnmarshalJSON(data []byte) error {
	temp := struct {
		AccessToken  string `json:"access_token"`
		TokenType    string `json:"token_type"`
		ExpiresIn    int    `json:"expires_in"`
		RefreshToken string `json:"refresh_token"`
		Scope        Scope  `json:"scope"`
	}{}

	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}

	t.accessToken = temp.AccessToken
	t.tokenType = temp.TokenType
	t.expiresIn = temp.ExpiresIn
	t.refreshToken = temp.RefreshToken
	t.scope = temp.Scope

	return nil
}

// AccessToken returns the response's access token.
func (t *TokensSuccessResponse) AccessToken() string {
	return t.accessToken
}

// TokenType returns the response's token type.
func (t *TokensSuccessResponse) TokenType() string {
	return t.tokenType
}

// ExpiresIn returns the response's expires in.
func (t *TokensSuccessResponse) ExpiresIn() int {
	return t.expiresIn
}

// RefreshToken returns the response's refresh token.
func (t *TokensSuccessResponse) RefreshToken() string {
	return t.refreshToken
}

// Scope returns the response's scope.
func (t *tokensSuccessResponse) Scope() Scope {
	return t.scope
}

func processAPIResponse(endpoint string, resp *http.Response, responseObject interface{}) error {
	if resp.StatusCode != http.StatusOK {
		errorResponse := APIStatusResponse{}

		switch err := json.NewDecoder(resp.Body).Decode(&errorResponse); {
		case err == io.EOF:
			return fmt.Errorf("%s: %s %q: %s", endpoint, resp.Request.Method, resp.Request.URL.String(), resp.Status)
		case err != nil:
			var e *json.SyntaxError

			if errors.As(err, &e) {
				return fmt.Errorf("%s: %s %q: %s", endpoint, resp.Request.Method, resp.Request.URL.String(), resp.Status)
			}

			return fmt.Errorf("%s: %s %q: %s: %w", endpoint, resp.Request.Method, resp.Request.URL.String(), resp.Status, err)
		}

		return &APIError{errorString: fmt.Sprintf("%s: %s %q: %s: code %d: %s", endpoint, resp.Request.Method, resp.Request.URL.String(), resp.Status, *errorResponse.status.Code, *errorResponse.status.Message)}
	}

	if err := json.NewDecoder(resp.Body).Decode(responseObject); err != nil {
		return fmt.Errorf("%s: %s %q: %s: %w", endpoint, resp.Request.Method, resp.Request.URL.String(), resp.Status, err)
	}

	return nil
}

func processAuthorizationResponse(endpoint string, resp *http.Response, responseObject interface{}) error {
	if resp.StatusCode != http.StatusOK {
		errorResponse := AuthorizationErrorResponse{}

		switch err := json.NewDecoder(resp.Body).Decode(&errorResponse); {
		case err == io.EOF:
			return fmt.Errorf("%s: %s %q: %s", endpoint, resp.Request.Method, resp.Request.URL.String(), resp.Status)
		case err != nil:
			var e *json.SyntaxError

			if errors.As(err, &e) {
				return fmt.Errorf("%s: %s %q: %s", endpoint, resp.Request.Method, resp.Request.URL.String(), resp.Status)
			}

			return fmt.Errorf("%s: %s %q: %s: %w", endpoint, resp.Request.Method, resp.Request.URL.String(), resp.Status, err)
		}

		return &AuthorizationError{errorString: fmt.Sprintf("%s: %s %q: %s: %s: %s: %s", endpoint, resp.Request.Method, resp.Request.URL.String(), resp.Status, errorResponse.errorType, errorResponse.errorDescription, errorResponse.errorURI)}
	}

	if err := json.NewDecoder(resp.Body).Decode(responseObject); err != nil {
		return fmt.Errorf("%s: %s %q: %s: %w", endpoint, resp.Request.Method, resp.Request.URL.String(), resp.Status, err)
	}

	return nil
}
