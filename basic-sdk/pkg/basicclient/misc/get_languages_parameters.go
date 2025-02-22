// Code generated by go-swagger; DO NOT EDIT.

package misc

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

// NewGetLanguagesParams creates a new GetLanguagesParams object
// with the default values initialized.
func NewGetLanguagesParams() *GetLanguagesParams {
	var ()
	return &GetLanguagesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetLanguagesParamsWithTimeout creates a new GetLanguagesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetLanguagesParamsWithTimeout(timeout time.Duration) *GetLanguagesParams {
	var ()
	return &GetLanguagesParams{

		timeout: timeout,
	}
}

// NewGetLanguagesParamsWithContext creates a new GetLanguagesParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetLanguagesParamsWithContext(ctx context.Context) *GetLanguagesParams {
	var ()
	return &GetLanguagesParams{

		Context: ctx,
	}
}

// NewGetLanguagesParamsWithHTTPClient creates a new GetLanguagesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetLanguagesParamsWithHTTPClient(client *http.Client) *GetLanguagesParams {
	var ()
	return &GetLanguagesParams{
		HTTPClient: client,
	}
}

/*GetLanguagesParams contains all the parameters to send to the API endpoint
for the get languages operation typically these are written to a http.Request
*/
type GetLanguagesParams struct {

	/*Namespace
	  namespace, only accept alphabet and numeric

	*/
	Namespace string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get languages params
func (o *GetLanguagesParams) WithTimeout(timeout time.Duration) *GetLanguagesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get languages params
func (o *GetLanguagesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get languages params
func (o *GetLanguagesParams) WithContext(ctx context.Context) *GetLanguagesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get languages params
func (o *GetLanguagesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get languages params
func (o *GetLanguagesParams) WithHTTPClient(client *http.Client) *GetLanguagesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get languages params
func (o *GetLanguagesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNamespace adds the namespace to the get languages params
func (o *GetLanguagesParams) WithNamespace(namespace string) *GetLanguagesParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the get languages params
func (o *GetLanguagesParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WriteToRequest writes these params to a swagger request
func (o *GetLanguagesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
