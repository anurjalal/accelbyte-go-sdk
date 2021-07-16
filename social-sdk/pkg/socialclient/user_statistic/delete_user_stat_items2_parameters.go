// Code generated by go-swagger; DO NOT EDIT.

package user_statistic

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

// NewDeleteUserStatItems2Params creates a new DeleteUserStatItems2Params object
// with the default values initialized.
func NewDeleteUserStatItems2Params() *DeleteUserStatItems2Params {
	var ()
	return &DeleteUserStatItems2Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteUserStatItems2ParamsWithTimeout creates a new DeleteUserStatItems2Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteUserStatItems2ParamsWithTimeout(timeout time.Duration) *DeleteUserStatItems2Params {
	var ()
	return &DeleteUserStatItems2Params{

		timeout: timeout,
	}
}

// NewDeleteUserStatItems2ParamsWithContext creates a new DeleteUserStatItems2Params object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteUserStatItems2ParamsWithContext(ctx context.Context) *DeleteUserStatItems2Params {
	var ()
	return &DeleteUserStatItems2Params{

		Context: ctx,
	}
}

// NewDeleteUserStatItems2ParamsWithHTTPClient creates a new DeleteUserStatItems2Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteUserStatItems2ParamsWithHTTPClient(client *http.Client) *DeleteUserStatItems2Params {
	var ()
	return &DeleteUserStatItems2Params{
		HTTPClient: client,
	}
}

/*DeleteUserStatItems2Params contains all the parameters to send to the API endpoint
for the delete user stat items 2 operation typically these are written to a http.Request
*/
type DeleteUserStatItems2Params struct {

	/*AdditionalKey
	  additional key

	*/
	AdditionalKey *string
	/*Namespace
	  namespace

	*/
	Namespace string
	/*StatCode
	  stat code

	*/
	StatCode string
	/*UserID
	  user ID

	*/
	UserID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete user stat items 2 params
func (o *DeleteUserStatItems2Params) WithTimeout(timeout time.Duration) *DeleteUserStatItems2Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete user stat items 2 params
func (o *DeleteUserStatItems2Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete user stat items 2 params
func (o *DeleteUserStatItems2Params) WithContext(ctx context.Context) *DeleteUserStatItems2Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete user stat items 2 params
func (o *DeleteUserStatItems2Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete user stat items 2 params
func (o *DeleteUserStatItems2Params) WithHTTPClient(client *http.Client) *DeleteUserStatItems2Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete user stat items 2 params
func (o *DeleteUserStatItems2Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAdditionalKey adds the additionalKey to the delete user stat items 2 params
func (o *DeleteUserStatItems2Params) WithAdditionalKey(additionalKey *string) *DeleteUserStatItems2Params {
	o.SetAdditionalKey(additionalKey)
	return o
}

// SetAdditionalKey adds the additionalKey to the delete user stat items 2 params
func (o *DeleteUserStatItems2Params) SetAdditionalKey(additionalKey *string) {
	o.AdditionalKey = additionalKey
}

// WithNamespace adds the namespace to the delete user stat items 2 params
func (o *DeleteUserStatItems2Params) WithNamespace(namespace string) *DeleteUserStatItems2Params {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the delete user stat items 2 params
func (o *DeleteUserStatItems2Params) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithStatCode adds the statCode to the delete user stat items 2 params
func (o *DeleteUserStatItems2Params) WithStatCode(statCode string) *DeleteUserStatItems2Params {
	o.SetStatCode(statCode)
	return o
}

// SetStatCode adds the statCode to the delete user stat items 2 params
func (o *DeleteUserStatItems2Params) SetStatCode(statCode string) {
	o.StatCode = statCode
}

// WithUserID adds the userID to the delete user stat items 2 params
func (o *DeleteUserStatItems2Params) WithUserID(userID string) *DeleteUserStatItems2Params {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the delete user stat items 2 params
func (o *DeleteUserStatItems2Params) SetUserID(userID string) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteUserStatItems2Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.AdditionalKey != nil {

		// query param additionalKey
		var qrAdditionalKey string
		if o.AdditionalKey != nil {
			qrAdditionalKey = *o.AdditionalKey
		}
		qAdditionalKey := qrAdditionalKey
		if qAdditionalKey != "" {
			if err := r.SetQueryParam("additionalKey", qAdditionalKey); err != nil {
				return err
			}
		}

	}

	// path param namespace
	if err := r.SetPathParam("namespace", o.Namespace); err != nil {
		return err
	}

	// path param statCode
	if err := r.SetPathParam("statCode", o.StatCode); err != nil {
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
