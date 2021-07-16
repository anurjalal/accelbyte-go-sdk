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

// PlayStationReconcileResult play station reconcile result
//
// swagger:model PlayStationReconcileResult
type PlayStationReconcileResult struct {

	// item Id
	ItemID string `json:"itemId,omitempty"`

	// psn item Id
	PsnItemID string `json:"psnItemId,omitempty"`

	// sku
	Sku string `json:"sku,omitempty"`

	// status
	// Enum: [VERIFIED FULFILLED FAILED]
	Status string `json:"status,omitempty"`

	// transaction Id
	TransactionID string `json:"transactionId,omitempty"`
}

// Validate validates this play station reconcile result
func (m *PlayStationReconcileResult) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var playStationReconcileResultTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["VERIFIED","FULFILLED","FAILED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		playStationReconcileResultTypeStatusPropEnum = append(playStationReconcileResultTypeStatusPropEnum, v)
	}
}

const (

	// PlayStationReconcileResultStatusVERIFIED captures enum value "VERIFIED"
	PlayStationReconcileResultStatusVERIFIED string = "VERIFIED"

	// PlayStationReconcileResultStatusFULFILLED captures enum value "FULFILLED"
	PlayStationReconcileResultStatusFULFILLED string = "FULFILLED"

	// PlayStationReconcileResultStatusFAILED captures enum value "FAILED"
	PlayStationReconcileResultStatusFAILED string = "FAILED"
)

// prop value enum
func (m *PlayStationReconcileResult) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, playStationReconcileResultTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *PlayStationReconcileResult) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PlayStationReconcileResult) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PlayStationReconcileResult) UnmarshalBinary(b []byte) error {
	var res PlayStationReconcileResult
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
