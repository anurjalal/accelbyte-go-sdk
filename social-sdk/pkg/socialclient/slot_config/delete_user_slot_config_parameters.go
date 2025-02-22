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

// NewDeleteUserSlotConfigParams creates a new DeleteUserSlotConfigParams object
// with the default values initialized.
func NewDeleteUserSlotConfigParams() *DeleteUserSlotConfigParams {
	var ()
	return &DeleteUserSlotConfigParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteUserSlotConfigParamsWithTimeout creates a new DeleteUserSlotConfigParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteUserSlotConfigParamsWithTimeout(timeout time.Duration) *DeleteUserSlotConfigParams {
	var ()
	return &DeleteUserSlotConfigParams{

		timeout: timeout,
	}
}

// NewDeleteUserSlotConfigParamsWithContext creates a new DeleteUserSlotConfigParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteUserSlotConfigParamsWithContext(ctx context.Context) *DeleteUserSlotConfigParams {
	var ()
	return &DeleteUserSlotConfigParams{

		Context: ctx,
	}
}

// NewDeleteUserSlotConfigParamsWithHTTPClient creates a new DeleteUserSlotConfigParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteUserSlotConfigParamsWithHTTPClient(client *http.Client) *DeleteUserSlotConfigParams {
	var ()
	return &DeleteUserSlotConfigParams{
		HTTPClient: client,
	}
}

/*DeleteUserSlotConfigParams contains all the parameters to send to the API endpoint
for the delete user slot config operation typically these are written to a http.Request
*/
type DeleteUserSlotConfigParams struct {

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

// WithTimeout adds the timeout to the delete user slot config params
func (o *DeleteUserSlotConfigParams) WithTimeout(timeout time.Duration) *DeleteUserSlotConfigParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete user slot config params
func (o *DeleteUserSlotConfigParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete user slot config params
func (o *DeleteUserSlotConfigParams) WithContext(ctx context.Context) *DeleteUserSlotConfigParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete user slot config params
func (o *DeleteUserSlotConfigParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete user slot config params
func (o *DeleteUserSlotConfigParams) WithHTTPClient(client *http.Client) *DeleteUserSlotConfigParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete user slot config params
func (o *DeleteUserSlotConfigParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNamespace adds the namespace to the delete user slot config params
func (o *DeleteUserSlotConfigParams) WithNamespace(namespace string) *DeleteUserSlotConfigParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the delete user slot config params
func (o *DeleteUserSlotConfigParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithUserID adds the userID to the delete user slot config params
func (o *DeleteUserSlotConfigParams) WithUserID(userID string) *DeleteUserSlotConfigParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the delete user slot config params
func (o *DeleteUserSlotConfigParams) SetUserID(userID string) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteUserSlotConfigParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
