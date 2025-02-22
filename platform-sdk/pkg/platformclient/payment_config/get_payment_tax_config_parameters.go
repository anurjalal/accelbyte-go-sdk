// Code generated by go-swagger; DO NOT EDIT.

package payment_config

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

// NewGetPaymentTaxConfigParams creates a new GetPaymentTaxConfigParams object
// with the default values initialized.
func NewGetPaymentTaxConfigParams() *GetPaymentTaxConfigParams {

	return &GetPaymentTaxConfigParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetPaymentTaxConfigParamsWithTimeout creates a new GetPaymentTaxConfigParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetPaymentTaxConfigParamsWithTimeout(timeout time.Duration) *GetPaymentTaxConfigParams {

	return &GetPaymentTaxConfigParams{

		timeout: timeout,
	}
}

// NewGetPaymentTaxConfigParamsWithContext creates a new GetPaymentTaxConfigParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetPaymentTaxConfigParamsWithContext(ctx context.Context) *GetPaymentTaxConfigParams {

	return &GetPaymentTaxConfigParams{

		Context: ctx,
	}
}

// NewGetPaymentTaxConfigParamsWithHTTPClient creates a new GetPaymentTaxConfigParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetPaymentTaxConfigParamsWithHTTPClient(client *http.Client) *GetPaymentTaxConfigParams {

	return &GetPaymentTaxConfigParams{
		HTTPClient: client,
	}
}

/*GetPaymentTaxConfigParams contains all the parameters to send to the API endpoint
for the get payment tax config operation typically these are written to a http.Request
*/
type GetPaymentTaxConfigParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get payment tax config params
func (o *GetPaymentTaxConfigParams) WithTimeout(timeout time.Duration) *GetPaymentTaxConfigParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get payment tax config params
func (o *GetPaymentTaxConfigParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get payment tax config params
func (o *GetPaymentTaxConfigParams) WithContext(ctx context.Context) *GetPaymentTaxConfigParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get payment tax config params
func (o *GetPaymentTaxConfigParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get payment tax config params
func (o *GetPaymentTaxConfigParams) WithHTTPClient(client *http.Client) *GetPaymentTaxConfigParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get payment tax config params
func (o *GetPaymentTaxConfigParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetPaymentTaxConfigParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
