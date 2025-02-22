// Code generated by go-swagger; DO NOT EDIT.

package users

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

// NewAdminDeleteUserRolesV3Params creates a new AdminDeleteUserRolesV3Params object
// with the default values initialized.
func NewAdminDeleteUserRolesV3Params() *AdminDeleteUserRolesV3Params {
	var ()
	return &AdminDeleteUserRolesV3Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewAdminDeleteUserRolesV3ParamsWithTimeout creates a new AdminDeleteUserRolesV3Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewAdminDeleteUserRolesV3ParamsWithTimeout(timeout time.Duration) *AdminDeleteUserRolesV3Params {
	var ()
	return &AdminDeleteUserRolesV3Params{

		timeout: timeout,
	}
}

// NewAdminDeleteUserRolesV3ParamsWithContext creates a new AdminDeleteUserRolesV3Params object
// with the default values initialized, and the ability to set a context for a request
func NewAdminDeleteUserRolesV3ParamsWithContext(ctx context.Context) *AdminDeleteUserRolesV3Params {
	var ()
	return &AdminDeleteUserRolesV3Params{

		Context: ctx,
	}
}

// NewAdminDeleteUserRolesV3ParamsWithHTTPClient creates a new AdminDeleteUserRolesV3Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAdminDeleteUserRolesV3ParamsWithHTTPClient(client *http.Client) *AdminDeleteUserRolesV3Params {
	var ()
	return &AdminDeleteUserRolesV3Params{
		HTTPClient: client,
	}
}

/*AdminDeleteUserRolesV3Params contains all the parameters to send to the API endpoint
for the admin delete user roles v3 operation typically these are written to a http.Request
*/
type AdminDeleteUserRolesV3Params struct {

	/*Body*/
	Body []string
	/*Namespace
	  Namespace, only accept alphabet and numeric

	*/
	Namespace string
	/*UserID
	  User id

	*/
	UserID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the admin delete user roles v3 params
func (o *AdminDeleteUserRolesV3Params) WithTimeout(timeout time.Duration) *AdminDeleteUserRolesV3Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the admin delete user roles v3 params
func (o *AdminDeleteUserRolesV3Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the admin delete user roles v3 params
func (o *AdminDeleteUserRolesV3Params) WithContext(ctx context.Context) *AdminDeleteUserRolesV3Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the admin delete user roles v3 params
func (o *AdminDeleteUserRolesV3Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the admin delete user roles v3 params
func (o *AdminDeleteUserRolesV3Params) WithHTTPClient(client *http.Client) *AdminDeleteUserRolesV3Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the admin delete user roles v3 params
func (o *AdminDeleteUserRolesV3Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the admin delete user roles v3 params
func (o *AdminDeleteUserRolesV3Params) WithBody(body []string) *AdminDeleteUserRolesV3Params {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the admin delete user roles v3 params
func (o *AdminDeleteUserRolesV3Params) SetBody(body []string) {
	o.Body = body
}

// WithNamespace adds the namespace to the admin delete user roles v3 params
func (o *AdminDeleteUserRolesV3Params) WithNamespace(namespace string) *AdminDeleteUserRolesV3Params {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the admin delete user roles v3 params
func (o *AdminDeleteUserRolesV3Params) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithUserID adds the userID to the admin delete user roles v3 params
func (o *AdminDeleteUserRolesV3Params) WithUserID(userID string) *AdminDeleteUserRolesV3Params {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the admin delete user roles v3 params
func (o *AdminDeleteUserRolesV3Params) SetUserID(userID string) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *AdminDeleteUserRolesV3Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
