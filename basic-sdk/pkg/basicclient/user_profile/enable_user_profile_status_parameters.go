// Code generated by go-swagger; DO NOT EDIT.

package user_profile

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

	"github.com/AccelByte/accelbyte-go-sdk/basic-sdk/pkg/basicclientmodels"
)

// NewEnableUserProfileStatusParams creates a new EnableUserProfileStatusParams object
// with the default values initialized.
func NewEnableUserProfileStatusParams() *EnableUserProfileStatusParams {
	var ()
	return &EnableUserProfileStatusParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewEnableUserProfileStatusParamsWithTimeout creates a new EnableUserProfileStatusParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewEnableUserProfileStatusParamsWithTimeout(timeout time.Duration) *EnableUserProfileStatusParams {
	var ()
	return &EnableUserProfileStatusParams{

		timeout: timeout,
	}
}

// NewEnableUserProfileStatusParamsWithContext creates a new EnableUserProfileStatusParams object
// with the default values initialized, and the ability to set a context for a request
func NewEnableUserProfileStatusParamsWithContext(ctx context.Context) *EnableUserProfileStatusParams {
	var ()
	return &EnableUserProfileStatusParams{

		Context: ctx,
	}
}

// NewEnableUserProfileStatusParamsWithHTTPClient creates a new EnableUserProfileStatusParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewEnableUserProfileStatusParamsWithHTTPClient(client *http.Client) *EnableUserProfileStatusParams {
	var ()
	return &EnableUserProfileStatusParams{
		HTTPClient: client,
	}
}

/*EnableUserProfileStatusParams contains all the parameters to send to the API endpoint
for the enable user profile status operation typically these are written to a http.Request
*/
type EnableUserProfileStatusParams struct {

	/*Body*/
	Body *basicclientmodels.UserProfileStatusUpdate
	/*Namespace
	  namespace, only accept alphabet and numeric

	*/
	Namespace string
	/*UserID
	  user's id, should follow UUID version 4 without hyphen

	*/
	UserID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the enable user profile status params
func (o *EnableUserProfileStatusParams) WithTimeout(timeout time.Duration) *EnableUserProfileStatusParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the enable user profile status params
func (o *EnableUserProfileStatusParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the enable user profile status params
func (o *EnableUserProfileStatusParams) WithContext(ctx context.Context) *EnableUserProfileStatusParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the enable user profile status params
func (o *EnableUserProfileStatusParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the enable user profile status params
func (o *EnableUserProfileStatusParams) WithHTTPClient(client *http.Client) *EnableUserProfileStatusParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the enable user profile status params
func (o *EnableUserProfileStatusParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the enable user profile status params
func (o *EnableUserProfileStatusParams) WithBody(body *basicclientmodels.UserProfileStatusUpdate) *EnableUserProfileStatusParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the enable user profile status params
func (o *EnableUserProfileStatusParams) SetBody(body *basicclientmodels.UserProfileStatusUpdate) {
	o.Body = body
}

// WithNamespace adds the namespace to the enable user profile status params
func (o *EnableUserProfileStatusParams) WithNamespace(namespace string) *EnableUserProfileStatusParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the enable user profile status params
func (o *EnableUserProfileStatusParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithUserID adds the userID to the enable user profile status params
func (o *EnableUserProfileStatusParams) WithUserID(userID string) *EnableUserProfileStatusParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the enable user profile status params
func (o *EnableUserProfileStatusParams) SetUserID(userID string) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *EnableUserProfileStatusParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
