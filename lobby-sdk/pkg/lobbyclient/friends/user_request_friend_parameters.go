// Code generated by go-swagger; DO NOT EDIT.

package friends

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

	"github.com/AccelByte/accelbyte-go-sdk/lobby-sdk/pkg/lobbyclientmodels"
)

// NewUserRequestFriendParams creates a new UserRequestFriendParams object
// with the default values initialized.
func NewUserRequestFriendParams() *UserRequestFriendParams {
	var ()
	return &UserRequestFriendParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUserRequestFriendParamsWithTimeout creates a new UserRequestFriendParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUserRequestFriendParamsWithTimeout(timeout time.Duration) *UserRequestFriendParams {
	var ()
	return &UserRequestFriendParams{

		timeout: timeout,
	}
}

// NewUserRequestFriendParamsWithContext creates a new UserRequestFriendParams object
// with the default values initialized, and the ability to set a context for a request
func NewUserRequestFriendParamsWithContext(ctx context.Context) *UserRequestFriendParams {
	var ()
	return &UserRequestFriendParams{

		Context: ctx,
	}
}

// NewUserRequestFriendParamsWithHTTPClient creates a new UserRequestFriendParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUserRequestFriendParamsWithHTTPClient(client *http.Client) *UserRequestFriendParams {
	var ()
	return &UserRequestFriendParams{
		HTTPClient: client,
	}
}

/*UserRequestFriendParams contains all the parameters to send to the API endpoint
for the user request friend operation typically these are written to a http.Request
*/
type UserRequestFriendParams struct {

	/*Body
	  add friend content

	*/
	Body *lobbyclientmodels.ModelRequestFriendsRequest
	/*Namespace
	  namespace

	*/
	Namespace string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the user request friend params
func (o *UserRequestFriendParams) WithTimeout(timeout time.Duration) *UserRequestFriendParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the user request friend params
func (o *UserRequestFriendParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the user request friend params
func (o *UserRequestFriendParams) WithContext(ctx context.Context) *UserRequestFriendParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the user request friend params
func (o *UserRequestFriendParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the user request friend params
func (o *UserRequestFriendParams) WithHTTPClient(client *http.Client) *UserRequestFriendParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the user request friend params
func (o *UserRequestFriendParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the user request friend params
func (o *UserRequestFriendParams) WithBody(body *lobbyclientmodels.ModelRequestFriendsRequest) *UserRequestFriendParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the user request friend params
func (o *UserRequestFriendParams) SetBody(body *lobbyclientmodels.ModelRequestFriendsRequest) {
	o.Body = body
}

// WithNamespace adds the namespace to the user request friend params
func (o *UserRequestFriendParams) WithNamespace(namespace string) *UserRequestFriendParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the user request friend params
func (o *UserRequestFriendParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WriteToRequest writes these params to a swagger request
func (o *UserRequestFriendParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
