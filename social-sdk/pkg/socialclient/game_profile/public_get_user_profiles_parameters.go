// Code generated by go-swagger; DO NOT EDIT.

package game_profile

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

// NewPublicGetUserProfilesParams creates a new PublicGetUserProfilesParams object
// with the default values initialized.
func NewPublicGetUserProfilesParams() *PublicGetUserProfilesParams {
	var ()
	return &PublicGetUserProfilesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPublicGetUserProfilesParamsWithTimeout creates a new PublicGetUserProfilesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPublicGetUserProfilesParamsWithTimeout(timeout time.Duration) *PublicGetUserProfilesParams {
	var ()
	return &PublicGetUserProfilesParams{

		timeout: timeout,
	}
}

// NewPublicGetUserProfilesParamsWithContext creates a new PublicGetUserProfilesParams object
// with the default values initialized, and the ability to set a context for a request
func NewPublicGetUserProfilesParamsWithContext(ctx context.Context) *PublicGetUserProfilesParams {
	var ()
	return &PublicGetUserProfilesParams{

		Context: ctx,
	}
}

// NewPublicGetUserProfilesParamsWithHTTPClient creates a new PublicGetUserProfilesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPublicGetUserProfilesParamsWithHTTPClient(client *http.Client) *PublicGetUserProfilesParams {
	var ()
	return &PublicGetUserProfilesParams{
		HTTPClient: client,
	}
}

/*PublicGetUserProfilesParams contains all the parameters to send to the API endpoint
for the public get user profiles operation typically these are written to a http.Request
*/
type PublicGetUserProfilesParams struct {

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

// WithTimeout adds the timeout to the public get user profiles params
func (o *PublicGetUserProfilesParams) WithTimeout(timeout time.Duration) *PublicGetUserProfilesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the public get user profiles params
func (o *PublicGetUserProfilesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the public get user profiles params
func (o *PublicGetUserProfilesParams) WithContext(ctx context.Context) *PublicGetUserProfilesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the public get user profiles params
func (o *PublicGetUserProfilesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the public get user profiles params
func (o *PublicGetUserProfilesParams) WithHTTPClient(client *http.Client) *PublicGetUserProfilesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the public get user profiles params
func (o *PublicGetUserProfilesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNamespace adds the namespace to the public get user profiles params
func (o *PublicGetUserProfilesParams) WithNamespace(namespace string) *PublicGetUserProfilesParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the public get user profiles params
func (o *PublicGetUserProfilesParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithUserID adds the userID to the public get user profiles params
func (o *PublicGetUserProfilesParams) WithUserID(userID string) *PublicGetUserProfilesParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the public get user profiles params
func (o *PublicGetUserProfilesParams) SetUserID(userID string) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *PublicGetUserProfilesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
