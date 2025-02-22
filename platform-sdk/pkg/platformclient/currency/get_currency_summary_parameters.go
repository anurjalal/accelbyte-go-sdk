// Code generated by go-swagger; DO NOT EDIT.

package currency

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

// NewGetCurrencySummaryParams creates a new GetCurrencySummaryParams object
// with the default values initialized.
func NewGetCurrencySummaryParams() *GetCurrencySummaryParams {
	var ()
	return &GetCurrencySummaryParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetCurrencySummaryParamsWithTimeout creates a new GetCurrencySummaryParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetCurrencySummaryParamsWithTimeout(timeout time.Duration) *GetCurrencySummaryParams {
	var ()
	return &GetCurrencySummaryParams{

		timeout: timeout,
	}
}

// NewGetCurrencySummaryParamsWithContext creates a new GetCurrencySummaryParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetCurrencySummaryParamsWithContext(ctx context.Context) *GetCurrencySummaryParams {
	var ()
	return &GetCurrencySummaryParams{

		Context: ctx,
	}
}

// NewGetCurrencySummaryParamsWithHTTPClient creates a new GetCurrencySummaryParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetCurrencySummaryParamsWithHTTPClient(client *http.Client) *GetCurrencySummaryParams {
	var ()
	return &GetCurrencySummaryParams{
		HTTPClient: client,
	}
}

/*GetCurrencySummaryParams contains all the parameters to send to the API endpoint
for the get currency summary operation typically these are written to a http.Request
*/
type GetCurrencySummaryParams struct {

	/*CurrencyCode
	  currencyCode

	*/
	CurrencyCode string
	/*Namespace
	  Namespace

	*/
	Namespace string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get currency summary params
func (o *GetCurrencySummaryParams) WithTimeout(timeout time.Duration) *GetCurrencySummaryParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get currency summary params
func (o *GetCurrencySummaryParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get currency summary params
func (o *GetCurrencySummaryParams) WithContext(ctx context.Context) *GetCurrencySummaryParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get currency summary params
func (o *GetCurrencySummaryParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get currency summary params
func (o *GetCurrencySummaryParams) WithHTTPClient(client *http.Client) *GetCurrencySummaryParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get currency summary params
func (o *GetCurrencySummaryParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCurrencyCode adds the currencyCode to the get currency summary params
func (o *GetCurrencySummaryParams) WithCurrencyCode(currencyCode string) *GetCurrencySummaryParams {
	o.SetCurrencyCode(currencyCode)
	return o
}

// SetCurrencyCode adds the currencyCode to the get currency summary params
func (o *GetCurrencySummaryParams) SetCurrencyCode(currencyCode string) {
	o.CurrencyCode = currencyCode
}

// WithNamespace adds the namespace to the get currency summary params
func (o *GetCurrencySummaryParams) WithNamespace(namespace string) *GetCurrencySummaryParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the get currency summary params
func (o *GetCurrencySummaryParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WriteToRequest writes these params to a swagger request
func (o *GetCurrencySummaryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param currencyCode
	if err := r.SetPathParam("currencyCode", o.CurrencyCode); err != nil {
		return err
	}

	// path param namespace
	if err := r.SetPathParam("namespace", o.Namespace); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
