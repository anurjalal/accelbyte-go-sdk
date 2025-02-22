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

// NewResetUserStatItemValue1Params creates a new ResetUserStatItemValue1Params object
// with the default values initialized.
func NewResetUserStatItemValue1Params() *ResetUserStatItemValue1Params {
	var ()
	return &ResetUserStatItemValue1Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewResetUserStatItemValue1ParamsWithTimeout creates a new ResetUserStatItemValue1Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewResetUserStatItemValue1ParamsWithTimeout(timeout time.Duration) *ResetUserStatItemValue1Params {
	var ()
	return &ResetUserStatItemValue1Params{

		timeout: timeout,
	}
}

// NewResetUserStatItemValue1ParamsWithContext creates a new ResetUserStatItemValue1Params object
// with the default values initialized, and the ability to set a context for a request
func NewResetUserStatItemValue1ParamsWithContext(ctx context.Context) *ResetUserStatItemValue1Params {
	var ()
	return &ResetUserStatItemValue1Params{

		Context: ctx,
	}
}

// NewResetUserStatItemValue1ParamsWithHTTPClient creates a new ResetUserStatItemValue1Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewResetUserStatItemValue1ParamsWithHTTPClient(client *http.Client) *ResetUserStatItemValue1Params {
	var ()
	return &ResetUserStatItemValue1Params{
		HTTPClient: client,
	}
}

/*ResetUserStatItemValue1Params contains all the parameters to send to the API endpoint
for the reset user stat item value 1 operation typically these are written to a http.Request
*/
type ResetUserStatItemValue1Params struct {

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

// WithTimeout adds the timeout to the reset user stat item value 1 params
func (o *ResetUserStatItemValue1Params) WithTimeout(timeout time.Duration) *ResetUserStatItemValue1Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the reset user stat item value 1 params
func (o *ResetUserStatItemValue1Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the reset user stat item value 1 params
func (o *ResetUserStatItemValue1Params) WithContext(ctx context.Context) *ResetUserStatItemValue1Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the reset user stat item value 1 params
func (o *ResetUserStatItemValue1Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the reset user stat item value 1 params
func (o *ResetUserStatItemValue1Params) WithHTTPClient(client *http.Client) *ResetUserStatItemValue1Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the reset user stat item value 1 params
func (o *ResetUserStatItemValue1Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNamespace adds the namespace to the reset user stat item value 1 params
func (o *ResetUserStatItemValue1Params) WithNamespace(namespace string) *ResetUserStatItemValue1Params {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the reset user stat item value 1 params
func (o *ResetUserStatItemValue1Params) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithStatCode adds the statCode to the reset user stat item value 1 params
func (o *ResetUserStatItemValue1Params) WithStatCode(statCode string) *ResetUserStatItemValue1Params {
	o.SetStatCode(statCode)
	return o
}

// SetStatCode adds the statCode to the reset user stat item value 1 params
func (o *ResetUserStatItemValue1Params) SetStatCode(statCode string) {
	o.StatCode = statCode
}

// WithUserID adds the userID to the reset user stat item value 1 params
func (o *ResetUserStatItemValue1Params) WithUserID(userID string) *ResetUserStatItemValue1Params {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the reset user stat item value 1 params
func (o *ResetUserStatItemValue1Params) SetUserID(userID string) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *ResetUserStatItemValue1Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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
