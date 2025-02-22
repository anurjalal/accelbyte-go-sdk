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

// NewDeleteUserInformationParams creates a new DeleteUserInformationParams object
// with the default values initialized.
func NewDeleteUserInformationParams() *DeleteUserInformationParams {
	var ()
	return &DeleteUserInformationParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteUserInformationParamsWithTimeout creates a new DeleteUserInformationParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteUserInformationParamsWithTimeout(timeout time.Duration) *DeleteUserInformationParams {
	var ()
	return &DeleteUserInformationParams{

		timeout: timeout,
	}
}

// NewDeleteUserInformationParamsWithContext creates a new DeleteUserInformationParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteUserInformationParamsWithContext(ctx context.Context) *DeleteUserInformationParams {
	var ()
	return &DeleteUserInformationParams{

		Context: ctx,
	}
}

// NewDeleteUserInformationParamsWithHTTPClient creates a new DeleteUserInformationParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteUserInformationParamsWithHTTPClient(client *http.Client) *DeleteUserInformationParams {
	var ()
	return &DeleteUserInformationParams{
		HTTPClient: client,
	}
}

/*DeleteUserInformationParams contains all the parameters to send to the API endpoint
for the delete user information operation typically these are written to a http.Request
*/
type DeleteUserInformationParams struct {

	/*Namespace
	  Namespace, only accept alphabet and numeric

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

// WithTimeout adds the timeout to the delete user information params
func (o *DeleteUserInformationParams) WithTimeout(timeout time.Duration) *DeleteUserInformationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete user information params
func (o *DeleteUserInformationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete user information params
func (o *DeleteUserInformationParams) WithContext(ctx context.Context) *DeleteUserInformationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete user information params
func (o *DeleteUserInformationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete user information params
func (o *DeleteUserInformationParams) WithHTTPClient(client *http.Client) *DeleteUserInformationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete user information params
func (o *DeleteUserInformationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNamespace adds the namespace to the delete user information params
func (o *DeleteUserInformationParams) WithNamespace(namespace string) *DeleteUserInformationParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the delete user information params
func (o *DeleteUserInformationParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithUserID adds the userID to the delete user information params
func (o *DeleteUserInformationParams) WithUserID(userID string) *DeleteUserInformationParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the delete user information params
func (o *DeleteUserInformationParams) SetUserID(userID string) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteUserInformationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
