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

	"github.com/AccelByte/accelbyte-go-sdk/dsmc-sdk/pkg/dsmcclientmodels"
)

// NewUpdatePodConfigParams creates a new UpdatePodConfigParams object
// with the default values initialized.
func NewUpdatePodConfigParams() *UpdatePodConfigParams {
	var ()
	return &UpdatePodConfigParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdatePodConfigParamsWithTimeout creates a new UpdatePodConfigParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdatePodConfigParamsWithTimeout(timeout time.Duration) *UpdatePodConfigParams {
	var ()
	return &UpdatePodConfigParams{

		timeout: timeout,
	}
}

// NewUpdatePodConfigParamsWithContext creates a new UpdatePodConfigParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdatePodConfigParamsWithContext(ctx context.Context) *UpdatePodConfigParams {
	var ()
	return &UpdatePodConfigParams{

		Context: ctx,
	}
}

// NewUpdatePodConfigParamsWithHTTPClient creates a new UpdatePodConfigParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdatePodConfigParamsWithHTTPClient(client *http.Client) *UpdatePodConfigParams {
	var ()
	return &UpdatePodConfigParams{
		HTTPClient: client,
	}
}

/*UpdatePodConfigParams contains all the parameters to send to the API endpoint
for the update pod config operation typically these are written to a http.Request
*/
type UpdatePodConfigParams struct {

	/*Body*/
	Body *dsmcclientmodels.ModelsUpdatePodConfigRequest
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

// WithTimeout adds the timeout to the update pod config params
func (o *UpdatePodConfigParams) WithTimeout(timeout time.Duration) *UpdatePodConfigParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update pod config params
func (o *UpdatePodConfigParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update pod config params
func (o *UpdatePodConfigParams) WithContext(ctx context.Context) *UpdatePodConfigParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update pod config params
func (o *UpdatePodConfigParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update pod config params
func (o *UpdatePodConfigParams) WithHTTPClient(client *http.Client) *UpdatePodConfigParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update pod config params
func (o *UpdatePodConfigParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update pod config params
func (o *UpdatePodConfigParams) WithBody(body *dsmcclientmodels.ModelsUpdatePodConfigRequest) *UpdatePodConfigParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update pod config params
func (o *UpdatePodConfigParams) SetBody(body *dsmcclientmodels.ModelsUpdatePodConfigRequest) {
	o.Body = body
}

// WithName adds the name to the update pod config params
func (o *UpdatePodConfigParams) WithName(name string) *UpdatePodConfigParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the update pod config params
func (o *UpdatePodConfigParams) SetName(name string) {
	o.Name = name
}

// WithNamespace adds the namespace to the update pod config params
func (o *UpdatePodConfigParams) WithNamespace(namespace string) *UpdatePodConfigParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the update pod config params
func (o *UpdatePodConfigParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WriteToRequest writes these params to a swagger request
func (o *UpdatePodConfigParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

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
