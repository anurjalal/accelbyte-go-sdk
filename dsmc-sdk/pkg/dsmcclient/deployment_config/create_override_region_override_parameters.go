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

	"github.com/AccelByte/accelbyte-go-sdk/dsmc-sdk/pkg/dsmcclientmodels"
)

// NewCreateOverrideRegionOverrideParams creates a new CreateOverrideRegionOverrideParams object
// with the default values initialized.
func NewCreateOverrideRegionOverrideParams() *CreateOverrideRegionOverrideParams {
	var ()
	return &CreateOverrideRegionOverrideParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateOverrideRegionOverrideParamsWithTimeout creates a new CreateOverrideRegionOverrideParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateOverrideRegionOverrideParamsWithTimeout(timeout time.Duration) *CreateOverrideRegionOverrideParams {
	var ()
	return &CreateOverrideRegionOverrideParams{

		timeout: timeout,
	}
}

// NewCreateOverrideRegionOverrideParamsWithContext creates a new CreateOverrideRegionOverrideParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateOverrideRegionOverrideParamsWithContext(ctx context.Context) *CreateOverrideRegionOverrideParams {
	var ()
	return &CreateOverrideRegionOverrideParams{

		Context: ctx,
	}
}

// NewCreateOverrideRegionOverrideParamsWithHTTPClient creates a new CreateOverrideRegionOverrideParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateOverrideRegionOverrideParamsWithHTTPClient(client *http.Client) *CreateOverrideRegionOverrideParams {
	var ()
	return &CreateOverrideRegionOverrideParams{
		HTTPClient: client,
	}
}

/*CreateOverrideRegionOverrideParams contains all the parameters to send to the API endpoint
for the create override region override operation typically these are written to a http.Request
*/
type CreateOverrideRegionOverrideParams struct {

	/*Body*/
	Body *dsmcclientmodels.ModelsCreateRegionOverrideRequest
	/*Deployment
	  deployment of the game

	*/
	Deployment string
	/*Namespace
	  namespace of the game

	*/
	Namespace string
	/*Region
	  region

	*/
	Region string
	/*Version
	  version

	*/
	Version string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create override region override params
func (o *CreateOverrideRegionOverrideParams) WithTimeout(timeout time.Duration) *CreateOverrideRegionOverrideParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create override region override params
func (o *CreateOverrideRegionOverrideParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create override region override params
func (o *CreateOverrideRegionOverrideParams) WithContext(ctx context.Context) *CreateOverrideRegionOverrideParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create override region override params
func (o *CreateOverrideRegionOverrideParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create override region override params
func (o *CreateOverrideRegionOverrideParams) WithHTTPClient(client *http.Client) *CreateOverrideRegionOverrideParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create override region override params
func (o *CreateOverrideRegionOverrideParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create override region override params
func (o *CreateOverrideRegionOverrideParams) WithBody(body *dsmcclientmodels.ModelsCreateRegionOverrideRequest) *CreateOverrideRegionOverrideParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create override region override params
func (o *CreateOverrideRegionOverrideParams) SetBody(body *dsmcclientmodels.ModelsCreateRegionOverrideRequest) {
	o.Body = body
}

// WithDeployment adds the deployment to the create override region override params
func (o *CreateOverrideRegionOverrideParams) WithDeployment(deployment string) *CreateOverrideRegionOverrideParams {
	o.SetDeployment(deployment)
	return o
}

// SetDeployment adds the deployment to the create override region override params
func (o *CreateOverrideRegionOverrideParams) SetDeployment(deployment string) {
	o.Deployment = deployment
}

// WithNamespace adds the namespace to the create override region override params
func (o *CreateOverrideRegionOverrideParams) WithNamespace(namespace string) *CreateOverrideRegionOverrideParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the create override region override params
func (o *CreateOverrideRegionOverrideParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithRegion adds the region to the create override region override params
func (o *CreateOverrideRegionOverrideParams) WithRegion(region string) *CreateOverrideRegionOverrideParams {
	o.SetRegion(region)
	return o
}

// SetRegion adds the region to the create override region override params
func (o *CreateOverrideRegionOverrideParams) SetRegion(region string) {
	o.Region = region
}

// WithVersion adds the version to the create override region override params
func (o *CreateOverrideRegionOverrideParams) WithVersion(version string) *CreateOverrideRegionOverrideParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the create override region override params
func (o *CreateOverrideRegionOverrideParams) SetVersion(version string) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *CreateOverrideRegionOverrideParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param deployment
	if err := r.SetPathParam("deployment", o.Deployment); err != nil {
		return err
	}

	// path param namespace
	if err := r.SetPathParam("namespace", o.Namespace); err != nil {
		return err
	}

	// path param region
	if err := r.SetPathParam("region", o.Region); err != nil {
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
