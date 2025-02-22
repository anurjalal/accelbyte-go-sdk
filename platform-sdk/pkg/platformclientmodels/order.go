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

// Order order
//
// swagger:model Order
type Order struct {

	// chargeback reversed time
	// Format: date-time
	ChargebackReversedTime strfmt.DateTime `json:"chargebackReversedTime,omitempty"`

	// chargeback time
	// Format: date-time
	ChargebackTime strfmt.DateTime `json:"chargebackTime,omitempty"`

	// charged
	Charged bool `json:"charged,omitempty"`

	// charged time
	// Format: date-time
	ChargedTime strfmt.DateTime `json:"chargedTime,omitempty"`

	// count item Id
	CountItemID string `json:"countItemId,omitempty"`

	// count namespace
	CountNamespace string `json:"countNamespace,omitempty"`

	// count user Id
	CountUserID string `json:"countUserId,omitempty"`

	// created at
	// Format: date-time
	CreatedAt strfmt.DateTime `json:"createdAt,omitempty"`

	// created time
	// Format: date-time
	CreatedTime strfmt.DateTime `json:"createdTime,omitempty"`

	// currency
	Currency *CurrencySummary `json:"currency,omitempty"`

	// discounted price
	DiscountedPrice int32 `json:"discountedPrice,omitempty"`

	// expire time
	// Format: date-time
	ExpireTime strfmt.DateTime `json:"expireTime,omitempty"`

	// ext
	Ext map[string]interface{} `json:"ext,omitempty"`

	// free
	Free bool `json:"free,omitempty"`

	// fulfilled time
	// Format: date-time
	FulfilledTime strfmt.DateTime `json:"fulfilledTime,omitempty"`

	// item Id
	ItemID string `json:"itemId,omitempty"`

	// item snapshot
	ItemSnapshot *ItemSnapshot `json:"itemSnapshot,omitempty"`

	// language
	Language string `json:"language,omitempty"`

	// namespace
	Namespace string `json:"namespace,omitempty"`

	// order no
	OrderNo string `json:"orderNo,omitempty"`

	// payment method
	PaymentMethod string `json:"paymentMethod,omitempty"`

	// payment method fee
	PaymentMethodFee int32 `json:"paymentMethodFee,omitempty"`

	// payment order no
	PaymentOrderNo string `json:"paymentOrderNo,omitempty"`

	// payment provider
	// Enum: [WALLET XSOLLA ADYEN STRIPE CHECKOUT ALIPAY WXPAY PAYPAL]
	PaymentProvider string `json:"paymentProvider,omitempty"`

	// payment provider fee
	PaymentProviderFee int32 `json:"paymentProviderFee,omitempty"`

	// payment remain seconds
	PaymentRemainSeconds int32 `json:"paymentRemainSeconds,omitempty"`

	// payment station Url
	PaymentStationURL string `json:"paymentStationUrl,omitempty"`

	// price
	Price int32 `json:"price,omitempty"`

	// quantity
	Quantity int32 `json:"quantity,omitempty"`

	// refunded time
	// Format: date-time
	RefundedTime strfmt.DateTime `json:"refundedTime,omitempty"`

	// region
	Region string `json:"region,omitempty"`

	// return Url
	ReturnURL string `json:"returnUrl,omitempty"`

	// rvn
	Rvn int32 `json:"rvn,omitempty"`

	// sales tax
	SalesTax int32 `json:"salesTax,omitempty"`

	// sandbox
	Sandbox bool `json:"sandbox,omitempty"`

	// status
	// Enum: [INIT CHARGED CHARGEBACK CHARGEBACK_REVERSED FULFILLED FULFILL_FAILED REFUNDING REFUNDED REFUND_FAILED CLOSED DELETED]
	Status string `json:"status,omitempty"`

	// status reason
	StatusReason string `json:"statusReason,omitempty"`

	// subtotal price
	SubtotalPrice int32 `json:"subtotalPrice,omitempty"`

	// tax
	Tax int32 `json:"tax,omitempty"`

	// total price
	TotalPrice int32 `json:"totalPrice,omitempty"`

	// total tax
	TotalTax int32 `json:"totalTax,omitempty"`

	// updated at
	// Format: date-time
	UpdatedAt strfmt.DateTime `json:"updatedAt,omitempty"`

	// user Id
	UserID string `json:"userId,omitempty"`

	// vat
	Vat int32 `json:"vat,omitempty"`
}

// Validate validates this order
func (m *Order) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateChargebackReversedTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateChargebackTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateChargedTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCurrency(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExpireTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFulfilledTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateItemSnapshot(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePaymentProvider(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRefundedTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdatedAt(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Order) validateChargebackReversedTime(formats strfmt.Registry) error {

	if swag.IsZero(m.ChargebackReversedTime) { // not required
		return nil
	}

	if err := validate.FormatOf("chargebackReversedTime", "body", "date-time", m.ChargebackReversedTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Order) validateChargebackTime(formats strfmt.Registry) error {

	if swag.IsZero(m.ChargebackTime) { // not required
		return nil
	}

	if err := validate.FormatOf("chargebackTime", "body", "date-time", m.ChargebackTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Order) validateChargedTime(formats strfmt.Registry) error {

	if swag.IsZero(m.ChargedTime) { // not required
		return nil
	}

	if err := validate.FormatOf("chargedTime", "body", "date-time", m.ChargedTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Order) validateCreatedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.CreatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("createdAt", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Order) validateCreatedTime(formats strfmt.Registry) error {

	if swag.IsZero(m.CreatedTime) { // not required
		return nil
	}

	if err := validate.FormatOf("createdTime", "body", "date-time", m.CreatedTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Order) validateCurrency(formats strfmt.Registry) error {

	if swag.IsZero(m.Currency) { // not required
		return nil
	}

	if m.Currency != nil {
		if err := m.Currency.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("currency")
			}
			return err
		}
	}

	return nil
}

func (m *Order) validateExpireTime(formats strfmt.Registry) error {

	if swag.IsZero(m.ExpireTime) { // not required
		return nil
	}

	if err := validate.FormatOf("expireTime", "body", "date-time", m.ExpireTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Order) validateFulfilledTime(formats strfmt.Registry) error {

	if swag.IsZero(m.FulfilledTime) { // not required
		return nil
	}

	if err := validate.FormatOf("fulfilledTime", "body", "date-time", m.FulfilledTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Order) validateItemSnapshot(formats strfmt.Registry) error {

	if swag.IsZero(m.ItemSnapshot) { // not required
		return nil
	}

	if m.ItemSnapshot != nil {
		if err := m.ItemSnapshot.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("itemSnapshot")
			}
			return err
		}
	}

	return nil
}

var orderTypePaymentProviderPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["WALLET","XSOLLA","ADYEN","STRIPE","CHECKOUT","ALIPAY","WXPAY","PAYPAL"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		orderTypePaymentProviderPropEnum = append(orderTypePaymentProviderPropEnum, v)
	}
}

const (

	// OrderPaymentProviderWALLET captures enum value "WALLET"
	OrderPaymentProviderWALLET string = "WALLET"

	// OrderPaymentProviderXSOLLA captures enum value "XSOLLA"
	OrderPaymentProviderXSOLLA string = "XSOLLA"

	// OrderPaymentProviderADYEN captures enum value "ADYEN"
	OrderPaymentProviderADYEN string = "ADYEN"

	// OrderPaymentProviderSTRIPE captures enum value "STRIPE"
	OrderPaymentProviderSTRIPE string = "STRIPE"

	// OrderPaymentProviderCHECKOUT captures enum value "CHECKOUT"
	OrderPaymentProviderCHECKOUT string = "CHECKOUT"

	// OrderPaymentProviderALIPAY captures enum value "ALIPAY"
	OrderPaymentProviderALIPAY string = "ALIPAY"

	// OrderPaymentProviderWXPAY captures enum value "WXPAY"
	OrderPaymentProviderWXPAY string = "WXPAY"

	// OrderPaymentProviderPAYPAL captures enum value "PAYPAL"
	OrderPaymentProviderPAYPAL string = "PAYPAL"
)

// prop value enum
func (m *Order) validatePaymentProviderEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, orderTypePaymentProviderPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Order) validatePaymentProvider(formats strfmt.Registry) error {

	if swag.IsZero(m.PaymentProvider) { // not required
		return nil
	}

	// value enum
	if err := m.validatePaymentProviderEnum("paymentProvider", "body", m.PaymentProvider); err != nil {
		return err
	}

	return nil
}

func (m *Order) validateRefundedTime(formats strfmt.Registry) error {

	if swag.IsZero(m.RefundedTime) { // not required
		return nil
	}

	if err := validate.FormatOf("refundedTime", "body", "date-time", m.RefundedTime.String(), formats); err != nil {
		return err
	}

	return nil
}

var orderTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["INIT","CHARGED","CHARGEBACK","CHARGEBACK_REVERSED","FULFILLED","FULFILL_FAILED","REFUNDING","REFUNDED","REFUND_FAILED","CLOSED","DELETED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		orderTypeStatusPropEnum = append(orderTypeStatusPropEnum, v)
	}
}

const (

	// OrderStatusINIT captures enum value "INIT"
	OrderStatusINIT string = "INIT"

	// OrderStatusCHARGED captures enum value "CHARGED"
	OrderStatusCHARGED string = "CHARGED"

	// OrderStatusCHARGEBACK captures enum value "CHARGEBACK"
	OrderStatusCHARGEBACK string = "CHARGEBACK"

	// OrderStatusCHARGEBACKREVERSED captures enum value "CHARGEBACK_REVERSED"
	OrderStatusCHARGEBACKREVERSED string = "CHARGEBACK_REVERSED"

	// OrderStatusFULFILLED captures enum value "FULFILLED"
	OrderStatusFULFILLED string = "FULFILLED"

	// OrderStatusFULFILLFAILED captures enum value "FULFILL_FAILED"
	OrderStatusFULFILLFAILED string = "FULFILL_FAILED"

	// OrderStatusREFUNDING captures enum value "REFUNDING"
	OrderStatusREFUNDING string = "REFUNDING"

	// OrderStatusREFUNDED captures enum value "REFUNDED"
	OrderStatusREFUNDED string = "REFUNDED"

	// OrderStatusREFUNDFAILED captures enum value "REFUND_FAILED"
	OrderStatusREFUNDFAILED string = "REFUND_FAILED"

	// OrderStatusCLOSED captures enum value "CLOSED"
	OrderStatusCLOSED string = "CLOSED"

	// OrderStatusDELETED captures enum value "DELETED"
	OrderStatusDELETED string = "DELETED"
)

// prop value enum
func (m *Order) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, orderTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Order) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

func (m *Order) validateUpdatedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.UpdatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("updatedAt", "body", "date-time", m.UpdatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Order) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Order) UnmarshalBinary(b []byte) error {
	var res Order
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
