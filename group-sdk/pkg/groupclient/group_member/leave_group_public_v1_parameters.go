// Code generated by go-swagger; DO NOT EDIT.

package group_member

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

// NewLeaveGroupPublicV1Params creates a new LeaveGroupPublicV1Params object
// with the default values initialized.
func NewLeaveGroupPublicV1Params() *LeaveGroupPublicV1Params {
	var ()
	return &LeaveGroupPublicV1Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewLeaveGroupPublicV1ParamsWithTimeout creates a new LeaveGroupPublicV1Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewLeaveGroupPublicV1ParamsWithTimeout(timeout time.Duration) *LeaveGroupPublicV1Params {
	var ()
	return &LeaveGroupPublicV1Params{

		timeout: timeout,
	}
}

// NewLeaveGroupPublicV1ParamsWithContext creates a new LeaveGroupPublicV1Params object
// with the default values initialized, and the ability to set a context for a request
func NewLeaveGroupPublicV1ParamsWithContext(ctx context.Context) *LeaveGroupPublicV1Params {
	var ()
	return &LeaveGroupPublicV1Params{

		Context: ctx,
	}
}

// NewLeaveGroupPublicV1ParamsWithHTTPClient creates a new LeaveGroupPublicV1Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewLeaveGroupPublicV1ParamsWithHTTPClient(client *http.Client) *LeaveGroupPublicV1Params {
	var ()
	return &LeaveGroupPublicV1Params{
		HTTPClient: client,
	}
}

/*LeaveGroupPublicV1Params contains all the parameters to send to the API endpoint
for the leave group public v1 operation typically these are written to a http.Request
*/
type LeaveGroupPublicV1Params struct {

	/*Namespace
	  namespace

	*/
	Namespace string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the leave group public v1 params
func (o *LeaveGroupPublicV1Params) WithTimeout(timeout time.Duration) *LeaveGroupPublicV1Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the leave group public v1 params
func (o *LeaveGroupPublicV1Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the leave group public v1 params
func (o *LeaveGroupPublicV1Params) WithContext(ctx context.Context) *LeaveGroupPublicV1Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the leave group public v1 params
func (o *LeaveGroupPublicV1Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the leave group public v1 params
func (o *LeaveGroupPublicV1Params) WithHTTPClient(client *http.Client) *LeaveGroupPublicV1Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the leave group public v1 params
func (o *LeaveGroupPublicV1Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNamespace adds the namespace to the leave group public v1 params
func (o *LeaveGroupPublicV1Params) WithNamespace(namespace string) *LeaveGroupPublicV1Params {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the leave group public v1 params
func (o *LeaveGroupPublicV1Params) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WriteToRequest writes these params to a swagger request
func (o *LeaveGroupPublicV1Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param namespace
	if err := r.SetPathParam("namespace", o.Namespace); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
