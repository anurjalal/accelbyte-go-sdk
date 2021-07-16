// Code generated by go-swagger; DO NOT EDIT.

package platformclientmodels

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PaymentOrderSyncResult payment order sync result
//
// swagger:model PaymentOrderSyncResult
type PaymentOrderSyncResult struct {

	// next evaluated key
	NextEvaluatedKey string `json:"nextEvaluatedKey,omitempty"`

	// payment orders
	PaymentOrders []*PaymentOrder `json:"paymentOrders"`
}

// Validate validates this payment order sync result
func (m *PaymentOrderSyncResult) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePaymentOrders(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PaymentOrderSyncResult) validatePaymentOrders(formats strfmt.Registry) error {

	if swag.IsZero(m.PaymentOrders) { // not required
		return nil
	}

	for i := 0; i < len(m.PaymentOrders); i++ {
		if swag.IsZero(m.PaymentOrders[i]) { // not required
			continue
		}

		if m.PaymentOrders[i] != nil {
			if err := m.PaymentOrders[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("paymentOrders" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *PaymentOrderSyncResult) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PaymentOrderSyncResult) UnmarshalBinary(b []byte) error {
	var res PaymentOrderSyncResult
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
