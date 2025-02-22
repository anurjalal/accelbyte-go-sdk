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

// NewGetPublisherUserParams creates a new GetPublisherUserParams object
// with the default values initialized.
func NewGetPublisherUserParams() *GetPublisherUserParams {
	var ()
	return &GetPublisherUserParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetPublisherUserParamsWithTimeout creates a new GetPublisherUserParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetPublisherUserParamsWithTimeout(timeout time.Duration) *GetPublisherUserParams {
	var ()
	return &GetPublisherUserParams{

		timeout: timeout,
	}
}

// NewGetPublisherUserParamsWithContext creates a new GetPublisherUserParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetPublisherUserParamsWithContext(ctx context.Context) *GetPublisherUserParams {
	var ()
	return &GetPublisherUserParams{

		Context: ctx,
	}
}

// NewGetPublisherUserParamsWithHTTPClient creates a new GetPublisherUserParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetPublisherUserParamsWithHTTPClient(client *http.Client) *GetPublisherUserParams {
	var ()
	return &GetPublisherUserParams{
		HTTPClient: client,
	}
}

/*GetPublisherUserParams contains all the parameters to send to the API endpoint
for the get publisher user operation typically these are written to a http.Request
*/
type GetPublisherUserParams struct {

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

// WithTimeout adds the timeout to the get publisher user params
func (o *GetPublisherUserParams) WithTimeout(timeout time.Duration) *GetPublisherUserParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get publisher user params
func (o *GetPublisherUserParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get publisher user params
func (o *GetPublisherUserParams) WithContext(ctx context.Context) *GetPublisherUserParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get publisher user params
func (o *GetPublisherUserParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get publisher user params
func (o *GetPublisherUserParams) WithHTTPClient(client *http.Client) *GetPublisherUserParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get publisher user params
func (o *GetPublisherUserParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNamespace adds the namespace to the get publisher user params
func (o *GetPublisherUserParams) WithNamespace(namespace string) *GetPublisherUserParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the get publisher user params
func (o *GetPublisherUserParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithUserID adds the userID to the get publisher user params
func (o *GetPublisherUserParams) WithUserID(userID string) *GetPublisherUserParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the get publisher user params
func (o *GetPublisherUserParams) SetUserID(userID string) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *GetPublisherUserParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
