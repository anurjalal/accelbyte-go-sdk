// Code generated by go-swagger; DO NOT EDIT.

package public

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

// NewListProvidersParams creates a new ListProvidersParams object
// with the default values initialized.
func NewListProvidersParams() *ListProvidersParams {

	return &ListProvidersParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewListProvidersParamsWithTimeout creates a new ListProvidersParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListProvidersParamsWithTimeout(timeout time.Duration) *ListProvidersParams {

	return &ListProvidersParams{

		timeout: timeout,
	}
}

// NewListProvidersParamsWithContext creates a new ListProvidersParams object
// with the default values initialized, and the ability to set a context for a request
func NewListProvidersParamsWithContext(ctx context.Context) *ListProvidersParams {

	return &ListProvidersParams{

		Context: ctx,
	}
}

// NewListProvidersParamsWithHTTPClient creates a new ListProvidersParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListProvidersParamsWithHTTPClient(client *http.Client) *ListProvidersParams {

	return &ListProvidersParams{
		HTTPClient: client,
	}
}

/*ListProvidersParams contains all the parameters to send to the API endpoint
for the list providers operation typically these are written to a http.Request
*/
type ListProvidersParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list providers params
func (o *ListProvidersParams) WithTimeout(timeout time.Duration) *ListProvidersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list providers params
func (o *ListProvidersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list providers params
func (o *ListProvidersParams) WithContext(ctx context.Context) *ListProvidersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list providers params
func (o *ListProvidersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list providers params
func (o *ListProvidersParams) WithHTTPClient(client *http.Client) *ListProvidersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list providers params
func (o *ListProvidersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *ListProvidersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
