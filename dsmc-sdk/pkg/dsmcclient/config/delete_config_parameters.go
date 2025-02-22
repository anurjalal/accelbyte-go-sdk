// Code generated by go-swagger; DO NOT EDIT.

package config

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

// NewDeleteConfigParams creates a new DeleteConfigParams object
// with the default values initialized.
func NewDeleteConfigParams() *DeleteConfigParams {
	var ()
	return &DeleteConfigParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteConfigParamsWithTimeout creates a new DeleteConfigParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteConfigParamsWithTimeout(timeout time.Duration) *DeleteConfigParams {
	var ()
	return &DeleteConfigParams{

		timeout: timeout,
	}
}

// NewDeleteConfigParamsWithContext creates a new DeleteConfigParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteConfigParamsWithContext(ctx context.Context) *DeleteConfigParams {
	var ()
	return &DeleteConfigParams{

		Context: ctx,
	}
}

// NewDeleteConfigParamsWithHTTPClient creates a new DeleteConfigParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteConfigParamsWithHTTPClient(client *http.Client) *DeleteConfigParams {
	var ()
	return &DeleteConfigParams{
		HTTPClient: client,
	}
}

/*DeleteConfigParams contains all the parameters to send to the API endpoint
for the delete config operation typically these are written to a http.Request
*/
type DeleteConfigParams struct {

	/*Namespace
	  namespace of the game

	*/
	Namespace string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete config params
func (o *DeleteConfigParams) WithTimeout(timeout time.Duration) *DeleteConfigParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete config params
func (o *DeleteConfigParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete config params
func (o *DeleteConfigParams) WithContext(ctx context.Context) *DeleteConfigParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete config params
func (o *DeleteConfigParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete config params
func (o *DeleteConfigParams) WithHTTPClient(client *http.Client) *DeleteConfigParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete config params
func (o *DeleteConfigParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNamespace adds the namespace to the delete config params
func (o *DeleteConfigParams) WithNamespace(namespace string) *DeleteConfigParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the delete config params
func (o *DeleteConfigParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteConfigParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
