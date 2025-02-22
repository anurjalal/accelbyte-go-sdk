// Code generated by go-swagger; DO NOT EDIT.

package anonymization

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

// NewAnonymizeEntitlementParams creates a new AnonymizeEntitlementParams object
// with the default values initialized.
func NewAnonymizeEntitlementParams() *AnonymizeEntitlementParams {
	var ()
	return &AnonymizeEntitlementParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAnonymizeEntitlementParamsWithTimeout creates a new AnonymizeEntitlementParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAnonymizeEntitlementParamsWithTimeout(timeout time.Duration) *AnonymizeEntitlementParams {
	var ()
	return &AnonymizeEntitlementParams{

		timeout: timeout,
	}
}

// NewAnonymizeEntitlementParamsWithContext creates a new AnonymizeEntitlementParams object
// with the default values initialized, and the ability to set a context for a request
func NewAnonymizeEntitlementParamsWithContext(ctx context.Context) *AnonymizeEntitlementParams {
	var ()
	return &AnonymizeEntitlementParams{

		Context: ctx,
	}
}

// NewAnonymizeEntitlementParamsWithHTTPClient creates a new AnonymizeEntitlementParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAnonymizeEntitlementParamsWithHTTPClient(client *http.Client) *AnonymizeEntitlementParams {
	var ()
	return &AnonymizeEntitlementParams{
		HTTPClient: client,
	}
}

/*AnonymizeEntitlementParams contains all the parameters to send to the API endpoint
for the anonymize entitlement operation typically these are written to a http.Request
*/
type AnonymizeEntitlementParams struct {

	/*Namespace*/
	Namespace string
	/*UserID*/
	UserID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the anonymize entitlement params
func (o *AnonymizeEntitlementParams) WithTimeout(timeout time.Duration) *AnonymizeEntitlementParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the anonymize entitlement params
func (o *AnonymizeEntitlementParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the anonymize entitlement params
func (o *AnonymizeEntitlementParams) WithContext(ctx context.Context) *AnonymizeEntitlementParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the anonymize entitlement params
func (o *AnonymizeEntitlementParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the anonymize entitlement params
func (o *AnonymizeEntitlementParams) WithHTTPClient(client *http.Client) *AnonymizeEntitlementParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the anonymize entitlement params
func (o *AnonymizeEntitlementParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNamespace adds the namespace to the anonymize entitlement params
func (o *AnonymizeEntitlementParams) WithNamespace(namespace string) *AnonymizeEntitlementParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the anonymize entitlement params
func (o *AnonymizeEntitlementParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithUserID adds the userID to the anonymize entitlement params
func (o *AnonymizeEntitlementParams) WithUserID(userID string) *AnonymizeEntitlementParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the anonymize entitlement params
func (o *AnonymizeEntitlementParams) SetUserID(userID string) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *AnonymizeEntitlementParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
