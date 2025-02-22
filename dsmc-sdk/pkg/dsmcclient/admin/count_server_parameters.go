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

// NewCountServerParams creates a new CountServerParams object
// with the default values initialized.
func NewCountServerParams() *CountServerParams {
	var ()
	return &CountServerParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCountServerParamsWithTimeout creates a new CountServerParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCountServerParamsWithTimeout(timeout time.Duration) *CountServerParams {
	var ()
	return &CountServerParams{

		timeout: timeout,
	}
}

// NewCountServerParamsWithContext creates a new CountServerParams object
// with the default values initialized, and the ability to set a context for a request
func NewCountServerParamsWithContext(ctx context.Context) *CountServerParams {
	var ()
	return &CountServerParams{

		Context: ctx,
	}
}

// NewCountServerParamsWithHTTPClient creates a new CountServerParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCountServerParamsWithHTTPClient(client *http.Client) *CountServerParams {
	var ()
	return &CountServerParams{
		HTTPClient: client,
	}
}

/*CountServerParams contains all the parameters to send to the API endpoint
for the count server operation typically these are written to a http.Request
*/
type CountServerParams struct {

	/*Namespace
	  namespace of the game

	*/
	Namespace string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the count server params
func (o *CountServerParams) WithTimeout(timeout time.Duration) *CountServerParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the count server params
func (o *CountServerParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the count server params
func (o *CountServerParams) WithContext(ctx context.Context) *CountServerParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the count server params
func (o *CountServerParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the count server params
func (o *CountServerParams) WithHTTPClient(client *http.Client) *CountServerParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the count server params
func (o *CountServerParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNamespace adds the namespace to the count server params
func (o *CountServerParams) WithNamespace(namespace string) *CountServerParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the count server params
func (o *CountServerParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WriteToRequest writes these params to a swagger request
func (o *CountServerParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param namespace
	if err := r.SetPathParam("namespace", o.Namespace); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
