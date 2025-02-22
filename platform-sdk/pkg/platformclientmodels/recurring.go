// Code generated by go-swagger; DO NOT EDIT.

package platformclientmodels

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Recurring recurring
//
// swagger:model Recurring
type Recurring struct {

	// recurring cycle
	// Required: true
	// Enum: [WEEKLY MONTHLY QUARTERLY YEARLY]
	Cycle *string `json:"cycle"`

	// fixed free days, 0 means not set
	// Required: true
	FixedFreeDays *int32 `json:"fixedFreeDays"`

	// fixed trial cycles, 0 means not set
	// Required: true
	FixedTrialCycles *int32 `json:"fixedTrialCycles"`

	// recurring grace days, retry recurring charge within configured days if charge fail, default 7
	// Required: true
	GraceDays *int32 `json:"graceDays"`
}

// Validate validates this recurring
func (m *Recurring) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCycle(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFixedFreeDays(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFixedTrialCycles(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGraceDays(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var recurringTypeCyclePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["WEEKLY","MONTHLY","QUARTERLY","YEARLY"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		recurringTypeCyclePropEnum = append(recurringTypeCyclePropEnum, v)
	}
}

const (

	// RecurringCycleWEEKLY captures enum value "WEEKLY"
	RecurringCycleWEEKLY string = "WEEKLY"

	// RecurringCycleMONTHLY captures enum value "MONTHLY"
	RecurringCycleMONTHLY string = "MONTHLY"

	// RecurringCycleQUARTERLY captures enum value "QUARTERLY"
	RecurringCycleQUARTERLY string = "QUARTERLY"

	// RecurringCycleYEARLY captures enum value "YEARLY"
	RecurringCycleYEARLY string = "YEARLY"
)

// prop value enum
func (m *Recurring) validateCycleEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, recurringTypeCyclePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Recurring) validateCycle(formats strfmt.Registry) error {

	if err := validate.Required("cycle", "body", m.Cycle); err != nil {
		return err
	}

	// value enum
	if err := m.validateCycleEnum("cycle", "body", *m.Cycle); err != nil {
		return err
	}

	return nil
}

func (m *Recurring) validateFixedFreeDays(formats strfmt.Registry) error {

	if err := validate.Required("fixedFreeDays", "body", m.FixedFreeDays); err != nil {
		return err
	}

	return nil
}

func (m *Recurring) validateFixedTrialCycles(formats strfmt.Registry) error {

	if err := validate.Required("fixedTrialCycles", "body", m.FixedTrialCycles); err != nil {
		return err
	}

	return nil
}

func (m *Recurring) validateGraceDays(formats strfmt.Registry) error {

	if err := validate.Required("graceDays", "body", m.GraceDays); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Recurring) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Recurring) UnmarshalBinary(b []byte) error {
	var res Recurring
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
