// Code generated by go-swagger; DO NOT EDIT.

package third_party

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

// NewAdminCreateThirdPartyConfigParams creates a new AdminCreateThirdPartyConfigParams object
// with the default values initialized.
func NewAdminCreateThirdPartyConfigParams() *AdminCreateThirdPartyConfigParams {
	var ()
	return &AdminCreateThirdPartyConfigParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAdminCreateThirdPartyConfigParamsWithTimeout creates a new AdminCreateThirdPartyConfigParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAdminCreateThirdPartyConfigParamsWithTimeout(timeout time.Duration) *AdminCreateThirdPartyConfigParams {
	var ()
	return &AdminCreateThirdPartyConfigParams{

		timeout: timeout,
	}
}

// NewAdminCreateThirdPartyConfigParamsWithContext creates a new AdminCreateThirdPartyConfigParams object
// with the default values initialized, and the ability to set a context for a request
func NewAdminCreateThirdPartyConfigParamsWithContext(ctx context.Context) *AdminCreateThirdPartyConfigParams {
	var ()
	return &AdminCreateThirdPartyConfigParams{

		Context: ctx,
	}
}

// NewAdminCreateThirdPartyConfigParamsWithHTTPClient creates a new AdminCreateThirdPartyConfigParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAdminCreateThirdPartyConfigParamsWithHTTPClient(client *http.Client) *AdminCreateThirdPartyConfigParams {
	var ()
	return &AdminCreateThirdPartyConfigParams{
		HTTPClient: client,
	}
}

/*AdminCreateThirdPartyConfigParams contains all the parameters to send to the API endpoint
for the admin create third party config operation typically these are written to a http.Request
*/
type AdminCreateThirdPartyConfigParams struct {

	/*Body
	  third party config

	*/
	Body *lobbyclientmodels.ModelsCreateConfigRequest
	/*Namespace
	  namespace

	*/
	Namespace string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the admin create third party config params
func (o *AdminCreateThirdPartyConfigParams) WithTimeout(timeout time.Duration) *AdminCreateThirdPartyConfigParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the admin create third party config params
func (o *AdminCreateThirdPartyConfigParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the admin create third party config params
func (o *AdminCreateThirdPartyConfigParams) WithContext(ctx context.Context) *AdminCreateThirdPartyConfigParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the admin create third party config params
func (o *AdminCreateThirdPartyConfigParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the admin create third party config params
func (o *AdminCreateThirdPartyConfigParams) WithHTTPClient(client *http.Client) *AdminCreateThirdPartyConfigParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the admin create third party config params
func (o *AdminCreateThirdPartyConfigParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the admin create third party config params
func (o *AdminCreateThirdPartyConfigParams) WithBody(body *lobbyclientmodels.ModelsCreateConfigRequest) *AdminCreateThirdPartyConfigParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the admin create third party config params
func (o *AdminCreateThirdPartyConfigParams) SetBody(body *lobbyclientmodels.ModelsCreateConfigRequest) {
	o.Body = body
}

// WithNamespace adds the namespace to the admin create third party config params
func (o *AdminCreateThirdPartyConfigParams) WithNamespace(namespace string) *AdminCreateThirdPartyConfigParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the admin create third party config params
func (o *AdminCreateThirdPartyConfigParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WriteToRequest writes these params to a swagger request
func (o *AdminCreateThirdPartyConfigParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
