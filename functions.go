package ecobee

import (
	"context"
	"time"

	"github.com/sherif-fanous/go-ecobee/objects"
)

// An AckType specifies an acknowledgment type.
type AckType string

// Supported AckType values.
const (
	AckTypeAccept         AckType = "accept"
	AckTypeDecline        AckType = "decline"
	AckTypeDefer          AckType = "defer"
	AckTypeUnacknowledged AckType = "unacknowledged"
)

// A FanMode specifies a fan mode.
type FanMode string

// Supported FanMode values.
const (
	FanModeAuto FanMode = "auto"
	FanModeOn   FanMode = "on"
)

// A HoldType specifies a hold type.
type HoldType string

// Supported HoldType values.
const (
	HoldTypeDateTime       HoldType = "dateTime"
	HoldTypeHoldHours      HoldType = "holdHours"
	HoldTypeIndefinite     HoldType = "indefinite"
	HoldTypeNextTransition HoldType = "nextTransition"
)

// A PlugState specifies a plug state.
type PlugState string

// Supported PlugState values.
const (
	PlugStateOff    PlugState = "off"
	PlugStateOn     PlugState = "on"
	PlugStateResume PlugState = "resume"
)

// An AcknowledgeParameters specifies the request parameters of the Acknowledge
// method.
type AcknowledgeParameters struct {
	// The thermostat identifier to acknowledge the alert for.
	ThermostatIdentifier *string
	// The acknowledge ref of alert.
	AckRef *string
	// The type of acknowledgement.
	AckType *AckType
	// Whether to remind at a later date, if this is a defer acknowledgement.
	RemindMeLater *bool
}

// Acknowledge allows an Alert to be acknowledged.
//
// For more information see: https://www.ecobee.com/home/developer/api/documentation/v1/functions/Acknowledge.shtml
func (c *Client) Acknowledge(ctx context.Context, selection *objects.Selection, parameters *AcknowledgeParameters) (*APIStatusResponse, error) {
	functions := []objects.Function{
		{
			Type: String("acknowledge"),
			Params: map[string]interface{}{
				"thermostatIdentifier": *parameters.ThermostatIdentifier,
				"ackRef":               *parameters.AckRef,
				"ackType":              *parameters.AckType,
			},
		},
	}

	if parameters.RemindMeLater != nil {
		functions[0].Params["remindMeLater"] = *parameters.RemindMeLater
	}

	return c.UpdateThermostat(ctx, selection, nil, functions)
}

// A ControlPlugParameters specifies the request parameters of the ControlPlug
// method.
type ControlPlugParameters struct {
	// The name of the plug.
	PlugName *string
	// The state to put the plug into.
	PlugState *PlugState
	// The start date & time in thermostat time.
	StartDateTime *time.Time
	// The end date & time in thermostat time.
	EndDateTime *time.Time
	// The hold duration type.
	HoldType *HoldType
	// The number of hours to hold for.
	HoldHours *int
}

// ControlPlug sets the on/off state of a plug by setting a hold on the plug.
//
// For more information see: https://www.ecobee.com/home/developer/api/documentation/v1/functions/ControlPlug.shtml
func (c *Client) ControlPlug(ctx context.Context, selection *objects.Selection, parameters *ControlPlugParameters) (*APIStatusResponse, error) {
	functions := []objects.Function{
		{
			Type: String("controlPlug"),
			Params: map[string]interface{}{
				"plugName":  *parameters.PlugName,
				"plugState": *parameters.PlugState,
			},
		},
	}

	if parameters.StartDateTime != nil {
		startDateTime := *parameters.StartDateTime

		functions[0].Params["startDate"] = startDateTime.Format("2006-01-02")
		functions[0].Params["startTime"] = startDateTime.Format("15:04:05")
	}

	if parameters.EndDateTime != nil {
		endDateTime := *parameters.EndDateTime

		functions[0].Params["endDate"] = endDateTime.Format("2006-01-02")
		functions[0].Params["endTime"] = endDateTime.Format("15:04:05")
	}

	if parameters.HoldType != nil {
		functions[0].Params["holdType"] = *parameters.HoldType
	}

	if parameters.HoldHours != nil {
		functions[0].Params["holdHours"] = *parameters.HoldHours
	}

	return c.UpdateThermostat(ctx, selection, nil, functions)
}

// A CreateVacationParameters specifies the request parameters of the
// CreateVacation method.
type CreateVacationParameters struct {
	// The vacation name to create.
	Name *string
	// The temperature to set the cool vacation hold at.
	CoolHoldTemp *int
	// The temperature to set the heat vacation hold at.
	HeatHoldTemp *int
	// The start date & time in thermostat time.
	StartDateTime *time.Time
	// The end date & time in thermostat time.
	EndDateTime *time.Time
	// The fan mode during the vacation.
	Fan *FanMode
	// The minimum number of minutes to run the fan each hour.
	FanMinOnTime *string
}

// CreateVacation creates a vacation event on the thermostat.
//
// For more information see: https://www.ecobee.com/home/developer/api/documentation/v1/functions/CreateVacation.shtml
func (c *Client) CreateVacation(ctx context.Context, selection *objects.Selection, parameters *CreateVacationParameters) (*APIStatusResponse, error) {
	functions := []objects.Function{
		{
			Type: String("createVacation"),
			Params: map[string]interface{}{
				"name":         *parameters.Name,
				"coolHoldTemp": *parameters.CoolHoldTemp,
				"heatHoldTemp": *parameters.HeatHoldTemp,
			},
		},
	}

	if parameters.StartDateTime != nil {
		startDateTime := *parameters.StartDateTime

		functions[0].Params["startDate"] = startDateTime.Format("2006-01-02")
		functions[0].Params["startTime"] = startDateTime.Format("15:04:05")
	}

	if parameters.EndDateTime != nil {
		endDateTime := *parameters.EndDateTime

		functions[0].Params["endDate"] = endDateTime.Format("2006-01-02")
		functions[0].Params["endTime"] = endDateTime.Format("15:04:05")
	}

	if parameters.Fan != nil {
		functions[0].Params["fan"] = *parameters.Fan
	}

	if parameters.FanMinOnTime != nil {
		functions[0].Params["fanMinOnTime"] = *parameters.FanMinOnTime
	}

	return c.UpdateThermostat(ctx, selection, nil, functions)
}

// A DeleteVacationParameters specifies the request parameters of the
// DeleteVacation method.
type DeleteVacationParameters struct {
	// The vacation name to delete.
	Name *string
}

// DeleteVacation deletes a vacation event from a thermostat.
//
// For more information see: https://www.ecobee.com/home/developer/api/documentation/v1/functions/DeleteVacation.shtml
func (c *Client) DeleteVacation(ctx context.Context, selection *objects.Selection, parameters *DeleteVacationParameters) (*APIStatusResponse, error) {
	return c.UpdateThermostat(ctx, selection, nil, []objects.Function{
		{
			Type: String("deleteVacation"),
			Params: map[string]interface{}{
				"name": *parameters.Name,
			},
		},
	})
}

// ResetPreferences sets all of the user configurable settings back to the
// factory default values.
//
// For more information see: https://www.ecobee.com/home/developer/api/documentation/v1/functions/ResetPreferences.shtml
func (c *Client) ResetPreferences(ctx context.Context, selection *objects.Selection) (*APIStatusResponse, error) {
	return c.UpdateThermostat(ctx, selection, nil, []objects.Function{
		{
			Type: String("resetPreferences"),
		},
	})
}

// A ResumeProgramParameters specifies the request parameters of the
// ResumeProgram method.
type ResumeProgramParameters struct {
	// Specifies whether the thermostat be resumed to next event (false) or to
	// it's program (true).
	ResumeAll *bool
}

// ResumeProgram removes the currently running event providing the event is
// not a mandatory demand response event.
//
// For more information see: https://www.ecobee.com/home/developer/api/documentation/v1/functions/ResumeProgram.shtml
func (c *Client) ResumeProgram(ctx context.Context, selection *objects.Selection, parameters *ResumeProgramParameters) (*APIStatusResponse, error) {
	return c.UpdateThermostat(ctx, selection, nil, []objects.Function{
		{
			Type: String("resumeProgram"),
			Params: map[string]interface{}{
				"resumeAll": *parameters.ResumeAll,
			},
		},
	})
}

// A SendMessageParameters specifies the request parameters of the
// SendMessage method.
type SendMessageParameters struct {
	// The message text to send.
	Text *string
}

// SendMessage allows an alert message to be sent to the thermostat.
//
// For more information see: https://www.ecobee.com/home/developer/api/documentation/v1/functions/SendMessage.shtml
func (c *Client) SendMessage(ctx context.Context, selection *objects.Selection, parameters *SendMessageParameters) (*APIStatusResponse, error) {
	return c.UpdateThermostat(ctx, selection, nil, []objects.Function{
		{
			Type: String("sendMessage"),
			Params: map[string]interface{}{
				"text": *parameters.Text,
			},
		},
	})
}

// A SetHoldParameters specifies the request parameters of the
// SetHold method.
type SetHoldParameters struct {
	// The temperature to set the cool hold at.
	CoolHoldTemp *int
	// The temperature to set the heat hold at.
	HeatHoldTemp *int
	// The Climate to use as reference for setting the coolHoldTemp,
	// heatHoldTemp and fan settings.
	HoldClimateRef *string
	// The start date & time in thermostat time.
	StartDateTime *time.Time
	// The end date & time in thermostat time.
	EndDateTime *time.Time
	// The hold duration type.
	HoldType *HoldType
	// The number of hours to hold for.
	HoldHours *int
}

// SetHold sets the thermostat into a hold with the specified temperature.
//
// For more information see: https://www.ecobee.com/home/developer/api/documentation/v1/functions/SetHold.shtml
func (c *Client) SetHold(ctx context.Context, selection *objects.Selection, parameters *SetHoldParameters) (*APIStatusResponse, error) {
	functions := []objects.Function{
		{
			Type:   String("setHold"),
			Params: map[string]interface{}{},
		},
	}

	if parameters.CoolHoldTemp != nil {
		functions[0].Params["coolHoldTemp"] = *parameters.CoolHoldTemp
	}

	if parameters.HeatHoldTemp != nil {
		functions[0].Params["heatHoldTemp"] = *parameters.HeatHoldTemp
	}

	if parameters.HoldClimateRef != nil {
		functions[0].Params["holdClimateRef"] = *parameters.HoldClimateRef
	}

	if parameters.StartDateTime != nil {
		startDateTime := *parameters.StartDateTime

		functions[0].Params["startDate"] = startDateTime.Format("2006-01-02")
		functions[0].Params["startTime"] = startDateTime.Format("15:04:05")
	}

	if parameters.EndDateTime != nil {
		endDateTime := *parameters.EndDateTime

		functions[0].Params["endDate"] = endDateTime.Format("2006-01-02")
		functions[0].Params["endTime"] = endDateTime.Format("15:04:05")
	}

	if parameters.HoldType != nil {
		functions[0].Params["holdType"] = *parameters.HoldType
	}

	if parameters.HoldHours != nil {
		functions[0].Params["holdHours"] = *parameters.HoldHours
	}

	return c.UpdateThermostat(ctx, selection, nil, functions)
}

// An UnlinkVoiceEngineParameters specifies the request parameters of the
// UnlinkVoiceEngine method.
type UnlinkVoiceEngineParameters struct {
	// The name of the engine to unlink.
	EngineName *string
}

// UnlinkVoiceEngine disables the voice assistant for the selected thermostat.
//
// For more information see: https://www.ecobee.com/home/developer/api/documentation/v1/functions/UnlinkVoiceEngine.shtml
func (c *Client) UnlinkVoiceEngine(ctx context.Context, selection *objects.Selection, parameters *UnlinkVoiceEngineParameters) (*APIStatusResponse, error) {
	return c.UpdateThermostat(ctx, selection, nil, []objects.Function{
		{
			Type: String("unlinkVoiceEngine"),
			Params: map[string]interface{}{
				"engineName": *parameters.EngineName,
			},
		},
	})
}

// An UpdateSensorParameters specifies the request parameters of the UpdateSensor
// method
type UpdateSensorParameters struct {
	// Name specifies the updated name to give the sensor.
	Name *string
	// DeviceID specifies the deviceId for the sensor.
	DeviceID *string
	// SensorID specifies the identifier for the sensor.
	SensorID *string
}

// UpdateSensor allows the caller to update the name of a remote sensor.
//
// For more information see: https://www.ecobee.com/home/developer/api/documentation/v1/functions/UpdateSensor.shtml
func (c *Client) UpdateSensor(ctx context.Context, selection *objects.Selection, parameters *UpdateSensorParameters) (*APIStatusResponse, error) {
	return c.UpdateThermostat(ctx, selection, nil, []objects.Function{
		{
			Type: String("updateSensor"),
			Params: map[string]interface{}{
				"name":     *parameters.Name,
				"deviceId": *parameters.DeviceID,
				"sensorId": *parameters.SensorID,
			},
		},
	})
}
