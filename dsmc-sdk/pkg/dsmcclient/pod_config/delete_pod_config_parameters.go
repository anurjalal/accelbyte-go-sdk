// Code generated by go-swagger; DO NOT EDIT.

package pod_config

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

// NewDeletePodConfigParams creates a new DeletePodConfigParams object
// with the default values initialized.
func NewDeletePodConfigParams() *DeletePodConfigParams {
	var ()
	return &DeletePodConfigParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeletePodConfigParamsWithTimeout creates a new DeletePodConfigParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeletePodConfigParamsWithTimeout(timeout time.Duration) *DeletePodConfigParams {
	var ()
	return &DeletePodConfigParams{

		timeout: timeout,
	}
}

// NewDeletePodConfigParamsWithContext creates a new DeletePodConfigParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeletePodConfigParamsWithContext(ctx context.Context) *DeletePodConfigParams {
	var ()
	return &DeletePodConfigParams{

		Context: ctx,
	}
}

// NewDeletePodConfigParamsWithHTTPClient creates a new DeletePodConfigParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeletePodConfigParamsWithHTTPClient(client *http.Client) *DeletePodConfigParams {
	var ()
	return &DeletePodConfigParams{
		HTTPClient: client,
	}
}

/*DeletePodConfigParams contains all the parameters to send to the API endpoint
for the delete pod config operation typically these are written to a http.Request
*/
type DeletePodConfigParams struct {

	/*Name
	  pod name

	*/
	Name string
	/*Namespace
	  namespace of the game

	*/
	Namespace string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete pod config params
func (o *DeletePodConfigParams) WithTimeout(timeout time.Duration) *DeletePodConfigParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete pod config params
func (o *DeletePodConfigParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete pod config params
func (o *DeletePodConfigParams) WithContext(ctx context.Context) *DeletePodConfigParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete pod config params
func (o *DeletePodConfigParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete pod config params
func (o *DeletePodConfigParams) WithHTTPClient(client *http.Client) *DeletePodConfigParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete pod config params
func (o *DeletePodConfigParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the delete pod config params
func (o *DeletePodConfigParams) WithName(name string) *DeletePodConfigParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the delete pod config params
func (o *DeletePodConfigParams) SetName(name string) {
	o.Name = name
}

// WithNamespace adds the namespace to the delete pod config params
func (o *DeletePodConfigParams) WithNamespace(namespace string) *DeletePodConfigParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the delete pod config params
func (o *DeletePodConfigParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WriteToRequest writes these params to a swagger request
func (o *DeletePodConfigParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
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
