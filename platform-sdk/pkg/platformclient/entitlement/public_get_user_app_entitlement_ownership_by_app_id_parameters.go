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

// NewPublicGetUserAppEntitlementOwnershipByAppIDParams creates a new PublicGetUserAppEntitlementOwnershipByAppIDParams object
// with the default values initialized.
func NewPublicGetUserAppEntitlementOwnershipByAppIDParams() *PublicGetUserAppEntitlementOwnershipByAppIDParams {
	var ()
	return &PublicGetUserAppEntitlementOwnershipByAppIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPublicGetUserAppEntitlementOwnershipByAppIDParamsWithTimeout creates a new PublicGetUserAppEntitlementOwnershipByAppIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPublicGetUserAppEntitlementOwnershipByAppIDParamsWithTimeout(timeout time.Duration) *PublicGetUserAppEntitlementOwnershipByAppIDParams {
	var ()
	return &PublicGetUserAppEntitlementOwnershipByAppIDParams{

		timeout: timeout,
	}
}

// NewPublicGetUserAppEntitlementOwnershipByAppIDParamsWithContext creates a new PublicGetUserAppEntitlementOwnershipByAppIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewPublicGetUserAppEntitlementOwnershipByAppIDParamsWithContext(ctx context.Context) *PublicGetUserAppEntitlementOwnershipByAppIDParams {
	var ()
	return &PublicGetUserAppEntitlementOwnershipByAppIDParams{

		Context: ctx,
	}
}

// NewPublicGetUserAppEntitlementOwnershipByAppIDParamsWithHTTPClient creates a new PublicGetUserAppEntitlementOwnershipByAppIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPublicGetUserAppEntitlementOwnershipByAppIDParamsWithHTTPClient(client *http.Client) *PublicGetUserAppEntitlementOwnershipByAppIDParams {
	var ()
	return &PublicGetUserAppEntitlementOwnershipByAppIDParams{
		HTTPClient: client,
	}
}

/*PublicGetUserAppEntitlementOwnershipByAppIDParams contains all the parameters to send to the API endpoint
for the public get user app entitlement ownership by app Id operation typically these are written to a http.Request
*/
type PublicGetUserAppEntitlementOwnershipByAppIDParams struct {

	/*AppID*/
	AppID string
	/*Namespace*/
	Namespace string
	/*UserID*/
	UserID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the public get user app entitlement ownership by app Id params
func (o *PublicGetUserAppEntitlementOwnershipByAppIDParams) WithTimeout(timeout time.Duration) *PublicGetUserAppEntitlementOwnershipByAppIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the public get user app entitlement ownership by app Id params
func (o *PublicGetUserAppEntitlementOwnershipByAppIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the public get user app entitlement ownership by app Id params
func (o *PublicGetUserAppEntitlementOwnershipByAppIDParams) WithContext(ctx context.Context) *PublicGetUserAppEntitlementOwnershipByAppIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the public get user app entitlement ownership by app Id params
func (o *PublicGetUserAppEntitlementOwnershipByAppIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the public get user app entitlement ownership by app Id params
func (o *PublicGetUserAppEntitlementOwnershipByAppIDParams) WithHTTPClient(client *http.Client) *PublicGetUserAppEntitlementOwnershipByAppIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the public get user app entitlement ownership by app Id params
func (o *PublicGetUserAppEntitlementOwnershipByAppIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAppID adds the appID to the public get user app entitlement ownership by app Id params
func (o *PublicGetUserAppEntitlementOwnershipByAppIDParams) WithAppID(appID string) *PublicGetUserAppEntitlementOwnershipByAppIDParams {
	o.SetAppID(appID)
	return o
}

// SetAppID adds the appId to the public get user app entitlement ownership by app Id params
func (o *PublicGetUserAppEntitlementOwnershipByAppIDParams) SetAppID(appID string) {
	o.AppID = appID
}

// WithNamespace adds the namespace to the public get user app entitlement ownership by app Id params
func (o *PublicGetUserAppEntitlementOwnershipByAppIDParams) WithNamespace(namespace string) *PublicGetUserAppEntitlementOwnershipByAppIDParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the public get user app entitlement ownership by app Id params
func (o *PublicGetUserAppEntitlementOwnershipByAppIDParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithUserID adds the userID to the public get user app entitlement ownership by app Id params
func (o *PublicGetUserAppEntitlementOwnershipByAppIDParams) WithUserID(userID string) *PublicGetUserAppEntitlementOwnershipByAppIDParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the public get user app entitlement ownership by app Id params
func (o *PublicGetUserAppEntitlementOwnershipByAppIDParams) SetUserID(userID string) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *PublicGetUserAppEntitlementOwnershipByAppIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param appId
	qrAppID := o.AppID
	qAppID := qrAppID
	if qAppID != "" {
		if err := r.SetQueryParam("appId", qAppID); err != nil {
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
