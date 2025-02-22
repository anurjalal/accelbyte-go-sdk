// Code generated by go-swagger; DO NOT EDIT.

package iamclientmodels

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ModelUserCreateRequestV3 model user create request v3
//
// swagger:model model.UserCreateRequestV3
type ModelUserCreateRequestV3 struct {

	// password m d5 sum
	PasswordMD5Sum string `json:"PasswordMD5Sum,omitempty"`

	// accepted policies
	AcceptedPolicies []*LegalAcceptedPoliciesRequest `json:"acceptedPolicies"`

	// auth type
	// Required: true
	AuthType *string `json:"authType"`

	// country
	// Required: true
	Country *string `json:"country"`

	// date of birth
	// Required: true
	DateOfBirth *string `json:"dateOfBirth"`

	// display name
	// Required: true
	DisplayName *string `json:"displayName"`

	// email address
	// Required: true
	EmailAddress *string `json:"emailAddress"`

	// password
	// Required: true
	Password *string `json:"password"`
}

// Validate validates this model user create request v3
func (m *ModelUserCreateRequestV3) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAcceptedPolicies(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAuthType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCountry(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDateOfBirth(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDisplayName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEmailAddress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePassword(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ModelUserCreateRequestV3) validateAcceptedPolicies(formats strfmt.Registry) error {

	if swag.IsZero(m.AcceptedPolicies) { // not required
		return nil
	}

	for i := 0; i < len(m.AcceptedPolicies); i++ {
		if swag.IsZero(m.AcceptedPolicies[i]) { // not required
			continue
		}

		if m.AcceptedPolicies[i] != nil {
			if err := m.AcceptedPolicies[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("acceptedPolicies" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ModelUserCreateRequestV3) validateAuthType(formats strfmt.Registry) error {

	if err := validate.Required("authType", "body", m.AuthType); err != nil {
		return err
	}

	return nil
}

func (m *ModelUserCreateRequestV3) validateCountry(formats strfmt.Registry) error {

	if err := validate.Required("country", "body", m.Country); err != nil {
		return err
	}

	return nil
}

func (m *ModelUserCreateRequestV3) validateDateOfBirth(formats strfmt.Registry) error {

	if err := validate.Required("dateOfBirth", "body", m.DateOfBirth); err != nil {
		return err
	}

	return nil
}

func (m *ModelUserCreateRequestV3) validateDisplayName(formats strfmt.Registry) error {

	if err := validate.Required("displayName", "body", m.DisplayName); err != nil {
		return err
	}

	return nil
}

func (m *ModelUserCreateRequestV3) validateEmailAddress(formats strfmt.Registry) error {

	if err := validate.Required("emailAddress", "body", m.EmailAddress); err != nil {
		return err
	}

	return nil
}

func (m *ModelUserCreateRequestV3) validatePassword(formats strfmt.Registry) error {

	if err := validate.Required("password", "body", m.Password); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ModelUserCreateRequestV3) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ModelUserCreateRequestV3) UnmarshalBinary(b []byte) error {
	var res ModelUserCreateRequestV3
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
