// Code generated by go-swagger; DO NOT EDIT.

package slot_config

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

// NewGetUserSlotConfigParams creates a new GetUserSlotConfigParams object
// with the default values initialized.
func NewGetUserSlotConfigParams() *GetUserSlotConfigParams {
	var ()
	return &GetUserSlotConfigParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetUserSlotConfigParamsWithTimeout creates a new GetUserSlotConfigParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetUserSlotConfigParamsWithTimeout(timeout time.Duration) *GetUserSlotConfigParams {
	var ()
	return &GetUserSlotConfigParams{

		timeout: timeout,
	}
}

// NewGetUserSlotConfigParamsWithContext creates a new GetUserSlotConfigParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetUserSlotConfigParamsWithContext(ctx context.Context) *GetUserSlotConfigParams {
	var ()
	return &GetUserSlotConfigParams{

		Context: ctx,
	}
}

// NewGetUserSlotConfigParamsWithHTTPClient creates a new GetUserSlotConfigParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetUserSlotConfigParamsWithHTTPClient(client *http.Client) *GetUserSlotConfigParams {
	var ()
	return &GetUserSlotConfigParams{
		HTTPClient: client,
	}
}

/*GetUserSlotConfigParams contains all the parameters to send to the API endpoint
for the get user slot config operation typically these are written to a http.Request
*/
type GetUserSlotConfigParams struct {

	/*Namespace
	  Namespace ID

	*/
	Namespace string
	/*UserID
	  User ID

	*/
	UserID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get user slot config params
func (o *GetUserSlotConfigParams) WithTimeout(timeout time.Duration) *GetUserSlotConfigParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get user slot config params
func (o *GetUserSlotConfigParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get user slot config params
func (o *GetUserSlotConfigParams) WithContext(ctx context.Context) *GetUserSlotConfigParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get user slot config params
func (o *GetUserSlotConfigParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get user slot config params
func (o *GetUserSlotConfigParams) WithHTTPClient(client *http.Client) *GetUserSlotConfigParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get user slot config params
func (o *GetUserSlotConfigParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNamespace adds the namespace to the get user slot config params
func (o *GetUserSlotConfigParams) WithNamespace(namespace string) *GetUserSlotConfigParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the get user slot config params
func (o *GetUserSlotConfigParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithUserID adds the userID to the get user slot config params
func (o *GetUserSlotConfigParams) WithUserID(userID string) *GetUserSlotConfigParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the get user slot config params
func (o *GetUserSlotConfigParams) SetUserID(userID string) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *GetUserSlotConfigParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
