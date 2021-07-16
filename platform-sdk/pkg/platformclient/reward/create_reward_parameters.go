// Code generated by go-swagger; DO NOT EDIT.

package reward

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

	"github.com/AccelByte/accelbyte-go-sdk/platform-sdk/pkg/platformclientmodels"
)

// NewCreateRewardParams creates a new CreateRewardParams object
// with the default values initialized.
func NewCreateRewardParams() *CreateRewardParams {
	var ()
	return &CreateRewardParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateRewardParamsWithTimeout creates a new CreateRewardParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateRewardParamsWithTimeout(timeout time.Duration) *CreateRewardParams {
	var ()
	return &CreateRewardParams{

		timeout: timeout,
	}
}

// NewCreateRewardParamsWithContext creates a new CreateRewardParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateRewardParamsWithContext(ctx context.Context) *CreateRewardParams {
	var ()
	return &CreateRewardParams{

		Context: ctx,
	}
}

// NewCreateRewardParamsWithHTTPClient creates a new CreateRewardParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateRewardParamsWithHTTPClient(client *http.Client) *CreateRewardParams {
	var ()
	return &CreateRewardParams{
		HTTPClient: client,
	}
}

/*CreateRewardParams contains all the parameters to send to the API endpoint
for the create reward operation typically these are written to a http.Request
*/
type CreateRewardParams struct {

	/*Body*/
	Body *platformclientmodels.RewardCreate
	/*Namespace*/
	Namespace string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create reward params
func (o *CreateRewardParams) WithTimeout(timeout time.Duration) *CreateRewardParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create reward params
func (o *CreateRewardParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create reward params
func (o *CreateRewardParams) WithContext(ctx context.Context) *CreateRewardParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create reward params
func (o *CreateRewardParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create reward params
func (o *CreateRewardParams) WithHTTPClient(client *http.Client) *CreateRewardParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create reward params
func (o *CreateRewardParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create reward params
func (o *CreateRewardParams) WithBody(body *platformclientmodels.RewardCreate) *CreateRewardParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create reward params
func (o *CreateRewardParams) SetBody(body *platformclientmodels.RewardCreate) {
	o.Body = body
}

// WithNamespace adds the namespace to the create reward params
func (o *CreateRewardParams) WithNamespace(namespace string) *CreateRewardParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the create reward params
func (o *CreateRewardParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WriteToRequest writes these params to a swagger request
func (o *CreateRewardParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
