// Code generated by go-swagger; DO NOT EDIT.

package admin

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

// NewCountSessionParams creates a new CountSessionParams object
// with the default values initialized.
func NewCountSessionParams() *CountSessionParams {
	var ()
	return &CountSessionParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCountSessionParamsWithTimeout creates a new CountSessionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCountSessionParamsWithTimeout(timeout time.Duration) *CountSessionParams {
	var ()
	return &CountSessionParams{

		timeout: timeout,
	}
}

// NewCountSessionParamsWithContext creates a new CountSessionParams object
// with the default values initialized, and the ability to set a context for a request
func NewCountSessionParamsWithContext(ctx context.Context) *CountSessionParams {
	var ()
	return &CountSessionParams{

		Context: ctx,
	}
}

// NewCountSessionParamsWithHTTPClient creates a new CountSessionParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCountSessionParamsWithHTTPClient(client *http.Client) *CountSessionParams {
	var ()
	return &CountSessionParams{
		HTTPClient: client,
	}
}

/*CountSessionParams contains all the parameters to send to the API endpoint
for the count session operation typically these are written to a http.Request
*/
type CountSessionParams struct {

	/*Namespace
	  namespace of the game

	*/
	Namespace string
	/*Region
	  region where session is located. if not specified it will count all sessions.

	*/
	Region *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the count session params
func (o *CountSessionParams) WithTimeout(timeout time.Duration) *CountSessionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the count session params
func (o *CountSessionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the count session params
func (o *CountSessionParams) WithContext(ctx context.Context) *CountSessionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the count session params
func (o *CountSessionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the count session params
func (o *CountSessionParams) WithHTTPClient(client *http.Client) *CountSessionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the count session params
func (o *CountSessionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNamespace adds the namespace to the count session params
func (o *CountSessionParams) WithNamespace(namespace string) *CountSessionParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the count session params
func (o *CountSessionParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithRegion adds the region to the count session params
func (o *CountSessionParams) WithRegion(region *string) *CountSessionParams {
	o.SetRegion(region)
	return o
}

// SetRegion adds the region to the count session params
func (o *CountSessionParams) SetRegion(region *string) {
	o.Region = region
}

// WriteToRequest writes these params to a swagger request
func (o *CountSessionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param namespace
	if err := r.SetPathParam("namespace", o.Namespace); err != nil {
		return err
	}

	if o.Region != nil {

		// query param region
		var qrRegion string
		if o.Region != nil {
			qrRegion = *o.Region
		}
		qRegion := qrRegion
		if qRegion != "" {
			if err := r.SetQueryParam("region", qRegion); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
