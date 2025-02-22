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

// NewGrantUserEntitlementParams creates a new GrantUserEntitlementParams object
// with the default values initialized.
func NewGrantUserEntitlementParams() *GrantUserEntitlementParams {
	var ()
	return &GrantUserEntitlementParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGrantUserEntitlementParamsWithTimeout creates a new GrantUserEntitlementParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGrantUserEntitlementParamsWithTimeout(timeout time.Duration) *GrantUserEntitlementParams {
	var ()
	return &GrantUserEntitlementParams{

		timeout: timeout,
	}
}

// NewGrantUserEntitlementParamsWithContext creates a new GrantUserEntitlementParams object
// with the default values initialized, and the ability to set a context for a request
func NewGrantUserEntitlementParamsWithContext(ctx context.Context) *GrantUserEntitlementParams {
	var ()
	return &GrantUserEntitlementParams{

		Context: ctx,
	}
}

// NewGrantUserEntitlementParamsWithHTTPClient creates a new GrantUserEntitlementParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGrantUserEntitlementParamsWithHTTPClient(client *http.Client) *GrantUserEntitlementParams {
	var ()
	return &GrantUserEntitlementParams{
		HTTPClient: client,
	}
}

/*GrantUserEntitlementParams contains all the parameters to send to the API endpoint
for the grant user entitlement operation typically these are written to a http.Request
*/
type GrantUserEntitlementParams struct {

	/*Body*/
	Body []*platformclientmodels.EntitlementGrant
	/*Namespace*/
	Namespace string
	/*UserID*/
	UserID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the grant user entitlement params
func (o *GrantUserEntitlementParams) WithTimeout(timeout time.Duration) *GrantUserEntitlementParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the grant user entitlement params
func (o *GrantUserEntitlementParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the grant user entitlement params
func (o *GrantUserEntitlementParams) WithContext(ctx context.Context) *GrantUserEntitlementParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the grant user entitlement params
func (o *GrantUserEntitlementParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the grant user entitlement params
func (o *GrantUserEntitlementParams) WithHTTPClient(client *http.Client) *GrantUserEntitlementParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the grant user entitlement params
func (o *GrantUserEntitlementParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the grant user entitlement params
func (o *GrantUserEntitlementParams) WithBody(body []*platformclientmodels.EntitlementGrant) *GrantUserEntitlementParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the grant user entitlement params
func (o *GrantUserEntitlementParams) SetBody(body []*platformclientmodels.EntitlementGrant) {
	o.Body = body
}

// WithNamespace adds the namespace to the grant user entitlement params
func (o *GrantUserEntitlementParams) WithNamespace(namespace string) *GrantUserEntitlementParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the grant user entitlement params
func (o *GrantUserEntitlementParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithUserID adds the userID to the grant user entitlement params
func (o *GrantUserEntitlementParams) WithUserID(userID string) *GrantUserEntitlementParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the grant user entitlement params
func (o *GrantUserEntitlementParams) SetUserID(userID string) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *GrantUserEntitlementParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
