// Code generated by go-swagger; DO NOT EDIT.

package groupclientmodels

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ResponseErrorResponse response error response
//
// swagger:model response.ErrorResponse
type ResponseErrorResponse struct {

	// error code
	// Required: true
	ErrorCode *int32 `json:"errorCode"`

	// error message
	// Required: true
	ErrorMessage *string `json:"errorMessage"`
}

// Validate validates this response error response
func (m *ResponseErrorResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateErrorCode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateErrorMessage(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ResponseErrorResponse) validateErrorCode(formats strfmt.Registry) error {

	if err := validate.Required("errorCode", "body", m.ErrorCode); err != nil {
		return err
	}

	return nil
}

func (m *ResponseErrorResponse) validateErrorMessage(formats strfmt.Registry) error {

	if err := validate.Required("errorMessage", "body", m.ErrorMessage); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ResponseErrorResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ResponseErrorResponse) UnmarshalBinary(b []byte) error {
	var res ResponseErrorResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
