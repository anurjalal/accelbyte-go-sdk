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

	"github.com/AccelByte/accelbyte-go-sdk/platform-sdk/pkg/platformclientmodels"
)

// NewConsumeUserEntitlementParams creates a new ConsumeUserEntitlementParams object
// with the default values initialized.
func NewConsumeUserEntitlementParams() *ConsumeUserEntitlementParams {
	var ()
	return &ConsumeUserEntitlementParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewConsumeUserEntitlementParamsWithTimeout creates a new ConsumeUserEntitlementParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewConsumeUserEntitlementParamsWithTimeout(timeout time.Duration) *ConsumeUserEntitlementParams {
	var ()
	return &ConsumeUserEntitlementParams{

		timeout: timeout,
	}
}

// NewConsumeUserEntitlementParamsWithContext creates a new ConsumeUserEntitlementParams object
// with the default values initialized, and the ability to set a context for a request
func NewConsumeUserEntitlementParamsWithContext(ctx context.Context) *ConsumeUserEntitlementParams {
	var ()
	return &ConsumeUserEntitlementParams{

		Context: ctx,
	}
}

// NewConsumeUserEntitlementParamsWithHTTPClient creates a new ConsumeUserEntitlementParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewConsumeUserEntitlementParamsWithHTTPClient(client *http.Client) *ConsumeUserEntitlementParams {
	var ()
	return &ConsumeUserEntitlementParams{
		HTTPClient: client,
	}
}

/*ConsumeUserEntitlementParams contains all the parameters to send to the API endpoint
for the consume user entitlement operation typically these are written to a http.Request
*/
type ConsumeUserEntitlementParams struct {

	/*Body*/
	Body *platformclientmodels.EntitlementDecrement
	/*EntitlementID*/
	EntitlementID string
	/*Namespace*/
	Namespace string
	/*UserID*/
	UserID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the consume user entitlement params
func (o *ConsumeUserEntitlementParams) WithTimeout(timeout time.Duration) *ConsumeUserEntitlementParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the consume user entitlement params
func (o *ConsumeUserEntitlementParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the consume user entitlement params
func (o *ConsumeUserEntitlementParams) WithContext(ctx context.Context) *ConsumeUserEntitlementParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the consume user entitlement params
func (o *ConsumeUserEntitlementParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the consume user entitlement params
func (o *ConsumeUserEntitlementParams) WithHTTPClient(client *http.Client) *ConsumeUserEntitlementParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the consume user entitlement params
func (o *ConsumeUserEntitlementParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the consume user entitlement params
func (o *ConsumeUserEntitlementParams) WithBody(body *platformclientmodels.EntitlementDecrement) *ConsumeUserEntitlementParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the consume user entitlement params
func (o *ConsumeUserEntitlementParams) SetBody(body *platformclientmodels.EntitlementDecrement) {
	o.Body = body
}

// WithEntitlementID adds the entitlementID to the consume user entitlement params
func (o *ConsumeUserEntitlementParams) WithEntitlementID(entitlementID string) *ConsumeUserEntitlementParams {
	o.SetEntitlementID(entitlementID)
	return o
}

// SetEntitlementID adds the entitlementId to the consume user entitlement params
func (o *ConsumeUserEntitlementParams) SetEntitlementID(entitlementID string) {
	o.EntitlementID = entitlementID
}

// WithNamespace adds the namespace to the consume user entitlement params
func (o *ConsumeUserEntitlementParams) WithNamespace(namespace string) *ConsumeUserEntitlementParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the consume user entitlement params
func (o *ConsumeUserEntitlementParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithUserID adds the userID to the consume user entitlement params
func (o *ConsumeUserEntitlementParams) WithUserID(userID string) *ConsumeUserEntitlementParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the consume user entitlement params
func (o *ConsumeUserEntitlementParams) SetUserID(userID string) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *ConsumeUserEntitlementParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param entitlementId
	if err := r.SetPathParam("entitlementId", o.EntitlementID); err != nil {
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
