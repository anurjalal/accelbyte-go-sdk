// Code generated by go-swagger; DO NOT EDIT.

package deployment_config

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

// NewDeleteDeploymentOverrideParams creates a new DeleteDeploymentOverrideParams object
// with the default values initialized.
func NewDeleteDeploymentOverrideParams() *DeleteDeploymentOverrideParams {
	var ()
	return &DeleteDeploymentOverrideParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteDeploymentOverrideParamsWithTimeout creates a new DeleteDeploymentOverrideParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteDeploymentOverrideParamsWithTimeout(timeout time.Duration) *DeleteDeploymentOverrideParams {
	var ()
	return &DeleteDeploymentOverrideParams{

		timeout: timeout,
	}
}

// NewDeleteDeploymentOverrideParamsWithContext creates a new DeleteDeploymentOverrideParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteDeploymentOverrideParamsWithContext(ctx context.Context) *DeleteDeploymentOverrideParams {
	var ()
	return &DeleteDeploymentOverrideParams{

		Context: ctx,
	}
}

// NewDeleteDeploymentOverrideParamsWithHTTPClient creates a new DeleteDeploymentOverrideParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteDeploymentOverrideParamsWithHTTPClient(client *http.Client) *DeleteDeploymentOverrideParams {
	var ()
	return &DeleteDeploymentOverrideParams{
		HTTPClient: client,
	}
}

/*DeleteDeploymentOverrideParams contains all the parameters to send to the API endpoint
for the delete deployment override operation typically these are written to a http.Request
*/
type DeleteDeploymentOverrideParams struct {

	/*Deployment
	  deployment name

	*/
	Deployment string
	/*Namespace
	  namespace of the game

	*/
	Namespace string
	/*Version
	  version

	*/
	Version string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete deployment override params
func (o *DeleteDeploymentOverrideParams) WithTimeout(timeout time.Duration) *DeleteDeploymentOverrideParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete deployment override params
func (o *DeleteDeploymentOverrideParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete deployment override params
func (o *DeleteDeploymentOverrideParams) WithContext(ctx context.Context) *DeleteDeploymentOverrideParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete deployment override params
func (o *DeleteDeploymentOverrideParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete deployment override params
func (o *DeleteDeploymentOverrideParams) WithHTTPClient(client *http.Client) *DeleteDeploymentOverrideParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete deployment override params
func (o *DeleteDeploymentOverrideParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDeployment adds the deployment to the delete deployment override params
func (o *DeleteDeploymentOverrideParams) WithDeployment(deployment string) *DeleteDeploymentOverrideParams {
	o.SetDeployment(deployment)
	return o
}

// SetDeployment adds the deployment to the delete deployment override params
func (o *DeleteDeploymentOverrideParams) SetDeployment(deployment string) {
	o.Deployment = deployment
}

// WithNamespace adds the namespace to the delete deployment override params
func (o *DeleteDeploymentOverrideParams) WithNamespace(namespace string) *DeleteDeploymentOverrideParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the delete deployment override params
func (o *DeleteDeploymentOverrideParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithVersion adds the version to the delete deployment override params
func (o *DeleteDeploymentOverrideParams) WithVersion(version string) *DeleteDeploymentOverrideParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the delete deployment override params
func (o *DeleteDeploymentOverrideParams) SetVersion(version string) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteDeploymentOverrideParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param deployment
	if err := r.SetPathParam("deployment", o.Deployment); err != nil {
		return err
	}

	// path param namespace
	if err := r.SetPathParam("namespace", o.Namespace); err != nil {
		return err
	}

	// path param version
	if err := r.SetPathParam("version", o.Version); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
