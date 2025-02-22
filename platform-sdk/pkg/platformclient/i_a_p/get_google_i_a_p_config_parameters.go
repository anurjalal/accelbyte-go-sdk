// Code generated by go-swagger; DO NOT EDIT.

package i_a_p

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

// NewGetGoogleIAPConfigParams creates a new GetGoogleIAPConfigParams object
// with the default values initialized.
func NewGetGoogleIAPConfigParams() *GetGoogleIAPConfigParams {
	var ()
	return &GetGoogleIAPConfigParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetGoogleIAPConfigParamsWithTimeout creates a new GetGoogleIAPConfigParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetGoogleIAPConfigParamsWithTimeout(timeout time.Duration) *GetGoogleIAPConfigParams {
	var ()
	return &GetGoogleIAPConfigParams{

		timeout: timeout,
	}
}

// NewGetGoogleIAPConfigParamsWithContext creates a new GetGoogleIAPConfigParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetGoogleIAPConfigParamsWithContext(ctx context.Context) *GetGoogleIAPConfigParams {
	var ()
	return &GetGoogleIAPConfigParams{

		Context: ctx,
	}
}

// NewGetGoogleIAPConfigParamsWithHTTPClient creates a new GetGoogleIAPConfigParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetGoogleIAPConfigParamsWithHTTPClient(client *http.Client) *GetGoogleIAPConfigParams {
	var ()
	return &GetGoogleIAPConfigParams{
		HTTPClient: client,
	}
}

/*GetGoogleIAPConfigParams contains all the parameters to send to the API endpoint
for the get google i a p config operation typically these are written to a http.Request
*/
type GetGoogleIAPConfigParams struct {

	/*Namespace*/
	Namespace string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get google i a p config params
func (o *GetGoogleIAPConfigParams) WithTimeout(timeout time.Duration) *GetGoogleIAPConfigParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get google i a p config params
func (o *GetGoogleIAPConfigParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get google i a p config params
func (o *GetGoogleIAPConfigParams) WithContext(ctx context.Context) *GetGoogleIAPConfigParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get google i a p config params
func (o *GetGoogleIAPConfigParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get google i a p config params
func (o *GetGoogleIAPConfigParams) WithHTTPClient(client *http.Client) *GetGoogleIAPConfigParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get google i a p config params
func (o *GetGoogleIAPConfigParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNamespace adds the namespace to the get google i a p config params
func (o *GetGoogleIAPConfigParams) WithNamespace(namespace string) *GetGoogleIAPConfigParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the get google i a p config params
func (o *GetGoogleIAPConfigParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WriteToRequest writes these params to a swagger request
func (o *GetGoogleIAPConfigParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
