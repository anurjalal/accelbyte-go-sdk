// Code generated by go-swagger; DO NOT EDIT.

package entitlement

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

// NewPublicDeleteUserDistributionReceiverParams creates a new PublicDeleteUserDistributionReceiverParams object
// with the default values initialized.
func NewPublicDeleteUserDistributionReceiverParams() *PublicDeleteUserDistributionReceiverParams {
	var ()
	return &PublicDeleteUserDistributionReceiverParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPublicDeleteUserDistributionReceiverParamsWithTimeout creates a new PublicDeleteUserDistributionReceiverParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPublicDeleteUserDistributionReceiverParamsWithTimeout(timeout time.Duration) *PublicDeleteUserDistributionReceiverParams {
	var ()
	return &PublicDeleteUserDistributionReceiverParams{

		timeout: timeout,
	}
}

// NewPublicDeleteUserDistributionReceiverParamsWithContext creates a new PublicDeleteUserDistributionReceiverParams object
// with the default values initialized, and the ability to set a context for a request
func NewPublicDeleteUserDistributionReceiverParamsWithContext(ctx context.Context) *PublicDeleteUserDistributionReceiverParams {
	var ()
	return &PublicDeleteUserDistributionReceiverParams{

		Context: ctx,
	}
}

// NewPublicDeleteUserDistributionReceiverParamsWithHTTPClient creates a new PublicDeleteUserDistributionReceiverParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPublicDeleteUserDistributionReceiverParamsWithHTTPClient(client *http.Client) *PublicDeleteUserDistributionReceiverParams {
	var ()
	return &PublicDeleteUserDistributionReceiverParams{
		HTTPClient: client,
	}
}

/*PublicDeleteUserDistributionReceiverParams contains all the parameters to send to the API endpoint
for the public delete user distribution receiver operation typically these are written to a http.Request
*/
type PublicDeleteUserDistributionReceiverParams struct {

	/*ExtUserID*/
	ExtUserID string
	/*Namespace
	  game namespace

	*/
	Namespace string
	/*UserID
	  game user id

	*/
	UserID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the public delete user distribution receiver params
func (o *PublicDeleteUserDistributionReceiverParams) WithTimeout(timeout time.Duration) *PublicDeleteUserDistributionReceiverParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the public delete user distribution receiver params
func (o *PublicDeleteUserDistributionReceiverParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the public delete user distribution receiver params
func (o *PublicDeleteUserDistributionReceiverParams) WithContext(ctx context.Context) *PublicDeleteUserDistributionReceiverParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the public delete user distribution receiver params
func (o *PublicDeleteUserDistributionReceiverParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the public delete user distribution receiver params
func (o *PublicDeleteUserDistributionReceiverParams) WithHTTPClient(client *http.Client) *PublicDeleteUserDistributionReceiverParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the public delete user distribution receiver params
func (o *PublicDeleteUserDistributionReceiverParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithExtUserID adds the extUserID to the public delete user distribution receiver params
func (o *PublicDeleteUserDistributionReceiverParams) WithExtUserID(extUserID string) *PublicDeleteUserDistributionReceiverParams {
	o.SetExtUserID(extUserID)
	return o
}

// SetExtUserID adds the extUserId to the public delete user distribution receiver params
func (o *PublicDeleteUserDistributionReceiverParams) SetExtUserID(extUserID string) {
	o.ExtUserID = extUserID
}

// WithNamespace adds the namespace to the public delete user distribution receiver params
func (o *PublicDeleteUserDistributionReceiverParams) WithNamespace(namespace string) *PublicDeleteUserDistributionReceiverParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the public delete user distribution receiver params
func (o *PublicDeleteUserDistributionReceiverParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithUserID adds the userID to the public delete user distribution receiver params
func (o *PublicDeleteUserDistributionReceiverParams) WithUserID(userID string) *PublicDeleteUserDistributionReceiverParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the public delete user distribution receiver params
func (o *PublicDeleteUserDistributionReceiverParams) SetUserID(userID string) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *PublicDeleteUserDistributionReceiverParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param extUserId
	if err := r.SetPathParam("extUserId", o.ExtUserID); err != nil {
		return err
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
