// Code generated by go-swagger; DO NOT EDIT.

package platformclientmodels

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// EpicGamesReconcileResult epic games reconcile result
//
// swagger:model EpicGamesReconcileResult
type EpicGamesReconcileResult struct {

	// Epic games item ID
	EpicGamesItemID string `json:"epicGamesItemId,omitempty"`

	// Item ID
	ItemID string `json:"itemId,omitempty"`

	// SKU
	Sku string `json:"sku,omitempty"`

	// IAP order status
	Status string `json:"status,omitempty"`

	// Transaction ID
	TransactionID string `json:"transactionId,omitempty"`
}

// Validate validates this epic games reconcile result
func (m *EpicGamesReconcileResult) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *EpicGamesReconcileResult) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EpicGamesReconcileResult) UnmarshalBinary(b []byte) error {
	var res EpicGamesReconcileResult
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
