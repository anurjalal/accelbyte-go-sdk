// Code generated by go-swagger; DO NOT EDIT.

package dsmcclientmodels

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ModelsCreatePodConfigRequest models create pod config request
//
// swagger:model models.CreatePodConfigRequest
type ModelsCreatePodConfigRequest struct {

	// cpu limit
	// Required: true
	CPULimit *int32 `json:"cpu_limit"`

	// mem limit
	// Required: true
	MemLimit *int32 `json:"mem_limit"`

	// params
	// Required: true
	Params *string `json:"params"`
}

// Validate validates this models create pod config request
func (m *ModelsCreatePodConfigRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCPULimit(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMemLimit(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateParams(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ModelsCreatePodConfigRequest) validateCPULimit(formats strfmt.Registry) error {

	if err := validate.Required("cpu_limit", "body", m.CPULimit); err != nil {
		return err
	}

	return nil
}

func (m *ModelsCreatePodConfigRequest) validateMemLimit(formats strfmt.Registry) error {

	if err := validate.Required("mem_limit", "body", m.MemLimit); err != nil {
		return err
	}

	return nil
}

func (m *ModelsCreatePodConfigRequest) validateParams(formats strfmt.Registry) error {

	if err := validate.Required("params", "body", m.Params); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ModelsCreatePodConfigRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ModelsCreatePodConfigRequest) UnmarshalBinary(b []byte) error {
	var res ModelsCreatePodConfigRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
