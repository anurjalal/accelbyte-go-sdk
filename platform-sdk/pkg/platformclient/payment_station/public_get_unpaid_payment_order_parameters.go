// Code generated by go-swagger; DO NOT EDIT.

package payment_station

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewPublicGetUnpaidPaymentOrderParams creates a new PublicGetUnpaidPaymentOrderParams object
// with the default values initialized.
func NewPublicGetUnpaidPaymentOrderParams() *PublicGetUnpaidPaymentOrderParams {
	var ()
	return &PublicGetUnpaidPaymentOrderParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPublicGetUnpaidPaymentOrderParamsWithTimeout creates a new PublicGetUnpaidPaymentOrderParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPublicGetUnpaidPaymentOrderParamsWithTimeout(timeout time.Duration) *PublicGetUnpaidPaymentOrderParams {
	var ()
	return &PublicGetUnpaidPaymentOrderParams{

		timeout: timeout,
	}
}

// NewPublicGetUnpaidPaymentOrderParamsWithContext creates a new PublicGetUnpaidPaymentOrderParams object
// with the default values initialized, and the ability to set a context for a request
func NewPublicGetUnpaidPaymentOrderParamsWithContext(ctx context.Context) *PublicGetUnpaidPaymentOrderParams {
	var ()
	return &PublicGetUnpaidPaymentOrderParams{

		Context: ctx,
	}
}

// NewPublicGetUnpaidPaymentOrderParamsWithHTTPClient creates a new PublicGetUnpaidPaymentOrderParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPublicGetUnpaidPaymentOrderParamsWithHTTPClient(client *http.Client) *PublicGetUnpaidPaymentOrderParams {
	var ()
	return &PublicGetUnpaidPaymentOrderParams{
		HTTPClient: client,
	}
}

/*PublicGetUnpaidPaymentOrderParams contains all the parameters to send to the API endpoint
for the public get unpaid payment order operation typically these are written to a http.Request
*/
type PublicGetUnpaidPaymentOrderParams struct {

	/*Namespace*/
	Namespace string
	/*PaymentOrderNo*/
	PaymentOrderNo string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the public get unpaid payment order params
func (o *PublicGetUnpaidPaymentOrderParams) WithTimeout(timeout time.Duration) *PublicGetUnpaidPaymentOrderParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the public get unpaid payment order params
func (o *PublicGetUnpaidPaymentOrderParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the public get unpaid payment order params
func (o *PublicGetUnpaidPaymentOrderParams) WithContext(ctx context.Context) *PublicGetUnpaidPaymentOrderParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the public get unpaid payment order params
func (o *PublicGetUnpaidPaymentOrderParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the public get unpaid payment order params
func (o *PublicGetUnpaidPaymentOrderParams) WithHTTPClient(client *http.Client) *PublicGetUnpaidPaymentOrderParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the public get unpaid payment order params
func (o *PublicGetUnpaidPaymentOrderParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNamespace adds the namespace to the public get unpaid payment order params
func (o *PublicGetUnpaidPaymentOrderParams) WithNamespace(namespace string) *PublicGetUnpaidPaymentOrderParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the public get unpaid payment order params
func (o *PublicGetUnpaidPaymentOrderParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithPaymentOrderNo adds the paymentOrderNo to the public get unpaid payment order params
func (o *PublicGetUnpaidPaymentOrderParams) WithPaymentOrderNo(paymentOrderNo string) *PublicGetUnpaidPaymentOrderParams {
	o.SetPaymentOrderNo(paymentOrderNo)
	return o
}

// SetPaymentOrderNo adds the paymentOrderNo to the public get unpaid payment order params
func (o *PublicGetUnpaidPaymentOrderParams) SetPaymentOrderNo(paymentOrderNo string) {
	o.PaymentOrderNo = paymentOrderNo
}

// WriteToRequest writes these params to a swagger request
func (o *PublicGetUnpaidPaymentOrderParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param namespace
	if err := r.SetPathParam("namespace", o.Namespace); err != nil {
		return err
	}

	// path param paymentOrderNo
	if err := r.SetPathParam("paymentOrderNo", o.PaymentOrderNo); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
