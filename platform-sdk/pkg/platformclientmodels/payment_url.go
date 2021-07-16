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

// PaymentURL payment Url
//
// swagger:model PaymentUrl
type PaymentURL struct {

	// payment provider
	// Required: true
	// Enum: [WALLET XSOLLA ADYEN STRIPE CHECKOUT ALIPAY WXPAY PAYPAL]
	PaymentProvider *string `json:"paymentProvider"`

	// payment type
	// Required: true
	// Enum: [QR_CODE LINK]
	PaymentType *string `json:"paymentType"`

	// payment url
	// Required: true
	PaymentURL *string `json:"paymentUrl"`

	// return url
	ReturnURL string `json:"returnUrl,omitempty"`
}

// Validate validates this payment Url
func (m *PaymentURL) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePaymentProvider(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePaymentType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePaymentURL(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var paymentUrlTypePaymentProviderPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["WALLET","XSOLLA","ADYEN","STRIPE","CHECKOUT","ALIPAY","WXPAY","PAYPAL"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		paymentUrlTypePaymentProviderPropEnum = append(paymentUrlTypePaymentProviderPropEnum, v)
	}
}

const (

	// PaymentURLPaymentProviderWALLET captures enum value "WALLET"
	PaymentURLPaymentProviderWALLET string = "WALLET"

	// PaymentURLPaymentProviderXSOLLA captures enum value "XSOLLA"
	PaymentURLPaymentProviderXSOLLA string = "XSOLLA"

	// PaymentURLPaymentProviderADYEN captures enum value "ADYEN"
	PaymentURLPaymentProviderADYEN string = "ADYEN"

	// PaymentURLPaymentProviderSTRIPE captures enum value "STRIPE"
	PaymentURLPaymentProviderSTRIPE string = "STRIPE"

	// PaymentURLPaymentProviderCHECKOUT captures enum value "CHECKOUT"
	PaymentURLPaymentProviderCHECKOUT string = "CHECKOUT"

	// PaymentURLPaymentProviderALIPAY captures enum value "ALIPAY"
	PaymentURLPaymentProviderALIPAY string = "ALIPAY"

	// PaymentURLPaymentProviderWXPAY captures enum value "WXPAY"
	PaymentURLPaymentProviderWXPAY string = "WXPAY"

	// PaymentURLPaymentProviderPAYPAL captures enum value "PAYPAL"
	PaymentURLPaymentProviderPAYPAL string = "PAYPAL"
)

// prop value enum
func (m *PaymentURL) validatePaymentProviderEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, paymentUrlTypePaymentProviderPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *PaymentURL) validatePaymentProvider(formats strfmt.Registry) error {

	if err := validate.Required("paymentProvider", "body", m.PaymentProvider); err != nil {
		return err
	}

	// value enum
	if err := m.validatePaymentProviderEnum("paymentProvider", "body", *m.PaymentProvider); err != nil {
		return err
	}

	return nil
}

var paymentUrlTypePaymentTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["QR_CODE","LINK"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		paymentUrlTypePaymentTypePropEnum = append(paymentUrlTypePaymentTypePropEnum, v)
	}
}

const (

	// PaymentURLPaymentTypeQRCODE captures enum value "QR_CODE"
	PaymentURLPaymentTypeQRCODE string = "QR_CODE"

	// PaymentURLPaymentTypeLINK captures enum value "LINK"
	PaymentURLPaymentTypeLINK string = "LINK"
)

// prop value enum
func (m *PaymentURL) validatePaymentTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, paymentUrlTypePaymentTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *PaymentURL) validatePaymentType(formats strfmt.Registry) error {

	if err := validate.Required("paymentType", "body", m.PaymentType); err != nil {
		return err
	}

	// value enum
	if err := m.validatePaymentTypeEnum("paymentType", "body", *m.PaymentType); err != nil {
		return err
	}

	return nil
}

func (m *PaymentURL) validatePaymentURL(formats strfmt.Registry) error {

	if err := validate.Required("paymentUrl", "body", m.PaymentURL); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PaymentURL) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PaymentURL) UnmarshalBinary(b []byte) error {
	var res PaymentURL
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
