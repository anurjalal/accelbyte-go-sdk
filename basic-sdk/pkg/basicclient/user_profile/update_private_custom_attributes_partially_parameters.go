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

// NewUpdatePrivateCustomAttributesPartiallyParams creates a new UpdatePrivateCustomAttributesPartiallyParams object
// with the default values initialized.
func NewUpdatePrivateCustomAttributesPartiallyParams() *UpdatePrivateCustomAttributesPartiallyParams {
	var ()
	return &UpdatePrivateCustomAttributesPartiallyParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdatePrivateCustomAttributesPartiallyParamsWithTimeout creates a new UpdatePrivateCustomAttributesPartiallyParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdatePrivateCustomAttributesPartiallyParamsWithTimeout(timeout time.Duration) *UpdatePrivateCustomAttributesPartiallyParams {
	var ()
	return &UpdatePrivateCustomAttributesPartiallyParams{

		timeout: timeout,
	}
}

// NewUpdatePrivateCustomAttributesPartiallyParamsWithContext creates a new UpdatePrivateCustomAttributesPartiallyParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdatePrivateCustomAttributesPartiallyParamsWithContext(ctx context.Context) *UpdatePrivateCustomAttributesPartiallyParams {
	var ()
	return &UpdatePrivateCustomAttributesPartiallyParams{

		Context: ctx,
	}
}

// NewUpdatePrivateCustomAttributesPartiallyParamsWithHTTPClient creates a new UpdatePrivateCustomAttributesPartiallyParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdatePrivateCustomAttributesPartiallyParamsWithHTTPClient(client *http.Client) *UpdatePrivateCustomAttributesPartiallyParams {
	var ()
	return &UpdatePrivateCustomAttributesPartiallyParams{
		HTTPClient: client,
	}
}

/*UpdatePrivateCustomAttributesPartiallyParams contains all the parameters to send to the API endpoint
for the update private custom attributes partially operation typically these are written to a http.Request
*/
type UpdatePrivateCustomAttributesPartiallyParams struct {

	/*Body*/
	Body map[string]interface{}
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

// WithTimeout adds the timeout to the update private custom attributes partially params
func (o *UpdatePrivateCustomAttributesPartiallyParams) WithTimeout(timeout time.Duration) *UpdatePrivateCustomAttributesPartiallyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update private custom attributes partially params
func (o *UpdatePrivateCustomAttributesPartiallyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update private custom attributes partially params
func (o *UpdatePrivateCustomAttributesPartiallyParams) WithContext(ctx context.Context) *UpdatePrivateCustomAttributesPartiallyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update private custom attributes partially params
func (o *UpdatePrivateCustomAttributesPartiallyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update private custom attributes partially params
func (o *UpdatePrivateCustomAttributesPartiallyParams) WithHTTPClient(client *http.Client) *UpdatePrivateCustomAttributesPartiallyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update private custom attributes partially params
func (o *UpdatePrivateCustomAttributesPartiallyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update private custom attributes partially params
func (o *UpdatePrivateCustomAttributesPartiallyParams) WithBody(body map[string]interface{}) *UpdatePrivateCustomAttributesPartiallyParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update private custom attributes partially params
func (o *UpdatePrivateCustomAttributesPartiallyParams) SetBody(body map[string]interface{}) {
	o.Body = body
}

// WithNamespace adds the namespace to the update private custom attributes partially params
func (o *UpdatePrivateCustomAttributesPartiallyParams) WithNamespace(namespace string) *UpdatePrivateCustomAttributesPartiallyParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the update private custom attributes partially params
func (o *UpdatePrivateCustomAttributesPartiallyParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithUserID adds the userID to the update private custom attributes partially params
func (o *UpdatePrivateCustomAttributesPartiallyParams) WithUserID(userID string) *UpdatePrivateCustomAttributesPartiallyParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the update private custom attributes partially params
func (o *UpdatePrivateCustomAttributesPartiallyParams) SetUserID(userID string) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *UpdatePrivateCustomAttributesPartiallyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

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
