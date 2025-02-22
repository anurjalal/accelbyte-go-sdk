// Code generated by go-swagger; DO NOT EDIT.

package user_profile

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

// NewGetCustomAttributesInfo1Params creates a new GetCustomAttributesInfo1Params object
// with the default values initialized.
func NewGetCustomAttributesInfo1Params() *GetCustomAttributesInfo1Params {
	var ()
	return &GetCustomAttributesInfo1Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetCustomAttributesInfo1ParamsWithTimeout creates a new GetCustomAttributesInfo1Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetCustomAttributesInfo1ParamsWithTimeout(timeout time.Duration) *GetCustomAttributesInfo1Params {
	var ()
	return &GetCustomAttributesInfo1Params{

		timeout: timeout,
	}
}

// NewGetCustomAttributesInfo1ParamsWithContext creates a new GetCustomAttributesInfo1Params object
// with the default values initialized, and the ability to set a context for a request
func NewGetCustomAttributesInfo1ParamsWithContext(ctx context.Context) *GetCustomAttributesInfo1Params {
	var ()
	return &GetCustomAttributesInfo1Params{

		Context: ctx,
	}
}

// NewGetCustomAttributesInfo1ParamsWithHTTPClient creates a new GetCustomAttributesInfo1Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetCustomAttributesInfo1ParamsWithHTTPClient(client *http.Client) *GetCustomAttributesInfo1Params {
	var ()
	return &GetCustomAttributesInfo1Params{
		HTTPClient: client,
	}
}

/*GetCustomAttributesInfo1Params contains all the parameters to send to the API endpoint
for the get custom attributes info 1 operation typically these are written to a http.Request
*/
type GetCustomAttributesInfo1Params struct {

	/*Namespace
	  namespace, only accept alphabet and numeric

	*/
	Namespace string
	/*UserID
	  user's id, should follow UUID version 4 without hyphen

	*/
	UserID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get custom attributes info 1 params
func (o *GetCustomAttributesInfo1Params) WithTimeout(timeout time.Duration) *GetCustomAttributesInfo1Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get custom attributes info 1 params
func (o *GetCustomAttributesInfo1Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get custom attributes info 1 params
func (o *GetCustomAttributesInfo1Params) WithContext(ctx context.Context) *GetCustomAttributesInfo1Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get custom attributes info 1 params
func (o *GetCustomAttributesInfo1Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get custom attributes info 1 params
func (o *GetCustomAttributesInfo1Params) WithHTTPClient(client *http.Client) *GetCustomAttributesInfo1Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get custom attributes info 1 params
func (o *GetCustomAttributesInfo1Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNamespace adds the namespace to the get custom attributes info 1 params
func (o *GetCustomAttributesInfo1Params) WithNamespace(namespace string) *GetCustomAttributesInfo1Params {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the get custom attributes info 1 params
func (o *GetCustomAttributesInfo1Params) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithUserID adds the userID to the get custom attributes info 1 params
func (o *GetCustomAttributesInfo1Params) WithUserID(userID string) *GetCustomAttributesInfo1Params {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the get custom attributes info 1 params
func (o *GetCustomAttributesInfo1Params) SetUserID(userID string) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *GetCustomAttributesInfo1Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param namespace
	if err := r.SetPathParam("namespace", o.Namespace); err != nil {
		return err
	}

	// path param userId
	if err := r.SetPathParam("userId", o.UserID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
