// Code generated by go-swagger; DO NOT EDIT.

package platformclientmodels

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// FulfillmentHistoryInfo fulfillment history info
//
// swagger:model FulfillmentHistoryInfo
type FulfillmentHistoryInfo struct {

	// redeemed code
	Code string `json:"code,omitempty"`

	// created at
	// Required: true
	// Format: date-time
	CreatedAt *strfmt.DateTime `json:"createdAt"`

	// credit granted
	CreditSummaries []*CreditSummary `json:"creditSummaries"`

	// entitlement granted
	EntitlementSummaries []*EntitlementSummary `json:"entitlementSummaries"`

	// items should do fulfillment
	FulfillItems []*FulfillmentItem `json:"fulfillItems"`

	// fulfillment error detail
	FulfillmentError *FulfillmentError `json:"fulfillmentError,omitempty"`

	// item ids already granted
	GrantedItemIds []string `json:"grantedItemIds"`

	// id
	// Required: true
	ID *string `json:"id"`

	// namespace
	// Required: true
	Namespace *string `json:"namespace"`

	// orderNo
	OrderNo string `json:"orderNo,omitempty"`

	// fulfillment status
	// Required: true
	// Enum: [SUCCESS FAIL]
	Status *string `json:"status"`

	// updated at
	// Required: true
	// Format: date-time
	UpdatedAt *strfmt.DateTime `json:"updatedAt"`

	// userId
	// Required: true
	UserID *string `json:"userId"`
}

// Validate validates this fulfillment history info
func (m *FulfillmentHistoryInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreditSummaries(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEntitlementSummaries(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFulfillItems(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFulfillmentError(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNamespace(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUserID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FulfillmentHistoryInfo) validateCreatedAt(formats strfmt.Registry) error {

	if err := validate.Required("createdAt", "body", m.CreatedAt); err != nil {
		return err
	}

	if err := validate.FormatOf("createdAt", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *FulfillmentHistoryInfo) validateCreditSummaries(formats strfmt.Registry) error {

	if swag.IsZero(m.CreditSummaries) { // not required
		return nil
	}

	for i := 0; i < len(m.CreditSummaries); i++ {
		if swag.IsZero(m.CreditSummaries[i]) { // not required
			continue
		}

		if m.CreditSummaries[i] != nil {
			if err := m.CreditSummaries[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("creditSummaries" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *FulfillmentHistoryInfo) validateEntitlementSummaries(formats strfmt.Registry) error {

	if swag.IsZero(m.EntitlementSummaries) { // not required
		return nil
	}

	for i := 0; i < len(m.EntitlementSummaries); i++ {
		if swag.IsZero(m.EntitlementSummaries[i]) { // not required
			continue
		}

		if m.EntitlementSummaries[i] != nil {
			if err := m.EntitlementSummaries[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("entitlementSummaries" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *FulfillmentHistoryInfo) validateFulfillItems(formats strfmt.Registry) error {

	if swag.IsZero(m.FulfillItems) { // not required
		return nil
	}

	for i := 0; i < len(m.FulfillItems); i++ {
		if swag.IsZero(m.FulfillItems[i]) { // not required
			continue
		}

		if m.FulfillItems[i] != nil {
			if err := m.FulfillItems[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("fulfillItems" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *FulfillmentHistoryInfo) validateFulfillmentError(formats strfmt.Registry) error {

	if swag.IsZero(m.FulfillmentError) { // not required
		return nil
	}

	if m.FulfillmentError != nil {
		if err := m.FulfillmentError.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("fulfillmentError")
			}
			return err
		}
	}

	return nil
}

func (m *FulfillmentHistoryInfo) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *FulfillmentHistoryInfo) validateNamespace(formats strfmt.Registry) error {

	if err := validate.Required("namespace", "body", m.Namespace); err != nil {
		return err
	}

	return nil
}

var fulfillmentHistoryInfoTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["SUCCESS","FAIL"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		fulfillmentHistoryInfoTypeStatusPropEnum = append(fulfillmentHistoryInfoTypeStatusPropEnum, v)
	}
}

const (

	// FulfillmentHistoryInfoStatusSUCCESS captures enum value "SUCCESS"
	FulfillmentHistoryInfoStatusSUCCESS string = "SUCCESS"

	// FulfillmentHistoryInfoStatusFAIL captures enum value "FAIL"
	FulfillmentHistoryInfoStatusFAIL string = "FAIL"
)

// prop value enum
func (m *FulfillmentHistoryInfo) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, fulfillmentHistoryInfoTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *FulfillmentHistoryInfo) validateStatus(formats strfmt.Registry) error {

	if err := validate.Required("status", "body", m.Status); err != nil {
		return err
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", *m.Status); err != nil {
		return err
	}

	return nil
}

func (m *FulfillmentHistoryInfo) validateUpdatedAt(formats strfmt.Registry) error {

	if err := validate.Required("updatedAt", "body", m.UpdatedAt); err != nil {
		return err
	}

	if err := validate.FormatOf("updatedAt", "body", "date-time", m.UpdatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *FulfillmentHistoryInfo) validateUserID(formats strfmt.Registry) error {

	if err := validate.Required("userId", "body", m.UserID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *FulfillmentHistoryInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FulfillmentHistoryInfo) UnmarshalBinary(b []byte) error {
	var res FulfillmentHistoryInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
