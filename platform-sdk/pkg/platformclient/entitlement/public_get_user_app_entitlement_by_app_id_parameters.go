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

// NewPublicGetUserAppEntitlementByAppIDParams creates a new PublicGetUserAppEntitlementByAppIDParams object
// with the default values initialized.
func NewPublicGetUserAppEntitlementByAppIDParams() *PublicGetUserAppEntitlementByAppIDParams {
	var ()
	return &PublicGetUserAppEntitlementByAppIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPublicGetUserAppEntitlementByAppIDParamsWithTimeout creates a new PublicGetUserAppEntitlementByAppIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPublicGetUserAppEntitlementByAppIDParamsWithTimeout(timeout time.Duration) *PublicGetUserAppEntitlementByAppIDParams {
	var ()
	return &PublicGetUserAppEntitlementByAppIDParams{

		timeout: timeout,
	}
}

// NewPublicGetUserAppEntitlementByAppIDParamsWithContext creates a new PublicGetUserAppEntitlementByAppIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewPublicGetUserAppEntitlementByAppIDParamsWithContext(ctx context.Context) *PublicGetUserAppEntitlementByAppIDParams {
	var ()
	return &PublicGetUserAppEntitlementByAppIDParams{

		Context: ctx,
	}
}

// NewPublicGetUserAppEntitlementByAppIDParamsWithHTTPClient creates a new PublicGetUserAppEntitlementByAppIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPublicGetUserAppEntitlementByAppIDParamsWithHTTPClient(client *http.Client) *PublicGetUserAppEntitlementByAppIDParams {
	var ()
	return &PublicGetUserAppEntitlementByAppIDParams{
		HTTPClient: client,
	}
}

/*PublicGetUserAppEntitlementByAppIDParams contains all the parameters to send to the API endpoint
for the public get user app entitlement by app Id operation typically these are written to a http.Request
*/
type PublicGetUserAppEntitlementByAppIDParams struct {

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

// WithTimeout adds the timeout to the public get user app entitlement by app Id params
func (o *PublicGetUserAppEntitlementByAppIDParams) WithTimeout(timeout time.Duration) *PublicGetUserAppEntitlementByAppIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the public get user app entitlement by app Id params
func (o *PublicGetUserAppEntitlementByAppIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the public get user app entitlement by app Id params
func (o *PublicGetUserAppEntitlementByAppIDParams) WithContext(ctx context.Context) *PublicGetUserAppEntitlementByAppIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the public get user app entitlement by app Id params
func (o *PublicGetUserAppEntitlementByAppIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the public get user app entitlement by app Id params
func (o *PublicGetUserAppEntitlementByAppIDParams) WithHTTPClient(client *http.Client) *PublicGetUserAppEntitlementByAppIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the public get user app entitlement by app Id params
func (o *PublicGetUserAppEntitlementByAppIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAppID adds the appID to the public get user app entitlement by app Id params
func (o *PublicGetUserAppEntitlementByAppIDParams) WithAppID(appID string) *PublicGetUserAppEntitlementByAppIDParams {
	o.SetAppID(appID)
	return o
}

// SetAppID adds the appId to the public get user app entitlement by app Id params
func (o *PublicGetUserAppEntitlementByAppIDParams) SetAppID(appID string) {
	o.AppID = appID
}

// WithNamespace adds the namespace to the public get user app entitlement by app Id params
func (o *PublicGetUserAppEntitlementByAppIDParams) WithNamespace(namespace string) *PublicGetUserAppEntitlementByAppIDParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the public get user app entitlement by app Id params
func (o *PublicGetUserAppEntitlementByAppIDParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithUserID adds the userID to the public get user app entitlement by app Id params
func (o *PublicGetUserAppEntitlementByAppIDParams) WithUserID(userID string) *PublicGetUserAppEntitlementByAppIDParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the public get user app entitlement by app Id params
func (o *PublicGetUserAppEntitlementByAppIDParams) SetUserID(userID string) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *PublicGetUserAppEntitlementByAppIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
