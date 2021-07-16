// Code generated by go-swagger; DO NOT EDIT.

package platformclientmodels

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// CategoryCreate A DTO object for creating category API call.
//
// swagger:model CategoryCreate
type CategoryCreate struct {

	// Category Path, A path separated by "/", start with "/" and end with combination of case of letters and numbers, max length is 255, min length is 2
	// Required: true
	CategoryPath *string `json:"categoryPath"`

	// Display name, key is language, value is display name, value max length is 255
	// Required: true
	LocalizationDisplayNames map[string]string `json:"localizationDisplayNames"`
}

// Validate validates this category create
func (m *CategoryCreate) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCategoryPath(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLocalizationDisplayNames(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CategoryCreate) validateCategoryPath(formats strfmt.Registry) error {

	if err := validate.Required("categoryPath", "body", m.CategoryPath); err != nil {
		return err
	}

	return nil
}

func (m *CategoryCreate) validateLocalizationDisplayNames(formats strfmt.Registry) error {

	return nil
}

// MarshalBinary interface implementation
func (m *CategoryCreate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CategoryCreate) UnmarshalBinary(b []byte) error {
	var res CategoryCreate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
